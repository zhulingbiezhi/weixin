package main

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"weixin/wxMsg"
)

func main() {
	http.HandleFunc("/wx", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Println("GET Method")
			w.Header().Add("Context-Type", "text/plain; charset=utf-8")
			token := handleToken(w, r)
			fmt.Println(token)
			w.Write([]byte(token))
		} else if r.Method == "POST" {
			fmt.Println("POST Method")
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println("read requestBody error: ", err.Error())
			} else {
				fmt.Print(r.Header)
				fmt.Println("\n", string(data))
				msg := new(wxMsg.WeixinMsg)
				if err := xml.Unmarshal(data, msg); err != nil {
					fmt.Println("xml Unmarshal to msg error: ", err.Error())
				} else {
					resp, err := xml.Marshal(msg.Reply())
					if err != nil {
						fmt.Println("xml Marshal response error: ", err.Error())
					} else {
						w.Header().Add("Context-Type", "application/xml; charset=utf-8")
						fmt.Println("the response body: ", string(resp))
						w.Write(resp)
					}
				}
			}
		} else {
			fmt.Println("unkown method:", r.Method)
		}
	})
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("ListenAndServe error:", err.Error())
	}
}

func handleToken(w http.ResponseWriter, r *http.Request) string {
	r.ParseForm()
	form := r.Form
	fmt.Print(form)
	dataMap := make(map[string]string)
	dataMap["nonce"] = form.Get("nonce")
	dataMap["signature"] = form.Get("signature")
	dataMap["echo"] = form.Get("echostr")
	dataMap["timestamp"] = form.Get("timestamp")
	dataMap["token"] = "ququ"
	var keys []string
	keys = append(keys, dataMap["nonce"])
	keys = append(keys, dataMap["token"])
	keys = append(keys, dataMap["timestamp"])
	sort.Strings(keys)
	h := sha1.New()
	for _, v := range keys {
		fmt.Println(v)
		h.Write([]byte(v))
	}
	strMd5 := fmt.Sprintf("%x", h.Sum(nil))
	if strMd5 == dataMap["signature"] {
		return dataMap["echo"]
	} else {
		fmt.Println("signature error")
		return ""
	}
}
