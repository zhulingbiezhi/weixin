package wxApi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var reqUrl = "https://api.weixin.qq.com"
var app_id_ququ = "wx5a8f22cee992c742"
var app_secret_ququ = "d1060a17e8e56ae8a0fb8535c2198f58"
var app_id_test = "wxe54bf4d47833a264"
var app_secret_test = "eff81765127bf4b2e8b102ab11454353"

var Access_token = "bcuCnjz6O2ntIUyVdEk7G7LF8PIK2_hWm1CHq0zo-U43yeg9-oJNQP11kFR4gkt0ya9qgf84_BRy_tR1PjuufrcEUQrIK3TeSVYweZKCtoNev1YmZKVO9FUHp96Zo37HHJQcAHAUWD"

type respToken struct {
	AccessToken string `json:"access_token"`
	ExpiresTime int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

func GetAccessToken(account string) {
	var url string
	if account == "ququ" {
		url = reqUrl + fmt.Sprintf("/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", app_id_ququ, app_secret_ququ)
	} else if account == "test" {
		url = reqUrl + fmt.Sprintf("/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", app_id_test, app_secret_test)
	} else {
		url = reqUrl + fmt.Sprintf("/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", app_id_test, app_secret_test)
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error: ", err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ioutil read response body error: ", err.Error())
	}
	respToken := &respToken{}
	if err := json.Unmarshal(body, respToken); err != nil {
		fmt.Println("josn Unmarshal GetAccessToken error: ", err.Error())
		fmt.Println(string(body))
	} else {
		fmt.Println(string(body))
		if respToken.AccessToken != "" {
			Access_token = respToken.AccessToken
			fmt.Println("###############")
			fmt.Println("new token---", Access_token)
			fmt.Println("###############")
			//strconv.ParseInt(respMap["expires_in"], 10, 64)
			t := time.Duration(respToken.ExpiresTime/2) * time.Second
			expTime := time.Duration(t)
			fmt.Println(expTime)
			expires_timer := time.After(expTime)
			<-expires_timer
			GetAccessToken("ququ")
		}
	}
}
