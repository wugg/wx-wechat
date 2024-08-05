package offiaccount

import (
	"fmt"
	"strconv"
	"time"
	"wechat/tool"
)

var JsApiList = []string{
	"checkJsApi",
	"onMenuShareTimeline",
	"onMenuShareAppMessage",
	"onMenuShareQQ",
	"onMenuShareWeibo",
	"onMenuShareQZone",
	"hideMenuItems",
	"showMenuItems",
	"hideAllNonBaseMenuItem",
	"showAllNonBaseMenuItem",
	"translateVoice",
	"startRecord",
	"stopRecord",
	"onRecordEnd",
	"playVoice",
}

type JsSign struct {
	AppId     string   `json:"appId"`
	NonceStr  string   `json:"nonceStr"`
	Timestamp string   `json:"timestamp"`
	Signature string   `json:"signature"`
	Url       string   `json:"url"`
	JsApiList []string `json:"jsApiList"`
}

func GetJsSign(authUrl string, appId string, nonceStr string, timestamp string) (*JsSign, error) {

	Ticket := WxObj.Ticket
	if Ticket == "" {
		newjst, err := (&JsTicket{}).GetJsTicket()
		if err != nil {
			return nil, err
		}
		Ticket = newjst.Ticket
	}
	if timestamp == "" {
		timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	}
	if nonceStr == "" {
		nonceStr = tool.GetRoundString(16)
	}
	if authUrl == "" {
		err := fmt.Errorf("authUrl is empty")
		return nil, err
	}
	param := map[string]string{
		"appId":     appId,
		"nonceStr":  nonceStr,
		"timestamp": timestamp,
		"url":       authUrl,
	}
	sign := tool.GetSingnature(param)
	return &JsSign{
		AppId:     appId,
		NonceStr:  nonceStr,
		Timestamp: timestamp,
		Signature: sign,
		Url:       authUrl,
		JsApiList: JsApiList,
	}, nil
}
