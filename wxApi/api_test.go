package wxApi

import (
	"flag"
	"fmt"
	"testing"
)

func TestGetToken(t *testing.T) {
	fmt.Println(flag.Arg(0))
	GetAccessToken(flag.Arg(0))
}

func TestUploadMaterial(t *testing.T) {
	reqUrl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s", Access_token, "image")
	filePath := "/Users/huhai/work/SFTP_SublimeText/golang/src/weixin/resource/upload.jpeg"
	respBytes, err := uploadFile(reqUrl, filePath, "image")
	if err != nil {
		fmt.Println("uploadFile error: ", err.Error())
	}
	fmt.Println(string(respBytes))
}

func TestDownloadMaterial(t *testing.T) {
	reqUrl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s", Access_token, "LKZmaRlooOTTXWCYLqLqXepB-v34pCylrITkh1rjU-cl9lynvAGx7yTNLVt2kH73")
	filePath := "/Users/huhai/work/SFTP_SublimeText/golang/src/weixin/resource/download.jpeg"
	respBytes, err := DownloadFile(reqUrl, filePath)
	if err != nil {
		fmt.Println("downloadFile error: ", err.Error())
	}
	fmt.Println("test")
	fmt.Println(string(respBytes))
}
