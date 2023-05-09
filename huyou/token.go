package huyou

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const AUTH_URL = "http://open.hydzfp.com/open_access/access/token.open"

func GetToken(openid, appSecret string) (string, error) {
	if openid == "" {
		return "", errors.New("openid不能为空")
	}
	if appSecret == "" {
		return "", errors.New("AppSecret不能为空")
	}
	tkmap := make(map[string]interface{})
	tkmap["openid"] = openid
	tkmap["app_secret"] = appSecret
	bs, err := json.Marshal(tkmap)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", AUTH_URL,
		bytes.NewReader(bs))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
