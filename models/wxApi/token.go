package wxApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/astaxie/beego/orm"
)

var reqUrl = "https://api.weixin.qq.com"

var DB orm.Ormer

func init() {
	orm.RegisterModel(new(UserInfo), new(AccessToken))
	DB = orm.NewOrm()
}

type UserInfo struct {
	Id        int          `orm:"column(id)"`
	WXToken   string       `orm:"column(wx_token)"`
	Access    *AccessToken `orm:"rel(one)"`
	AppId     string       `orm:"column(app_id)"`
	AppSecret string       `orm:"column(app_secret)"`
}

func (u *UserInfo) TableName() string {
	return "user_info"
}

type AccessToken struct {
	Id          int       `orm:"column(id)"`
	Token       string    `json:"access_token" orm:"column(access_token)"`
	ExpiresTime int       `json:"expires_in" orm:"column(exp_time)"`
	ErrCode     int       `json:"errcode" orm:"column(err_code)"`
	ErrMsg      string    `json:"errmsg" orm:"column(err_msg)"`
	CreateTime  string    `json:"time" orm:"column(request_time)"`
	UserId      *UserInfo `orm:"reverse(one)"`
}

func (ac *AccessToken) TableName() string {
	return "access_token"
}
func CreateAccessToken(account string) error {
	user := &UserInfo{WXToken: account}
	err := DB.Read(&user, "wx_token")
	if err == orm.ErrNoRows {
		return errors.New(fmt.Sprintf("user token find error: %s", err.Error()))
	} else {
		fmt.Println(user.WXToken, user.AppId)
	}
	url := reqUrl + fmt.Sprintf("/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", user.AppId, user.AppSecret)
	res, err := http.Get(url)
	if err != nil {
		return errors.New(fmt.Sprintf("http get error: %s", err.Error()))
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("ioutil read response body error: %s", err.Error()))
	}
	respToken := &AccessToken{}
	if err := json.Unmarshal(body, respToken); err != nil {
		return errors.New(fmt.Sprintf("json Unmarshal GetAccessToken error: %s, body: %s", err.Error(), string(body)))
	} else {
		if respToken.Token != "" {
			user.Access = respToken
			DB.Insert(respToken)
			DB.Update(user)
			fmt.Println("new token---", respToken.Token)
			t := time.Duration(respToken.ExpiresTime/2) * time.Second
			expTime := time.Duration(t)
			expires_timer := time.After(expTime)
			<-expires_timer
			CreateAccessToken(account)
		} else {
			fmt.Println(string(body))
			return errors.New("can't get access token !")
		}
	}
	return nil
}
