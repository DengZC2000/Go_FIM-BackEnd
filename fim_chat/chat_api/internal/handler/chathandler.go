package handler

import (
	"FIM/common/models/ctype"
	"FIM/common/response"
	"FIM/fim_chat/chat_api/internal/svc"
	"FIM/fim_chat/chat_api/internal/types"
	"FIM/fim_chat/chat_models"
	"FIM/fim_file/file_rpc/types/file_rpc"
	"FIM/fim_user/user_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

type ChatRequest struct {
	RevUserID uint      `json:"rev_user_id"`
	Msg       ctype.Msg `json:"msg"`
}

type ChatResponse struct {
	SendUser  ctype.UserInfo `json:"send_user"`
	RevUser   ctype.UserInfo `json:"rev_user"`
	Msg       ctype.Msg      `json:"msg"`
	CreatedAt string         `json:"created_at"`
}

type UserWsInfo struct {
	UserInfo user_models.UserModel //用户信息
	Conn     *websocket.Conn       //用户的ws连接对象
}

var UserWsMap = map[uint]UserWsInfo{}

func chat_Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRequest
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

		defer func() {
			conn.Close()
			delete(UserWsMap, req.UserID)
			svcCtx.Redis.HDel(context.Background(), "online", fmt.Sprintf("%d", req.UserID))
		}()
		//调用户服务，获取当前用户信息
		res, err := svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{
			UserId: uint32(req.UserID),
		})
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		var userInfo user_models.UserModel
		err = json.Unmarshal(res.Data, &userInfo)
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		var userWsInfo = UserWsInfo{
			UserInfo: userInfo,
			Conn:     conn,
		}
		UserWsMap[req.UserID] = userWsInfo

		//把在线的用户存进一个公共的地方，哎~ redis又用上了
		svcCtx.Redis.HSet(context.Background(), "online", fmt.Sprintf("%d", req.UserID), req.UserID)

		//查一下自己的好友是不是上线了
		friendRes, err := svcCtx.UserRpc.FriendList(context.Background(), &user_rpc.FriendListRequest{User: uint32(req.UserID)})
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		for _, info := range friendRes.FriendList {
			friend, ok := UserWsMap[uint(info.UserId)]
			if ok {
				//那他是否开启了好友上线提醒功能
				if friend.UserInfo.UserConfModel.FriendOnline {
					//好友上线了,向已经在线的我的好友，发送一个消息
					friend.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("好友 %s 上线了", userInfo.NickName)))
				}

			}
		}
		//查一下自己的好友列表，返回用户id列表，看看在不在这个UserMsMap中，在的话，就给自己发送个好友上线的消息

		fmt.Println(UserWsMap)
		for {
			//消息类型，消息，错误
			_, p, err := conn.ReadMessage()
			if err != nil {
				//用户断开聊天
				fmt.Println(err)
				break
			}
			var request ChatRequest
			err1 := json.Unmarshal(p, &request)
			if err1 != nil {
				//用户乱发消息
				logx.Error(err1)
				SendTipErrMsg(conn, "参数解析失败")
				continue
			}
			if req.UserID != request.RevUserID {

				//判断是否是好友
				isFriendRes, err := svcCtx.UserRpc.IsFriend(context.Background(), &user_rpc.IsFriendRequest{
					User1: uint32(req.UserID),
					User2: uint32(request.RevUserID),
				})
				if err != nil {
					logx.Error(err)
					SendTipErrMsg(conn, "用户服务错误")
					continue
				}
				if !isFriendRes.IsFriend {
					SendTipErrMsg(conn, "你们还不是好友哦")
					continue
				}
			}
			//判断是否是文件类型
			switch request.Msg.Type {
			case ctype.FileMsgType:
				//如果是文件类型,那么就要去请求rpc服务了,获取文件信息
				nameList := strings.Split(request.Msg.FileMsg.Src, "/")
				if len(nameList) == 0 {
					SendTipErrMsg(conn, "请上传文件")
					continue
				}
				fileID := nameList[len(nameList)-1]
				fileResponse, err3 := svcCtx.FileRpc.FileInfo(context.Background(), &file_rpc.FileInfoRequest{
					FileId: fileID,
				})
				if err3 != nil {
					logx.Error(err3)
					SendTipErrMsg(conn, err3.Error())
					continue
				}
				request.Msg.FileMsg.Title = fileResponse.FileName
				request.Msg.FileMsg.Size = fileResponse.FileSize
				request.Msg.FileMsg.Type = fileResponse.FileType
			case ctype.WithdrawMsgType:
				//撤回消息的id是必填的
				if request.Msg.WithdrawMsg.MsgID == 0 {
					SendTipErrMsg(conn, "撤回消息是必填的")
					continue
				}
				//自己只能撤回自己的
				//找这个消息是谁发的
				var msgModel chat_models.ChatModel
				err = svcCtx.DB.Take(&msgModel, request.Msg.WithdrawMsg.MsgID).Error
				if err != nil {
					SendTipErrMsg(conn, "消息不存在")
					continue
				}
				//判断是不是自己发的
				if msgModel.SendUserID != req.UserID {
					SendTipErrMsg(conn, "只能撤回自己的消息")
					continue
				}
				now := time.Now()
				subTime := now.Sub(msgModel.CreatedAt)
				if subTime >= time.Minute*2 {
					SendTipErrMsg(conn, "只能撤回2分钟之内的消息")
					continue
				}
				//撤回逻辑
				//收到撤回请求之后，服务端这边把原来消息类型修改为撤回消息，并且记录原消息
				//然后通知前端的收发双方，重新拉取聊天记录
				var content = "xx 撤回了一条消息"
				if userInfo.UserConfModel.RecallMessage != nil {
					content += *userInfo.UserConfModel.RecallMessage
				}
				svcCtx.DB.Model(&msgModel).Updates(chat_models.ChatModel{
					Msg: &ctype.Msg{
						Type: ctype.WithdrawMsgType,
						WithdrawMsg: &ctype.WithdrawMsg{
							Content:   content,
							MsgID:     request.Msg.WithdrawMsg.MsgID,
							OriginMsg: msgModel.Msg,
						},
					},
				})

			}
			//入库
			ChatMsgIntoDataBase(svcCtx.DB, req.UserID, request.RevUserID, &request.Msg)

			//调用封装方法，发送信息,其中判断了是否在线
			//并且应该双方都发消息，真实的场景啊
			SendMsgByUser(req.UserID, request.RevUserID, request.Msg)

		}
	}
}

// ChatMsgIntoDataBase 数据入库
func ChatMsgIntoDataBase(db *gorm.DB, sendUserID uint, revUserID uint, msg *ctype.Msg) {
	switch msg.Type {
	case ctype.WithdrawMsgType:
		fmt.Println("撤回消息是不入库的")
		return
	}
	chatModel := chat_models.ChatModel{
		SendUserID: sendUserID,
		RevUserID:  revUserID,
		MsgType:    msg.Type,
		Msg:        msg,
	}
	chatModel.MsgPreview = chatModel.MsgPreviewMethod()
	err := db.Create(&chatModel).Error
	if err != nil {
		logx.Error(err)
		sendUser, ok := UserWsMap[sendUserID]
		if !ok {
			return
		}
		SendTipErrMsg(sendUser.Conn, "消息保存失败")
	}
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
		CreatedAt: time.Now().String(),
	}
	byteData, _ := json.Marshal(resp)
	Conn.WriteMessage(websocket.TextMessage, byteData)
}

// SendMsgByUser 发消息 谁发的 给谁发
func SendMsgByUser(sendUserID uint, revUserID uint, msg ctype.Msg) {
	revUser, ok := UserWsMap[revUserID]
	if !ok {
		return
	}
	sendUser, ok := UserWsMap[sendUserID]
	if !ok {
		return
	}
	//构造响应
	resp := ChatResponse{
		RevUser: ctype.UserInfo{
			ID:       revUserID,
			Nickname: revUser.UserInfo.NickName,
			Avatar:   revUser.UserInfo.Avatar,
		},
		SendUser: ctype.UserInfo{
			ID:       sendUserID,
			Nickname: sendUser.UserInfo.NickName,
			Avatar:   sendUser.UserInfo.Avatar,
		},
		Msg:       msg,
		CreatedAt: time.Now().String(),
	}
	byteData, _ := json.Marshal(resp)
	revUser.Conn.WriteMessage(websocket.TextMessage, byteData)

	//并且应该双方都发消息，真实的场景啊
	sendUser.Conn.WriteMessage(websocket.TextMessage, byteData)

}
