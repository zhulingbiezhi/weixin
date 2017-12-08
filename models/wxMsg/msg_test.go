package wxMsg

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var xmlStr = `<xml><ToUserName><![CDATA[gh_b49ac5c95640]]></ToUserName>
<FromUserName><![CDATA[opANV1BtEJ7hAJsA-AHrTs5wdZHU]]></FromUserName>
<CreateTime>1512543556</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[给你]]></Content>
<MsgId>6496325107232953614</MsgId>
</xml>`

func TestXML(t *testing.T) {
	msg := new(WeixinMsg)
	if err := xml.Unmarshal([]byte(xmlStr), msg); err != nil {
		fmt.Println("xml Unmarshal to msg error: ", err.Error())
	} else {
		fmt.Print(msg)
		resp, err := xml.Marshal(msg.Reply())
		if err != nil {
			fmt.Println("xml Marshal to resp error: ", err.Error())
		}
		fmt.Println("\n")
		fmt.Println(string(resp))
	}
}
