package logic

import (
	"context"
	"fmt"
	"reflect"

	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type User_updateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUser_updateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *User_updateLogic {
	return &User_updateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *User_updateLogic) User_update(req *types.UserInfoUpdateRequest) (resp *types.UserInfoUpdateResponse, err error) {
	// todo: add your logic here and delete this line
	fmt.Println(req.UserID)
	fmt.Println(RefToMap(*req, "user"))
	fmt.Println(RefToMap(*req, "user_conf"))

	return
}
func RefToMap(data any, tag string) map[string]any {
	maps := map[string]any{}
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		getTag, ok := field.Tag.Lookup(tag)
		if !ok {
			continue
		}
		val := v.Field(i)
		if val.IsZero() {
			continue
		}
		if field.Type.Kind() == reflect.Ptr {
			maps[getTag] = val.Elem().Interface()
		} else {
			maps[getTag] = val.Interface()
		}
	}
	return maps
}
