package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	request_url := "http://localhost:7000/"
	// 要 POST的 参数
	form := url.Values{
		"username": {"xiaoming"},
		"address":  {"beijing"},
		"subject":  {"Hello"},
		"from":     {"china"},
	}

	// func Post(url string, bodyType string, body io.Reader) (resp *Response, err error) {
	//对form进行编码
	body := bytes.NewBufferString(form.Encode())
	rsp, err := http.Post(request_url, "application/x-www-form-urlencoded", body)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	body_byte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body_byte))
}
