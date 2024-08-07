package tool

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   100,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			DisableKeepAlives:     false,
			DisableCompression:    false,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		},
	}
}

func HttpGet(url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
func HttpPost(url string, body io.Reader, contentType string) ([]byte, error) {
	if contentType == "" {
		contentType = "application/json"
	}
	resp, err := client.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func CreateXmlBody(data map[string]interface{}) string {
	buffer := bytes.NewBufferString("")
	buffer.WriteString("<xml>")
	for key, value := range data {
		if value == nil {
			continue
		}
		buffer.WriteString(fmt.Sprintf("<%s><![CDATA[%v]]></%s>", key, value, key))
	}
	buffer.WriteString("</xml>")
	return buffer.String()
}

func GetSingnature(data map[string]string) (signature string) {
	if data == nil {
		return
	}
	haser := sha1.New()
	urlStr := ""
	for key, value := range data {
		url.JoinPath(urlStr, fmt.Sprintf("%s=%s", key, value))
	}
	haser.Write([]byte(urlStr))
	return hex.EncodeToString(haser.Sum(nil))
}

func GetRoundString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var b [64]byte
	for i := 0; i < length; i++ {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b[:length])
}

func MapToIoReader(data map[string]interface{}) io.Reader {
	if data == nil {
		return nil
	}
	Jsondata, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	buffer := bytes.NewReader(Jsondata)
	return buffer
}
