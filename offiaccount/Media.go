package offiaccount

import (
	"encoding/json"
	"wechat/config"
	"wechat/tool"
)

type Media struct {
}

type MeidaResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	VideoUrl   string `json:"video_url"`
	MediaId    string `json:"media_id"`
	Url        string `json:"url"`
	VoiceCount int    `json:"voice_count"`
	VideoCount int    `json:"video_count"`
	ImageCount int    `json:"image_count"`
	NewsCount  int    `json:"news_count"`
	TotalCount int    `json:"total_count"`
	ItemCount  int    `json:"item_count"`
	Item       []struct {
	} `json:"item"`
}

/**
 * 上传多媒体文件(认证后的订阅号可用)
 * 注意：上传大文件时可能需要先调用 set_time_limit(0) 避免超时
 * 注意：数组的键值任意，但文件名前必须加@，使用单引号以避免本地路径斜杠被转义
 * @param array $data {"media":'@Path\filename.jpg'}
 * @param type 类型：图片:image 语音:voice 视频:video 缩略图:thumb
 * @return boolean|array
 */
func (m *Media) UploadMedia(mediaType string, data map[string]interface{}) (mediaResponse *MeidaResponse) {
	ApiUrl := config.API_URL_PREFIX + config.MEDIA_UPLOAD_URL + "access_token=" + AccessToken + "&type=" + mediaType
	Response, _ := tool.HttpPost(ApiUrl, tool.MapToIoReader(data), "application/json")

	json.Unmarshal(Response, &mediaResponse)
	return mediaResponse
}

func (m *Media) GetMedia(mediaId string) (mediaResponse *MeidaResponse) {
	ApiUrl := config.API_URL_PREFIX + config.MEDIA_GET_URL + "access_token=" + AccessToken + "&media_id=" + mediaId
	body, _ := tool.HttpGet(ApiUrl)
	json.Unmarshal(body, &mediaResponse)
	return mediaResponse
}

func (m *Media) UploadImg(data map[string]interface{}) (response *MeidaResponse) {
	ApiUrl := config.API_URL_PREFIX + config.MEDIA_IMG_UPLOAD + "access_token=" + AccessToken
	body, _ := tool.HttpPost(ApiUrl, tool.MapToIoReader(data), "application/json")
	json.Unmarshal(body, &response)
	return response
}

// 添加其他素材
func (m *Media) UploadOther(data map[string]interface{}) (response *MeidaResponse) {
	ApiUrl := config.API_URL_PREFIX + config.ADD_MATERIAL_URL + "access_token=" + AccessToken
	body, _ := tool.HttpPost(ApiUrl, tool.MapToIoReader(data), "application/json")
	json.Unmarshal(body, &response)
	return response
}

func (m *Media) UploadVideo(data map[string]interface{}) (response *MeidaResponse) {
	ApiUrl := config.API_URL_PREFIX + config.MEDIA_VIDEO_UPLOAD + "access_token=" + AccessToken
	body, _ := tool.HttpPost(ApiUrl, tool.MapToIoReader(data), "application/json")
	json.Unmarshal(body, &response)
	return response
}

// 获取永久素材
func (m *Media) GetOther(data map[string]interface{}) (response *MeidaResponse) {
	ApiUrl := config.API_URL_PREFIX + config.MEDIA_GET_MATERIAL + "access_token=" + AccessToken
	body, _ := tool.HttpPost(ApiUrl, tool.MapToIoReader(data), "application/json")
	json.Unmarshal(body, &response)
	return response
}

// 删除素材
func (m *Media) DeleteMaterial(data map[string]interface{}) (response *MeidaResponse) {
	ApiUrl := config.API_URL_PREFIX + config.MEDIA_DEL_MATERIAL + "access_token=" + AccessToken
	body, _ := tool.HttpPost(ApiUrl, tool.MapToIoReader(data), "application/json")
	json.Unmarshal(body, &response)
	return response
}

// 获取素材总数
func (m *Media) GetMaterialCount() (response *MeidaResponse) {
	ApiUrl := config.API_URL_PREFIX + config.MEDIA_GET_MATERIAL_COUNT + "access_token=" + AccessToken
	body, _ := tool.HttpGet(ApiUrl)
	json.Unmarshal(body, &response)
	return response
}

func (m *Media) GetMaterialList(data map[string]interface{}) (response *MeidaResponse) {
	ApiUrl := config.API_URL_PREFIX + config.MEDIA_GET_MATERIAL_LIST + "access_token=" + AccessToken
	body, _ := tool.HttpPost(ApiUrl, tool.MapToIoReader(data), "application/json")
	json.Unmarshal(body, &response)
	return response
}
