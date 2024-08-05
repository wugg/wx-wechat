package offiaccount

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	config2 "wechat/config"
	"wechat/tool"
)

// 微信基础信息获取 access_token  jsticket
type WxAccess struct {
	AccessToken       string `json:"access_token"`
	ExpiresIn         int    `json:"expires_in"`
	ErrCode           int    `json:"errcode"`
	ErrMsg            string `json:"errmsg"`
	RefreshToken      string `json:"refresh_token"`
	OpenId            string `json:"openid"`
	Ticket            string `json:"ticket"`
	AccessExpiresTime int64
}

var AccessToken string
var WxObj = &WxAccess{}
var redisClient = tool.RedisClient

func init() {
	WxObj, _ := WxObj.GetAccessToken()
	AccessToken = WxObj.AccessToken
	ticket, _ := (&JsTicket{}).GetJsTicket()
	WxObj.Ticket = ticket.Ticket
}
func (wx *WxAccess) isAccessExpired() bool {
	return time.Now().Unix() > wx.AccessExpiresTime
}

type JsTicket struct {
	ErrCode           int    `json:"errcode"`
	ErrMsg            string `json:"errmsg"`
	Ticket            string `json:"ticket"`
	ExpiresIn         int    `json:"expires_in"`
	TicketExpiresTime int64
}

func (j *JsTicket) isTicketExpired() bool {
	return time.Now().Unix() > j.TicketExpiresTime
}

func (wx *WxAccess) GetAccessToken() (*WxAccess, error) {
	ctx := context.Background()
	config := &config2.Conf
	if config.AppID == "" || config.AppSecret == "" {
		return nil, fmt.Errorf("appid or appsecret is empty")
	}
	redisv := redisClient.Get(ctx, config.AppID+"_access_token").Val()
	if redisv != "" {
		err := json.Unmarshal([]byte(redisv), wx)
		if err != nil {
			return nil, err
		}
		if !wx.isAccessExpired() {
			return wx, nil
		}
	}
	apiUrl := config2.API_URL_PREFIX + config2.AUTH_URL + "appid=" + config.AppID + "&secret=" + config.AppSecret
	body, err := tool.HttpGet(apiUrl)
	if err != nil {
		return nil, err
	}
	var wxToken = &WxAccess{}
	err = json.Unmarshal(body, wxToken)
	if err != nil {
		return nil, err
	}
	if wxToken.ErrCode != 0 {
		return nil, fmt.Errorf("获取access_token失败,errcode:%d,errmsg:%s", wxToken.ErrCode, wxToken.ErrMsg)
	}
	wxToken.AccessExpiresTime = time.Now().Unix() + int64(wxToken.ExpiresIn)
	jsonwx, _ := json.Marshal(wxToken)
	//redis 缓存access_token 7200s
	biol := redisClient.Set(ctx, config.AppID+"_access_token", jsonwx, time.Duration(int64(wxToken.ExpiresIn))*time.Second)
	if biol.Err() != nil {
		return nil, biol.Err()
	}

	return wxToken, nil
}
func (j *JsTicket) GetJsTicket() (*JsTicket, error) {
	jsTicket := &JsTicket{}
	Ticket := redisClient.Get(context.Background(), config2.Conf.AppID+"_jsapi_ticket").Val()
	if Ticket != "" {
		err := json.Unmarshal([]byte(Ticket), jsTicket)
		if err != nil {
			return nil, err
		}
		if !jsTicket.isTicketExpired() {
			return jsTicket, nil
		}
	}
	apiUrl := config2.API_URL_PREFIX + config2.GET_TICKET_URL + "access_token=" + AccessToken + "&type=jsapi"
	body, err := tool.HttpGet(apiUrl)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, jsTicket)
	if err != nil {
		return nil, err
	}
	if jsTicket.ErrCode != 0 {
		return nil, fmt.Errorf("获取jsapi_ticket失败,errcode:%d,errmsg:%s", jsTicket.ErrCode, jsTicket.ErrMsg)
	}
	jsTicket.TicketExpiresTime = time.Now().Unix() + int64(jsTicket.ExpiresIn)
	jsonJsTicket, _ := json.Marshal(jsTicket)
	redisClient.Set(context.Background(), config2.Conf.AppID+"_jsapi_ticket", jsonJsTicket, time.Duration(int64(jsTicket.ExpiresIn))*time.Second)
	return jsTicket, nil
}
