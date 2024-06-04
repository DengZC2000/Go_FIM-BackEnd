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
				Path:    "/api/group/group_add",
				Handler: group_addHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/group/group_add_member",
				Handler: group_add_memberHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/group/group_create",
				Handler: group_createHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/group/group_delete/:id",
				Handler: group_deleteHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/group/group_delete_member",
				Handler: group_delete_memberHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/group_history/:id",
				Handler: group_historyHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/group/group_history_delete/:id",
				Handler: group_history_deleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/group_info/:id",
				Handler: group_infoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/group_member",
				Handler: group_memberHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/group_my_friends",
				Handler: group_my_friendsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/group_session",
				Handler: group_sessionHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/group/group_update",
				Handler: group_updateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/group/group_update_nickname",
				Handler: group_update_nicknameHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/group/group_update_role",
				Handler: group_update_roleHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/group/group_user_prohibition",
				Handler: group_user_prohibitionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/group_valid",
				Handler: group_validHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/group_valid_list",
				Handler: group_valid_listHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/group/group_valid_status",
				Handler: group_valid_statusHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/search",
				Handler: group_searchHandler(serverCtx),
			},
		},
	)
}
