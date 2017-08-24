package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"wefacer/core"
)

type AutoGenerated struct {
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int    `json:"expires_in"`
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	SessionSecret string `json:"session_secret"`
}

type FaceAddKey struct {
	APIKey    string
	APISecret string
}

var (
	FaceAddKeyValue    FaceAddKey
	AutoGeneratedValue AutoGenerated
)

//Get BaiduToken
func InitBaiduToken() bool {
	res, err := http.PostForm(core.WefacerConfig.ConfigMap["baidu_token_url"], url.Values{"grant_type": {"client_credentials"}, "client_id": {core.WefacerConfig.ConfigMap["baidu_api_key"]}, "client_secret": {core.WefacerConfig.ConfigMap["baidu_secret_key"]}})
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	err_unmarshal := json.Unmarshal(body, &AutoGeneratedValue)
	if err_unmarshal != nil {
		log.Println(err_unmarshal.Error())
		return false
	}
	log.Println("baidu token get success")
	return true
}
func InitFaceAddToken() {
	FaceAddKeyValue = FaceAddKey{core.WefacerConfig.ConfigMap["faceadd_api_key"], core.WefacerConfig.ConfigMap["faceadd_secret_secret_key"]}
}

func init() {
	InitBaiduToken()
	InitFaceAddToken()
}
