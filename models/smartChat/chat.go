package smartChat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"weixin/models"
)

type AppInfo struct {
	Id        int    `orm:"column(id)"`
	AppKey    string `orm:"column(smart_app_key)"`
	AppCode   string `orm:"column(smart_app_code)"`
	AppSecret string `orm:"column(smart_app_secret)"`
}

type smartReply struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Result struct {
		Type        string `json:"type"`
		Content     string `json:"content"`
		Relquestion string `json:"relquestion"`
	} `json:"result"`
}

var appCode = ""

func init() {
	//fmt.Println("REGISTER AppInfo")
	models.RegisterModel(new(AppInfo))
}

func Speak(question string) (string, error) {
	v := url.Values{
		"question": {question},
	}
	fmt.Println(v.Encode())
	reqest, err := http.NewRequest("GET", fmt.Sprintf("http://jisuznwd.market.alicloudapi.com/iqa/query?%s", v.Encode()), nil)
	if err != nil {
		return "", fmt.Errorf("http.NewRequest error: %s", err.Error())
	}
	reqest.Header.Add("Authorization", fmt.Sprintf("APPCODE %s", getAppCode()))
	resp, err := http.DefaultClient.Do(reqest)
	if err != nil {
		return "", fmt.Errorf("DefaultClient request error: %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll error: %s", err.Error())
	}
	smartResp := new(smartReply)
	fmt.Println(string(body))
	err = json.Unmarshal(body, smartResp)
	if err != nil {
		return "", fmt.Errorf("json.Unmarshal error: %s", err.Error())
	}
	return smartResp.Result.Content, nil
}

func getAppCode() string {
	if true {
		sql := "select * from app_info"

		infos := make([]AppInfo, 10)
		_, err := models.DB.Raw(sql).QueryRows(&infos)
		if err != nil {
			fmt.Println("find app_code error: ", err.Error())
		} else {
			fmt.Println(infos)
			appCode = infos[0].AppCode
		}
	}
	//fmt.Println("appcode -----", appCode)
	return appCode
}
