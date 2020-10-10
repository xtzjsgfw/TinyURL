package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SendResult struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func YunPian(mobile string, code string) int {
	apikey := "55149caa594f0de975f86640180c40e8"
	text := fmt.Sprintf("【张文彪test】您的验证码是%s。如非本人操作，请忽略本短信", code)
	url_send_sms := "https://sms.yunpian.com/v1/sms/send.json"
	data_send_sms := url.Values{"apikey": {apikey}, "mobile": {mobile}, "text": {text}}
	//httpsPostForm(url_send_sms, data_send_sms)

	resp, err := http.PostForm(url_send_sms, data_send_sms)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	var sendResult SendResult
	if err := json.Unmarshal(body, &sendResult); err != nil {
		panic(err)
	}
	return sendResult.Code
}

func httpsPostForm(url string, data url.Values) {
	resp, err := http.PostForm(url, data)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}
