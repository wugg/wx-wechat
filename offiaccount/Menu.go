package offiaccount

import (
	"context"
	"encoding/json"
	"fmt"
	config2 "wechat/config"
	"wechat/tool"
)

type Menu struct {
	IsMenuOpen   int                    `json:"is_menu_open"`
	SelfmenuInfo map[string]interface{} `json:"selfmenu_info"`
	Errcode      int                    `json:"errcode"`
	Errmsg       string                 `json:"errmsg"`
}

type WxResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// \wechat\offiaccount\Menu.go  button_type 菜单类型
// click: 点击推事件  view: 跳转URL  scancode_push: 扫码推事件  scancode_waitmsg: 扫码推事件且弹出“消息接收中”提示框
// pic_sysphoto: 弹出系统拍照发图  pic_photo_or_album: 弹出拍照或者相册发图  pic_weixin: 弹出微信相册发图器  location_select: 弹出地理位置选择器
// miniprogram: 小程序类型 media_id: 图片  view_limited: 图文消息  article_id: 发布后的图文消息  article_view_limited: 发布后的图文消息
var button_type = []string{"click", "view", "scancode_push", "scancode_waitmsg", "pic_sysphoto", "pic_photo_or_album",
	"pic_weixin", "location_select", "miniprogram", "media_id", "view_limited", "article_id", "article_view_limited"}

func (m *Menu) CreateMenu(menu map[string]interface{}) (err error) {
	Apiurl := config2.API_URL_PREFIX + config2.MENU_CREATE_URL + "access_token=" + AccessToken
	body, _ := tool.HttpPost(Apiurl, tool.MapToIoReader(menu), "application/json")
	var response WxResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("json.Unmarshal(body, &response) error(%v)", err)
	}
	if response.Errcode != 0 {
		return fmt.Errorf("CreateMenu error(%v)", response.Errmsg)
	}
	return nil
}

func (m *Menu) GetMenu() (response Menu, err error) {
	var ctx context.Context
	var body []byte
	resp := redisClient.Get(ctx, "menu").Val()
	if resp == "" {
		Apiurl := config2.API_URL_PREFIX + config2.MENU_GET_URL + "access_token=" + AccessToken
		body, _ = tool.HttpGet(Apiurl)
		redisClient.Set(ctx, "menu", body, 0)
		if response.Errcode != 0 {
			return response, fmt.Errorf("GetMenu error(%v)", response.Errmsg)
		}
	} else {
		body = []byte(resp)
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf("json.Unmarshal(body, &response) error(%v)", err)
	}
	return response, nil
}

func (m *Menu) DeleteMenu() (err error) {
	ApiUrl := config2.API_URL_PREFIX + config2.MENU_DELETE_URL + "access_token=" + AccessToken
	body, _ := tool.HttpGet(ApiUrl)
	var response WxResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	if response.Errcode != 0 {
		return fmt.Errorf("DeleteMenu error(%v)", response.Errmsg)
	}
	return nil
}
