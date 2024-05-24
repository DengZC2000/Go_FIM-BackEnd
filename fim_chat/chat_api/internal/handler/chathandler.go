package handler

import (
	"FIM/common/models/ctype"
	"FIM/common/response"
	"FIM/fim_chat/chat_api/internal/svc"
	"FIM/fim_chat/chat_api/internal/types"
	"FIM/fim_chat/chat_models"
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
			//入库
			ChatMsgIntoDataBase(svcCtx.DB, req.UserID, request.RevUserID, &request.Msg)

			//调用封装方法，发送信息,其中判断了是否在线
			SendMsgByUser(req.UserID, request.RevUserID, request.Msg)

		}
	}
}

// ChatMsgIntoDataBase 数据入库
func ChatMsgIntoDataBase(db *gorm.DB, sendUserID uint, revUserID uint, msg *ctype.Msg) {
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

}
