package controllers

import (
	"encoding/xml"
	"fmt"

	"weixin/models/wxMsg"

	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

var strToken = "test"

func (c *TestController) Get() {
	token := new(UserToken)
	err := c.ParseForm(token)
	if err != nil {
		fmt.Println("ParseForm error")
		c.Ctx.WriteString("fail")
	} else {
		token.Token = strToken
		err := token.Verify()
		if err != nil {
			fmt.Println("token verify error: ", err.Error())
		} else {
			c.Ctx.WriteString(token.EchoStr)
		}
	}
}

func (c *TestController) Post() {
	data := c.Ctx.Input.RequestBody
	fmt.Print(c.Ctx.Input.Context.Request.Header)
	fmt.Println("\n", string(data))
	msg := new(wxMsg.WeixinMsg)
	if err := xml.Unmarshal(data, msg); err != nil {
		fmt.Println("xml Unmarshal to msg error: ", err.Error())
	} else {
		resp, err := xml.Marshal(msg.Reply())
		if err != nil {
			fmt.Println("xml Marshal response error: ", err.Error())
		} else {
			c.Ctx.Output.Header("Context-Type", "application/xml; charset=utf-8")
			fmt.Println("the response body: ", string(resp))
			c.Ctx.WriteString(string(resp))
		}
	}

}
