package offiaccount

import (
	"encoding/json"
	config "wechat/config"
	"wechat/tool"
)

// 模板消息
type WxTemplate struct {
}

type TemplateResponse struct {
	PrimaryIndustry   string                 `json:"primary_industry"`
	SecondaryIndustry string                 `json:"secondary_industry"`
	ErrCode           int                    `json:"errcode"`
	ErrMsg            string                 `json:"errmsg"`
	TemplateId        string                 `json:"template_id"`
	TemplateList      map[string]interface{} `json:"template_list"`
}

func (w *WxTemplate) SetIndustry(data map[string]interface{}) (industry TemplateResponse) {
	ApiUrl := config.API_URL_PREFIX + config.TEMPLATE_SET_INDUSTRY_URL + "access_token=" + AccessToken
	body, _ := tool.HttpPost(ApiUrl, tool.MapToIoReader(data), "application/json")
	json.Unmarshal(body, &industry)
	return industry
}

func (w *WxTemplate) AddTemplate(data map[string]interface{}) (template TemplateResponse) {
	ApiUrl := config.API_URL_PREFIX + config.TEMPLATE_ADD_TPL_URL + "access_token=" + AccessToken
	body, _ := tool.HttpPost(ApiUrl, tool.MapToIoReader(data), "application/json")
	json.Unmarshal(body, &template)
	return template
}

func (w *WxTemplate) GetTemplateList() (template TemplateResponse) {
	ApiUrl := config.API_URL_PREFIX + config.GET_All_PRIVATE_TEMPLATE_URL + "access_token=" + AccessToken
	body, _ := tool.HttpGet(ApiUrl)
	json.Unmarshal(body, &template)
	return template
}

func (w *WxTemplate) DelPrivateTemplate(data map[string]interface{}) (template TemplateResponse) {
	ApiUrl := config.API_URL_PREFIX + config.DEL_PRIVATE_TEMPLATE_URL + "access_token=" + AccessToken
	body, _ := tool.HttpPost(ApiUrl, tool.MapToIoReader(data), "application/json")
	json.Unmarshal(body, &template)
	return template
}

func (w *WxTemplate) SendTemplate(data map[string]interface{}) (template TemplateResponse) {
	ApiUrl := config.API_URL_PREFIX + config.TEMPLATE_SEND_URL + "access_token=" + AccessToken
	body, _ := tool.HttpPost(ApiUrl, tool.MapToIoReader(data), "application/json")
	json.Unmarshal(body, &template)
	return template
}
