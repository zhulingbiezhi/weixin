package wxApi

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"sort"
	"weixin/models"
)

type UserToken struct {
	Nonce     string `form:"nonce"`
	TimeStamp string `form:"timestamp"`
	Signature string `form:"signature"`
	EchoStr   string `form:"echostr"`
}

func (token *UserToken) Verify() error {
	var keys []string
	keys = append(keys, token.Nonce)
	keys = append(keys, token.TimeStamp)
	tokens := QueryAllUser()
	fmt.Println(tokens)
	for _, t := range tokens {
		newkeys := append(keys, t)
		sort.Strings(newkeys)
		h := sha1.New()
		for _, v := range newkeys {
			h.Write([]byte(v))
		}
		strMd5 := fmt.Sprintf("%x", h.Sum(nil))
		if strMd5 == token.Signature {
			fmt.Println("token verify success !")
			return nil
		}
	}

	fmt.Println("token verify error, the body: ", fmt.Sprint(token))
	return errors.New("token verify error !")
}

func QueryAllUser() []string {
	sql := "select wx_token from user_info"
	users := make([]UserInfo, 5)
	var userToken []string
	_, err := models.DB.Raw(sql).QueryRows(&users)
	if err != nil {
		fmt.Println("find user error: ", err.Error())
	} else {
		for _, u := range users {
			userToken = append(userToken, u.WXToken)
		}
	}
	return userToken
}
