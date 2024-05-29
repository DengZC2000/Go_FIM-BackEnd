// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"FIM/fim_group/group_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/group/group_create",
				Handler: group_createHandler(serverCtx),
			},
		},
	)
}