package handler

import (
	"FIM/common/models/ctype"
	"FIM/common/response"
	"FIM/fim_chat/chat_api/internal/svc"
	"FIM/fim_chat/chat_api/internal/types"
	"FIM/fim_user/user_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"time"
)

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

		//把在线的用户存进一个公共的地方，哎~redis又用上了
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
				conn.WriteMessage(websocket.TextMessage, []byte("消息格式错误"))
				continue
			}
			//判断是否是好友
			isFriendRes, err := svcCtx.UserRpc.IsFriend(context.Background(), &user_rpc.IsFriendRequest{
				User1: uint32(req.UserID),
				User2: uint32(request.RevUserID),
			})
			if err != nil {
				logx.Error(err)
				conn.WriteMessage(websocket.TextMessage, []byte("用户服务错误"))
				continue
			}
			if !isFriendRes.IsFriend {
				conn.WriteMessage(websocket.TextMessage, []byte("你们还不是好友哦	"))
				continue
			}
			//入库
			//看看目标用户在不在线
			targetUserWs, ok := UserWsMap[request.RevUserID]
			if ok {
				//构造响应
				resp := ChatResponse{
					RevUser: ctype.UserInfo{
						ID:       request.RevUserID,
						Nickname: targetUserWs.UserInfo.NickName,
						Avatar:   targetUserWs.UserInfo.Avatar,
					},
					SendUser: ctype.UserInfo{
						ID:       req.UserID,
						Nickname: userInfo.NickName,
						Avatar:   userInfo.Avatar,
					},
					Msg:       request.Msg,
					CreatedAt: time.Now().String(),
				}
				byteData, _ := json.Marshal(resp)
				targetUserWs.Conn.WriteMessage(websocket.TextMessage, byteData)
			}

		}
	}
}

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
