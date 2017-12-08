package wxMsg

import (
	"encoding/xml"
	"fmt"
	"time"
)

type WeixinMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content,omitempty"`
	MsgId        string   `xml:"MsgId,omitempty"`
	PicUrl       string   `xml:"PicUrl,omitempty"`
	MediaId      string   `xml:"MediaId,omitempty"`
	Format       string   `xml:"Format,omitempty"`
	Recognition  string   `xml:"Recognition,omitempty"`
	ThumbMediaId string   `xml:"ThumbMediaId,omitempty"`
	Location_X   string   `xml:"Location_X,omitempty"`
	Location_Y   string   `xml:"Location_Y,omitempty"`
	Scale        string   `xml:"Scale,omitempty"`
	Title        string   `xml:"Title,omitempty"`
	Description  string   `xml:"Description,omitempty"`
	Url          string   `xml:"Url,omitempty"`
	ArticleCount string   `xml:"ArticleCount,omitempty"`
	Event        string   `xml:"Event,omitempty"`
	EventKey     string   `xml:"EventKey,omitempty"`
	WX_image     *wxImage `xml:"Image,omitempty"`
	WX_voice     *wxVoice `xml:"Voice,omitempty"`
	WX_video     *wxVideo `xml:"Video,omitempty"`
	WX_music     *wxMusic `xml:"Music,omitempty"`
	WX_news      *wxNews  `xml:"Articles,omitempty"`
}
type wxImage struct {
	XMLName xml.Name `xml:"Image"`
	MediaId string   `xml:"MediaId"`
}
type wxVoice struct {
	XMLName xml.Name `xml:"Voice"`
	MediaId string   `xml:"MediaId"`
}
type wxVideo struct {
	XMLName     xml.Name `xml:"Video"`
	MediaId     string   `xml:"MediaId"`
	Title       string   `xml:"Title"`
	Description string   `xml:"Description"`
}
type wxMusic struct {
	XMLName      xml.Name `xml:"Music"`
	MusicUrl     string   `xml:"MusicUrl"`
	HQMusicUrl   string   `xml:"HQMusicUrl"`
	ThumbMediaId string   `xml:"ThumbMediaId"`
	Title        string   `xml:"Title"`
	Description  string   `xml:"Description"`
}
type wxNews struct {
	XMLName xml.Name `xml:"Articles"`
	Items   []struct {
		XMLName     xml.Name `xml:"item"`
		Title       string   `xml:"Title"`
		PicUrl      string   `xml:"PicUrl"`
		Description string   `xml:"Description"`
		Url         string   `xml:"Url"`
	} `xml:"item"`
}

func (msg *WeixinMsg) Reply() WeixinMsg {
	//var strReply string
	switch msg.MsgType {
	case "text":
		return msg.dealWithText()
	case "image":
		return msg.dealWithImage()
	case "voice":
		return msg.dealWithVoice()
	case "video":
	case "shortvideo":
		return msg.dealWithVideo()
	case "location":
		return msg.dealWithLocation()
	case "link":
		return msg.dealWithLink()
	case "event":
		return msg.dealWithEvent()
	default:
	}
	return WeixinMsg{}
}

func (msg *WeixinMsg) dealWithText() WeixinMsg {
	var newMsg WeixinMsg
	newMsg.FromUserName = msg.ToUserName
	newMsg.ToUserName = msg.FromUserName
	newMsg.MsgType = "text"
	newMsg.CreateTime = time.Now().String()
	newMsg.Content = fmt.Sprintf("this is the type : %s,the context is : %s", msg.MsgType, msg.Content)
	return newMsg
}
func (msg *WeixinMsg) dealWithImage() WeixinMsg {
	var newMsg WeixinMsg
	newMsg.FromUserName = msg.ToUserName
	newMsg.ToUserName = msg.FromUserName
	newMsg.MsgType = "image"
	newMsg.CreateTime = time.Now().String()
	newMsg.WX_image = new(wxImage)
	newMsg.WX_image.MediaId = "LKZmaRlooOTTXWCYLqLqXfOiPFdcJL5XMkYYb984KxPjhPg45TPZuCm46NFQUanb"
	return newMsg
}
func (msg *WeixinMsg) dealWithVoice() WeixinMsg {
	return defaultMsg(msg)
}
func (msg *WeixinMsg) dealWithVideo() WeixinMsg {
	return defaultMsg(msg)
}
func (msg *WeixinMsg) dealWithShortVideo() WeixinMsg {
	return defaultMsg(msg)
}
func (msg *WeixinMsg) dealWithLocation() WeixinMsg {
	return defaultMsg(msg)
}
func (msg *WeixinMsg) dealWithLink() WeixinMsg {
	return defaultMsg(msg)
}
func (msg *WeixinMsg) dealWithEvent() WeixinMsg {
	newMsg := WeixinMsg{}
	if msg.Event == "subscribe" {
		newMsg.FromUserName = msg.ToUserName
		newMsg.ToUserName = msg.FromUserName
		newMsg.MsgType = "text"
		newMsg.CreateTime = time.Now().String()
		newMsg.Content = "谢谢宝宝的关注哦，么么哒"
	}
	return newMsg
}

func defaultMsg(msg *WeixinMsg) WeixinMsg {
	newMsg := WeixinMsg{}
	newMsg.FromUserName = msg.ToUserName
	newMsg.ToUserName = msg.FromUserName
	newMsg.MsgType = "text"
	newMsg.CreateTime = time.Now().String()
	newMsg.Content = "不支持的类型哦，么么哒"
	return newMsg
}
