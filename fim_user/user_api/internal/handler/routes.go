// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"FIM/fim_user/user_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/user/friend_info",
				Handler: friend_infoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/user_info",
				Handler: user_infoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/user/user_update",
				Handler: user_updateHandler(serverCtx),
			},
		},
	)
}
