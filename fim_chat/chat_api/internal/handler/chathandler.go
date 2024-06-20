package handler

import (
	"FIM/common/models/ctype"
	"FIM/common/response"
	"FIM/common/service/redis_service"
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
	ChatID    uint           `json:"chat_id"`
	IsMe      bool           `json:"is_me"`
	SendUser  ctype.UserInfo `json:"send_user"`
	RevUser   ctype.UserInfo `json:"rev_user"`
	Msg       ctype.Msg      `json:"msg"`
	CreatedAt string         `json:"created_at"`
}

type UserWsInfo struct {
	UserInfo    user_models.UserModel      //用户信息
	WsClientMap map[string]*websocket.Conn //这个用户管理的所有ws客户端
	CurrentConn *websocket.Conn            //当前的连接对象
}

var UserOnlineWsMap = map[uint]*UserWsInfo{}
var VideoCallStartTime = map[string]time.Time{}

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
		//获取这次的连接地址
		addr := conn.RemoteAddr().String()
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
				svcCtx.Redis.HDel(context.Background(), "online", fmt.Sprintf("%d", req.UserID))
			}

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

		logx.Info(addr)
		userWsInfo, ok := UserOnlineWsMap[req.UserID]
		if !ok {
			//如果这个用户第一次来
			UserOnlineWsMap[req.UserID] = &UserWsInfo{
				UserInfo:    userInfo,
				WsClientMap: map[string]*websocket.Conn{addr: conn},
				CurrentConn: conn,
			}
		} else {
			_, ok1 := userWsInfo.WsClientMap[addr]
			if !ok1 {
				//这个用户不是第一次来，那判断是不是这个用户二开
				userWsInfo.WsClientMap[addr] = conn
				userWsInfo.CurrentConn = conn
			}
		}

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
			friend, ok := UserOnlineWsMap[uint(info.UserId)]
			if ok {
				//那他是否开启了好友上线提醒功能
				if friend.UserInfo.UserConfModel.FriendOnline {
					//好友上线了,向已经在线的我的好友，发送一个消息
					DistributeWsMsgMap(friend.WsClientMap, []byte(fmt.Sprintf("好友 %s 上线了", userInfo.NickName)))
				}

			}
		}
		//查一下自己的好友列表，返回用户id列表，看看在不在这个UserMsMap中，在的话，就给自己发送个好友上线的消息

		fmt.Println(UserOnlineWsMap)
		for {
			//消息类型，消息，错误
			_, p, err := conn.ReadMessage()
			if err != nil {
				//用户断开聊天
				fmt.Println(err)
				break
			}
			// 目前这里没做实时更新
			// 要做到实时更新，把用户的这些配置放到缓存里去
			// 用户聊天之前向缓存中去拿用户的相关配置信息，拿不到的情况下，去调用户rpc返回信息，顺便把这些信息放到缓存中
			// 在后台，用户的配置更新之后，就让这条缓存失效
			if userInfo.UserConfModel.RestrictChat {
				SendTipErrMsg(conn, "当前用户被限制聊天")
				continue
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
			// 判断type 1 - 12
			if !(request.Msg.Type >= 1 && request.Msg.Type <= 12) {
				SendTipErrMsg(conn, "消息类型错误，未知的消息类型")
				continue
			}
			// 判断是否是文件类型
			switch request.Msg.Type {
			case ctype.TextMsgType:
				if request.Msg.TextMsg == nil {
					SendTipErrMsg(conn, "请输入消息内容")
					continue
				}
				if request.Msg.TextMsg.Content == "" {
					SendTipErrMsg(conn, "请输入消息内容")
					continue
				}
			case ctype.FileMsgType:
				//如果是文件类型,那么就要去请求rpc服务了,获取文件信息
				if request.Msg.FileMsg == nil {
					SendTipErrMsg(conn, "请上传文件")
					continue
				}
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
				if request.Msg.WithdrawMsg == nil {
					SendTipErrMsg(conn, "撤回消息是必填的")
					continue
				}
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
				//不能撤回已撤回的消息
				if msgModel.MsgType == ctype.WithdrawMsgType {
					SendTipErrMsg(conn, "撤回消息不能再撤回了")
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
					logx.Info(subTime)
					SendTipErrMsg(conn, "只能撤回2分钟之内的消息")
					continue
				}
				//撤回逻辑
				//收到撤回请求之后，服务端这边把原来消息类型修改为撤回消息，并且记录原消息
				//然后通知前端的收发双方，重新拉取聊天记录
				var content = "撤回了一条消息"
				if userInfo.UserConfModel.RecallMessage != nil {
					content += *userInfo.UserConfModel.RecallMessage
				}
				content = "你" + content
				//前端可以判断，这个消息如果不是isMe，就可以把你替换成对方的昵称
				originMsg := msgModel.Msg
				originMsg.WithdrawMsg = nil
				svcCtx.DB.Model(&msgModel).Updates(chat_models.ChatModel{
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
			case ctype.ReplyMsgType:
				//回复消息
				//先校验
				if request.Msg.ReplyMsg == nil || request.Msg.ReplyMsg.MsgID == 0 {
					SendTipErrMsg(conn, "回复消息必填")
					continue
				}
				//找这个原消息
				var msgModel chat_models.ChatModel
				err = svcCtx.DB.Take(&msgModel, request.Msg.ReplyMsg.MsgID).Error
				if err != nil {
					SendTipErrMsg(conn, "消息不存在")
					continue
				}
				//不能引用撤回消息
				if msgModel.MsgType == ctype.WithdrawMsgType {
					SendTipErrMsg(conn, "该消息已撤回")
					continue
				}
				//回复的这个消息，必须是你自己或者当前和你聊天这个人发出来的
				if !((msgModel.SendUserID == req.UserID && msgModel.RevUserID == request.RevUserID) ||
					(msgModel.SendUserID == request.RevUserID && msgModel.RevUserID == req.UserID)) {
					SendTipErrMsg(conn, "只能回复自己或者对方的消息")
					continue
				}
				SendBaseInfo, err2 := redis_service.GetUserBaseInfoByRedis(svcCtx.Redis, svcCtx.UserRpc, msgModel.SendUserID)
				if err2 != nil {
					logx.Error(err2)
					return
				}
				request.Msg.ReplyMsg.Msg = msgModel.Msg
				request.Msg.ReplyMsg.UserID = msgModel.SendUserID
				request.Msg.ReplyMsg.UserNickName = SendBaseInfo.Nickname
				request.Msg.ReplyMsg.OriginMsgDate = msgModel.CreatedAt
			case ctype.QuoteMsgType:
				//回复消息
				//先校验
				if request.Msg.QuoteMsg == nil || request.Msg.QuoteMsg.MsgID == 0 {
					SendTipErrMsg(conn, "引用消息必填")
					continue
				}
				//找这个原消息
				var msgModel chat_models.ChatModel
				err = svcCtx.DB.Take(&msgModel, request.Msg.QuoteMsg.MsgID).Error
				if err != nil {
					SendTipErrMsg(conn, "消息不存在")
					continue
				}
				//不能引用撤回消息
				if msgModel.MsgType == ctype.WithdrawMsgType {
					SendTipErrMsg(conn, "该消息已撤回")
					continue
				}
				//引用的这个消息，必须是你自己或者当前和你聊天这个人发出来的
				if !((msgModel.SendUserID == req.UserID && msgModel.RevUserID == request.RevUserID) ||
					(msgModel.SendUserID == request.RevUserID && msgModel.RevUserID == req.UserID)) {
					SendTipErrMsg(conn, "只能引用自己或者对方的消息")
					continue
				}
				SendBaseInfo, err2 := redis_service.GetUserBaseInfoByRedis(svcCtx.Redis, svcCtx.UserRpc, msgModel.SendUserID)
				if err2 != nil {
					logx.Error(err2)
					return
				}
				request.Msg.QuoteMsg.Msg = msgModel.Msg
				request.Msg.QuoteMsg.UserID = msgModel.SendUserID
				request.Msg.QuoteMsg.UserNickName = SendBaseInfo.Nickname
				request.Msg.QuoteMsg.OriginMsgDate = msgModel.CreatedAt
			case ctype.VideoCallMsgType:
				data := request.Msg.VideoCallMsg
				//先判断对方是否在线
				_, ok2 := UserOnlineWsMap[request.RevUserID]
				if !ok2 {
					SendTipErrMsg(conn, "对方不在线")
					continue
				}
				key := fmt.Sprintf("%d_%d", userInfo.ID, request.RevUserID)
				switch data.Flag {
				case 0:
					// 给自己的页面展示一个等待对方接听的弹框
					conn.WriteJSON(ChatResponse{
						Msg: ctype.Msg{
							VideoCallMsg: &ctype.VideoCallMsg{
								Flag: 1,
							},
						},
					})
					// 给对方的页面展示一个等待对方接听的弹框
					sendRevUserMsg(request.RevUserID, ctype.Msg{
						VideoCallMsg: &ctype.VideoCallMsg{
							Flag: 2,
						},
					})
				case 1: // 自己挂断
					sendRevUserMsg(request.RevUserID, ctype.Msg{
						VideoCallMsg: &ctype.VideoCallMsg{
							Flag: 3,
							Msg:  "发起者已挂断",
						},
					})
				case 2: // 对方挂断
					// 对方点击挂断，那么它的目标就是revUserID，也就是上面的conn
					sendRevUserMsg(request.RevUserID, ctype.Msg{
						VideoCallMsg: &ctype.VideoCallMsg{
							Flag: 3,
							Msg:  "接收者已挂断",
						},
					})
				case 3: //对方接受
					// 让发起者准备去发offer
					sendRevUserMsg(request.RevUserID, ctype.Msg{
						VideoCallMsg: &ctype.VideoCallMsg{
							Flag: 5, // 让发起者准备去发offer
							Type: "create_offer",
						},
					})
				case 4:
					// 我方正常挂断
					// 算你们的通话时长
					// 从发offer开始，算一个开始时间，到这里算一个结束时间，就是视频通话的时间
					fmt.Println("用户正常挂断")
					startTime, ok3 := VideoCallStartTime[key]
					if !ok3 {
						fmt.Println("没有开始时间")
						continue
					}
					CallTime := time.Now().Sub(startTime)
					fmt.Printf("通话时长 %s", CallTime.String())
				case 5:
					// 对方挂断
					key = fmt.Sprintf("%d_%d", request.RevUserID, userInfo.ID)
					startTime, ok3 := VideoCallStartTime[key]
					if !ok3 {
						fmt.Println("没有开始时间")
						continue
					}
					CallTime := time.Now().Sub(startTime)
					fmt.Printf("通话时长 %s", CallTime.String())
				}
				switch data.Type {
				case "offer":
					sendRevUserMsg(request.RevUserID, ctype.Msg{
						VideoCallMsg: &ctype.VideoCallMsg{
							Flag: 5, // 让发起者准备去发offer
							Type: "offer",
							Data: data.Data,
						},
					})
					VideoCallStartTime[key] = time.Now()
				case "answer":
					conn.WriteJSON(ChatResponse{
						Msg: ctype.Msg{
							VideoCallMsg: &ctype.VideoCallMsg{
								Type: "answer",
								Data: data.Data,
							},
						},
					})
				case "offer_ice":
					sendRevUserMsg(request.RevUserID, ctype.Msg{
						VideoCallMsg: &ctype.VideoCallMsg{
							Flag: 5, // 让发起者准备去发offer
							Type: "offer_ice",
							Data: data.Data,
						},
					})
				case "answer_ice":
					conn.WriteJSON(ChatResponse{
						Msg: ctype.Msg{
							VideoCallMsg: &ctype.VideoCallMsg{
								Type: "answer_ice",
								Data: data.Data,
							},
						},
					})
				}
				// 自己这方可以挂断

				// 对方也可以挂断

				// 如果对方开了多个浏览器，只用找其中的一个，找第一个
				continue

			}
			//入库,里面有撤回消息是不入库的逻辑
			chatID := ChatMsgIntoDataBase(svcCtx.DB, req.UserID, request.RevUserID, &request.Msg)
			//如果是撤回消息，此时的chatID是0，所以给一个被撤回的原消息的id值。
			if request.Msg.Type == ctype.WithdrawMsgType {
				chatID = request.Msg.WithdrawMsg.MsgID
			}
			//调用封装方法，发送信息,其中判断了是否在线
			//并且应该双方都发消息，真实的场景啊
			SendMsgByUser(svcCtx, req.UserID, request.RevUserID, request.Msg, chatID)

		}
	}
}

// DistributeWsMsgMap 给一组的ws对象发消息
func DistributeWsMsgMap(wsMap map[string]*websocket.Conn, byteData []byte) {
	for _, conn := range wsMap {
		conn.WriteMessage(websocket.TextMessage, byteData)
	}
}

// sendRevUserMsg 给目标客户端发消息
func sendRevUserMsg(revUserID uint, msg ctype.Msg) {
	userRes, ok := UserOnlineWsMap[revUserID]
	if !ok {
		return
	}
	for _, conn := range userRes.WsClientMap {
		conn.WriteJSON(ChatResponse{
			SendUser: ctype.UserInfo{},
			RevUser: ctype.UserInfo{
				ID:       userRes.UserInfo.ID,
				Nickname: userRes.UserInfo.NickName,
				Avatar:   userRes.UserInfo.Avatar,
			},
			Msg:       msg,
			CreatedAt: time.Now().String(),
		})
		break
	}
}

// ChatMsgIntoDataBase 数据入库
func ChatMsgIntoDataBase(db *gorm.DB, sendUserID uint, revUserID uint, msg *ctype.Msg) (MsgID uint) {
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
		sendUser, ok := UserOnlineWsMap[sendUserID]
		if !ok {
			return
		}
		SendTipErrMsg(sendUser.CurrentConn, "消息保存失败")
	}
	return chatModel.ID
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
func SendMsgByUser(SvcCtx *svc.ServiceContext, sendUserID uint, revUserID uint, msg ctype.Msg, chatID uint) {
	revUser, ok1 := UserOnlineWsMap[revUserID]
	sendUser, ok2 := UserOnlineWsMap[sendUserID]
	resp := ChatResponse{
		Msg:       msg,
		CreatedAt: time.Now().String(),
		ChatID:    chatID,
	}

	if ok1 && ok2 && sendUserID == revUserID {
		resp.IsMe = true
		//百分百是自己与自己发消息了
		resp.SendUser = ctype.UserInfo{
			ID:       sendUserID,
			Nickname: sendUser.UserInfo.NickName,
			Avatar:   sendUser.UserInfo.Avatar,
		}
		resp.RevUser = ctype.UserInfo{
			ID:       revUserID,
			Nickname: revUser.UserInfo.NickName,
			Avatar:   revUser.UserInfo.Avatar,
		}
		byteData, _ := json.Marshal(resp)
		DistributeWsMsgMap(sendUser.WsClientMap, byteData)
		return
	}
	if !ok1 {
		//如果接收者不在线，调redis获得接收者信息
		RevBaseInfo, err2 := redis_service.GetUserBaseInfoByRedis(SvcCtx.Redis, SvcCtx.UserRpc, revUserID)
		if err2 != nil {
			logx.Error(err2)
			return
		}
		resp.RevUser = ctype.UserInfo{
			ID:       RevBaseInfo.ID,
			Nickname: RevBaseInfo.Nickname,
			Avatar:   RevBaseInfo.Avatar,
		}

	} else {
		//接收者在线
		resp.RevUser = ctype.UserInfo{
			ID:       revUserID,
			Nickname: revUser.UserInfo.NickName,
			Avatar:   revUser.UserInfo.Avatar,
		}

	}
	//发送者肯定在线
	resp.SendUser = ctype.UserInfo{
		ID:       sendUserID,
		Nickname: sendUser.UserInfo.NickName,
		Avatar:   sendUser.UserInfo.Avatar,
	}

	RevByteData, _ := json.Marshal(resp)
	if ok1 {
		DistributeWsMsgMap(revUser.WsClientMap, RevByteData)
	}
	resp.IsMe = true
	SendByteData, _ := json.Marshal(resp)
	DistributeWsMsgMap(sendUser.WsClientMap, SendByteData)

}
