package qywxpush

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//  参考文档 : https://developer.work.weixin.qq.com/document/path/91770

type data struct {
	MsgType  string           `json:"msgtype"`
	Markdown MarkdownTemplate `json:"markdown"`
}
type MarkdownTemplate struct {
	Content string `json:"content"`
}

type push struct {
	Code int    `json:"code"`
	Msg  string `json:"Msg"`
}

func qywxpush(webhookurl string, str string) error {
	url := webhookurl
	markdown := MarkdownTemplate{str}
	data := data{"markdown", markdown}
	s, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	//client := &http.Client{}
	d := strings.NewReader(string(s))
	// 忽略 https 证书验证
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	req, err := http.NewRequest("POST", url, d)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var push push
	err = json.Unmarshal(bodyText, &push)
	if err != nil {
		fmt.Println(err)
	}
	if push.Code == 200 {
		return nil
	} else {
		return errors.New(push.Msg)
	}
}
