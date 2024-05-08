// Code generated by goctl. DO NOT EDIT.
package types

type Authentication struct {
}

type AuthenticationResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

type LoginInfo struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code int       `json:"code"`
	Data LoginInfo `json:"data"`
	Msg  string    `json:"msg"`
}

type OpenLoginInfo struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Href string `json:"href"` //��ת��ַ
}

type OpenLoginInfoResponse struct {
	Code int             `json:"code"`
	Data []OpenLoginInfo `json:"data"`
	Msg  string          `json:"msg"`
}

type Response struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}
