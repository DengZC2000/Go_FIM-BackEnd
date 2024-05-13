// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoRequest struct {
	UserID uint `header:"User-ID"`
	Role   int8 `header:"User-Role"`
}

type UserInfoResponse struct {
	UserID         uint   `json:"user_id"`
	Nickname       string `json:"nickname"`
	Role           int8   `json:"role"`
	Profile        string `json:"profile"`         //���˼���
	Avatar         string `json:"avatar"`          //ͷ��
	RegisterSource string `json:"register_source"` //ע����Դ qq ����
}
