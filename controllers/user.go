package controllers

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"sort"
)

type UserToken struct {
	//	ErrorCode int    `form:"errorcode"`
	//	ErrorMsg  string `form:"errmsg"`
	Nonce     string `form:"nonce"`
	TimeStamp string `form:"timestamp"`
	Signature string `form:"signature"`
	EchoStr   string `form:"echostr"`
	Token     string
}

func (token *UserToken) Verify() error {
	if token.Token != "" {
		var keys []string
		keys = append(keys, token.Nonce)
		keys = append(keys, token.TimeStamp)
		keys = append(keys, token.Token)
		sort.Strings(keys)

		h := sha1.New()
		for _, v := range keys {
			h.Write([]byte(v))
		}
		strMd5 := fmt.Sprintf("%x", h.Sum(nil))
		if strMd5 == token.Signature {
			fmt.Println("token verify success !")
			return nil
		} else {
			fmt.Println("token verify error, the body: ", fmt.Sprint(token))
			return errors.New("token verify error !")
		}
	} else {
		return errors.New("the token key is empty !")
	}
}
