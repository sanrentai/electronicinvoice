package nuonuo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const AUTH_URL = "https://open.nuonuo.com/accessToken"

func GetToken(appkey, appSecret string) (string, error) {
	if appkey == "" {
		return "", errors.New("AppKey不能为空")
	}
	if appSecret == "" {
		return "", errors.New("AppSecret不能为空")
	}
	req, err := http.NewRequest("POST", AUTH_URL,
		strings.NewReader(fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=client_credentials",
			appkey, appSecret)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
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
