package open_login

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type QQInfo struct {
	Nickname    string `json:"nickname"`
	Gender      string `json:"gender"`
	FigureurlQQ string `json:"figureurl_qq"` //头像大图
	OpenID      string `json:"open_id"`
}
type QQLogin struct {
	appID     string
	appKey    string
	redirect  string
	code      string
	accessTok string
	openID    string
}
type QQConfig struct {
	AppID    string
	AppKey   string
	Redirect string
}

func NewQQLogin(code string, QQconfig QQConfig) (qqInfo QQInfo, err error) {
	qqLogin := &QQLogin{
		appID:    QQconfig.AppID,
		appKey:   QQconfig.AppKey,
		redirect: QQconfig.Redirect,
		code:     code,
	}
	err = qqLogin.GetAccessToken()
	if err != nil {
		fmt.Println(err)
		return qqInfo, err
	}
	err = qqLogin.GetOpenID()
	if err != nil {
		return qqInfo, err
	}
	qqInfo, err = qqLogin.GetUserInfo()
	if err != nil {
		return qqInfo, err
	}
	qqInfo.OpenID = qqLogin.openID
	return qqInfo, nil
}
func (q *QQLogin) GetAccessToken() error {
	//获取access_token
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", q.appID)
	params.Add("client_secret", q.appKey)
	params.Add("code", q.code)
	params.Add("redirect_uri", q.redirect)
	u, err := url.Parse("https://graph.qq.com/oauth2.0/token")
	if err != nil {
		return err
	}
	u.RawQuery = params.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()
	qs, err := parseQS(res.Body)
	if err != nil {
		return err
	}
	q.accessTok = qs[`access_token`][0]

	return nil

}
func (q *QQLogin) GetOpenID() error {
	u, err := url.Parse(fmt.Sprintf("https://graph.qq.com/oauth2.0/me?access_token=%s", q.accessTok))

	if err != nil {
		return err
	}
	res, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()
	openID, err := getOpenID(res.Body)
	if err != nil {
		return err
	}
	q.openID = openID
	return nil

}
func (q *QQLogin) GetUserInfo() (qqinfo QQInfo, err error) {
	params := url.Values{}
	params.Add("access_token", q.accessTok)
	params.Add("oauth_consumer_key", q.appID)
	params.Add("openid", q.openID)
	u, err := url.Parse("https://graph.qq.com/user/get_user_info")
	if err != nil {
		return qqinfo, err
	}
	u.RawQuery = params.Encode()
	res, err := http.Get(u.String())
	data, err := io.ReadAll(res.Body)
	err = json.Unmarshal(data, &qqinfo)
	if err != nil {
		return qqinfo, err
	}
	return qqinfo, nil
}

// parseQS 将HTTP响应的正文解析为键值对的形式
func parseQS(r io.Reader) (val map[string][]string, err error) {
	val, err = url.ParseQuery(readAll(r))
	if err != nil {
		return val, err
	}
	return val, nil
}

// 将HTTP响应的正文解析出openid
func getOpenID(r io.Reader) (string, error) {
	body := readAll(r)
	start := strings.Index(body, `"openid":"`) + len(`"openid":"`)
	if start == -1 {
		fmt.Errorf("openid not found")
	}
	end := strings.Index(body[start:], `"`)
	if end == -1 {
		fmt.Errorf("openid not found")

	}
	return body[start : start+end], nil
}

// readAll读取所有数据并将其转换为字符串
func readAll(r io.Reader) string {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
