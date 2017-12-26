package controllers

import (
	"encoding/xml"
	"fmt"

	"weixin/models/wxApi"
	"weixin/models/wxMsg"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	token := new(wxApi.UserToken)
	err := c.ParseForm(token)
	if err != nil {
		fmt.Println("ParseForm error")
		c.Ctx.WriteString("fail")
	} else {
		err := token.Verify()
		if err != nil {
			fmt.Println("token verify error: ", err.Error())
		} else {
			c.Ctx.WriteString(token.EchoStr)
		}
	}
}

func (c *MainController) Post() {
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
