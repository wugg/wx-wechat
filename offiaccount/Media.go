package offiaccount

type Media struct {
}

type MeidaResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (m *Media) UploadMedia(mediaType string, filePath string) (mediaResponse *MeidaResponse, err error) {
	//ApiUrl := config.API_BASE_URL_PREFIX + config.MEDIA_UPLOAD_URL + "access_token=" + AccessToken + "&type=" + mediaType
	return mediaResponse, nil
}
