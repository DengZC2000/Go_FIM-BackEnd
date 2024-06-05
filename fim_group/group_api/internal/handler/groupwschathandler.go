package handler

import (
	"FIM/common/models/ctype"
	"FIM/common/response"
	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"
	"FIM/fim_group/group_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type UserWsInfo struct {
	UserInfo    ctype.UserInfo
	WsClientMap map[string]*websocket.Conn
}

var UserOnlineWsMap = map[uint]*UserWsInfo{}

type ChatRequest struct {
	GroupID uint      `json:"group_id"` //群id
	Msg     ctype.Msg `json:"msg"`      //消息
}
type ChatResponse struct {
	UserID       uint          `json:"user_id"`
	UserNickname string        `json:"user_nickname"`
	UserAvatar   string        `json:"user_avatar"`
	Msg          ctype.Msg     `json:"msg"`
	ID           uint          `json:"id"`
	MsgType      ctype.MsgType `json:"msg_type"`
	CreatedAt    time.Time     `json:"created_at"`
	IsMe         bool          `json:"is_me"`
}

func group_ws_chatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupChatRequest
		if err := httpx.ParseHeaders(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r, w, nil, err)
			return
		}

		var upGrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				//鉴权 true表示放行，false表示拦截
				return true
			},
		}
		conn, err := upGrader.Upgrade(w, r, nil)
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		//获取这次的连接地址
		addr := conn.RemoteAddr().String()
		fmt.Printf("用户连接成功！地址：%s", addr)
		defer func() {
			conn.Close()
			userWsTarget, ok := UserOnlineWsMap[req.UserID]
			if ok {
				//删除退出的那个conn
				delete(userWsTarget.WsClientMap, addr)
			}
			if userWsTarget != nil && len(userWsTarget.WsClientMap) == 0 {
				//如果都退出了，就下线
				delete(UserOnlineWsMap, req.UserID)
			}

		}()
		baseInfoResponse, err := svcCtx.UserRpc.UserBaseInfo(context.Background(), &user_rpc.UserBaseInfoRequest{
			UserId: uint32(req.UserID),
		})
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		userInfo := ctype.UserInfo{
			ID:       req.UserID,
			Nickname: baseInfoResponse.NickName,
			Avatar:   baseInfoResponse.Avatar,
		}
		userWsInfo, ok := UserOnlineWsMap[req.UserID]
		if !ok {
			//如果这个用户第一次来
			UserOnlineWsMap[req.UserID] = &UserWsInfo{
				UserInfo:    userInfo,
				WsClientMap: map[string]*websocket.Conn{addr: conn},
			}
		} else {
			_, ok1 := userWsInfo.WsClientMap[addr]
			if !ok1 {
				//这个用户不是第一次来，那判断是不是这个用户二开
				userWsInfo.WsClientMap[addr] = conn
			}
		}
		for {
			//消息类型，消息，错误
			_, p, err := conn.ReadMessage()
			if err != nil {
				//用户断开聊天
				fmt.Println(err)
				break
			}
			var request ChatRequest
			err = json.Unmarshal(p, &request)
			if err != nil {
				SendTipErrMsg(conn, "参数解析失败")
				continue
			}
			// 判断自己是不是这个群的成员
			var member group_models.GroupMemberModel
			err = svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", request.GroupID, req.UserID).Error
			if err != nil {
				SendTipErrMsg(conn, "你不是该群成员")
				continue
			}
			switch request.Msg.Type {
			case ctype.WithdrawMsgType:
				//校验
				withdrawMsg := request.Msg.WithdrawMsg
				if withdrawMsg == nil {
					SendTipErrMsg(conn, "撤回消息的格式错误")
					continue
				}
				if withdrawMsg.MsgID == 0 {
					SendTipErrMsg(conn, "撤回消息id为空")
					continue
				}
				// 去找消息
				var groupMsg group_models.GroupMsgModel
				err = svcCtx.DB.Take(&groupMsg, "group_id = ? and id = ?", request.GroupID, withdrawMsg.MsgID).Error
				if err != nil {
					SendTipErrMsg(conn, "原消息不存在")
					continue
				}
				// 原消息不能是撤回消息
				if groupMsg.MsgType == ctype.WithdrawMsgType {
					SendTipErrMsg(conn, "不能撤回已经撤回的消息")
					continue
				}
				// 拿我在这个群的角色
				// 如果是我自己撤的 并且是普通用户
				// 要判断时间是否在2分钟之内
				if member.Role == 3 && req.UserID != groupMsg.SendUserID {
					SendTipErrMsg(conn, "普通用户不能撤回别人的消息")
					continue
				}
				if req.UserID == groupMsg.SendUserID && member.Role == 3 {
					now := time.Now()
					if now.Sub(groupMsg.CreatedAt) > 2*time.Minute {
						SendTipErrMsg(conn, "只能撤回两分钟之内的消息")
						continue
					}
				}
				// 撤回不是自己的消息，先去查原消息的人的角色
				var msgUserRole int8 = 3
				err = svcCtx.DB.Model(group_models.GroupMemberModel{}).
					Where("group_id = ? and user_id = ?", request.GroupID, groupMsg.SendUserID).
					Select("role").
					Scan(&msgUserRole).Error

				// 这里有可能查不到， 原因是用户退群了
				// 如果是管理员 能撤回自己和用户的，没有时间限制
				if member.Role == 2 {
					if msgUserRole == 1 || (msgUserRole == 2 && req.UserID != groupMsg.SendUserID) {
						SendTipErrMsg(conn, "管理员只能撤回自己和用户的")
						continue
					}
				}
				// 如果是群主，直接没有限制
				// 层层筛选，来到最终地方
				var content = "撤回了一条消息"
				content = "你" + content
				//前端可以判断，这个消息如果不是isMe，就可以把你替换成对方的昵称
				originMsg := groupMsg.Msg
				originMsg.WithdrawMsg = nil // 要不然会出现循环引用
				svcCtx.DB.Model(&groupMsg).Updates(group_models.GroupMsgModel{
					MsgType:    ctype.WithdrawMsgType,
					MsgPreview: "[撤回消息]- " + content,
					Msg: &ctype.Msg{
						Type: ctype.WithdrawMsgType,
						WithdrawMsg: &ctype.WithdrawMsg{
							Content:   content,
							MsgID:     request.Msg.WithdrawMsg.MsgID,
							OriginMsg: originMsg,
						},
					},
				})
			}
			//消息入库
			groupMsgID := GroupMsgIntoDataBase(svcCtx.DB, conn, request.GroupID, req.UserID, &request.Msg)
			// 遍历这个用户列表，去找ws的客户端
			sendGroupOnlineUserMsg(svcCtx.DB,
				request.GroupID,
				req.UserID,
				request.Msg,
				groupMsgID,
			)
		}
	}
}
func GroupMsgIntoDataBase(DB *gorm.DB, conn *websocket.Conn, groupID uint, userID uint, msg *ctype.Msg) uint {
	switch msg.Type {
	case ctype.WithdrawMsgType:
		fmt.Println("撤回消息是不入库的")
		return 0
	}
	groupModel := group_models.GroupMsgModel{
		GroupID:    groupID,
		SendUserID: userID,
		MsgType:    msg.Type,
		Msg:        msg,
	}
	groupModel.MsgPreview = groupModel.MsgPreviewMethod()
	err := DB.Create(&groupModel).Error
	if err != nil {
		logx.Error(err)
		SendTipErrMsg(conn, "消息保存失败")
		return 0
	}
	return groupModel.ID
}

// sendGroupOnlineUserMsg 给这个群的用户发消息
func sendGroupOnlineUserMsg(DB *gorm.DB, groupID uint, userID uint, msg ctype.Msg, msgID uint) {
	// 查在线用户
	onlineUserIDList := getOnlineUserIDList()

	// 查这个群的成员，且在线
	var groupMemberOnlineIDList []uint
	DB.Model(group_models.GroupMemberModel{}).
		Where("group_id = ? and user_id in ?", groupID, onlineUserIDList).
		Select("user_id").
		Scan(&groupMemberOnlineIDList)
	// 构造响应
	var chatResponse = ChatResponse{
		UserID:    userID,
		Msg:       msg,
		ID:        msgID,
		MsgType:   msg.Type,
		CreatedAt: time.Now(),
	}
	wsInfo, ok := UserOnlineWsMap[userID]
	if ok {
		chatResponse.UserAvatar = wsInfo.UserInfo.Avatar
		chatResponse.UserNickname = wsInfo.UserInfo.Nickname
	}

	// 遍历这个用户列表，去找ws的客户端
	for _, u := range groupMemberOnlineIDList {
		wsUserInfo, ok2 := UserOnlineWsMap[u]
		if !ok2 {
			continue
		}
		//判断isMe
		if wsUserInfo.UserInfo.ID == userID {
			chatResponse.IsMe = true
		} else {
			chatResponse.IsMe = false
		}
		byteData, _ := json.Marshal(chatResponse)
		for _, ws := range wsUserInfo.WsClientMap {
			ws.WriteMessage(websocket.TextMessage, byteData)
		}
	}
}

func getOnlineUserIDList() (OnlineUserIDList []uint) {
	OnlineUserIDList = make([]uint, 0)
	for u, _ := range UserOnlineWsMap {
		OnlineUserIDList = append(OnlineUserIDList, u)
	}
	return
}

// SendTipErrMsg 发送错误提示的消息
func SendTipErrMsg(Conn *websocket.Conn, msg string) {
	resp := ChatResponse{
		Msg: ctype.Msg{
			Type: ctype.TipMsgType,
			TipMsg: &ctype.TipMsg{
				Status:  "error",
				Content: msg,
			},
		},
		CreatedAt: time.Now(),
	}
	byteData, _ := json.Marshal(resp)
	Conn.WriteMessage(websocket.TextMessage, byteData)
}
