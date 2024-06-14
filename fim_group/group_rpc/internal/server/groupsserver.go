// Code generated by goctl. DO NOT EDIT.
// Source: group_rpc.proto

package server

import (
	"context"

	"FIM/fim_group/group_rpc/internal/logic"
	"FIM/fim_group/group_rpc/internal/svc"
	"FIM/fim_group/group_rpc/types/group_rpc"
)

type GroupsServer struct {
	svcCtx *svc.ServiceContext
	group_rpc.UnimplementedGroupsServer
}

func NewGroupsServer(svcCtx *svc.ServiceContext) *GroupsServer {
	return &GroupsServer{
		svcCtx: svcCtx,
	}
}

func (s *GroupsServer) IsInGroup(ctx context.Context, in *group_rpc.IsInGroupRequest) (*group_rpc.IsInGroupResponse, error) {
	l := logic.NewIsInGroupLogic(ctx, s.svcCtx)
	return l.IsInGroup(in)
}

func (s *GroupsServer) UserGroupSearch(ctx context.Context, in *group_rpc.UserGroupSearchRequest) (*group_rpc.UserGroupSearchResponse, error) {
	l := logic.NewUserGroupSearchLogic(ctx, s.svcCtx)
	return l.UserGroupSearch(in)
}
