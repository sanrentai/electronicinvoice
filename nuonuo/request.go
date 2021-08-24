package nuonuo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	urlpkg "net/url"
	"strconv"
	"strings"
	"time"
)

func SendPostSyncRequest(url, senid, appKey, appSecret, token, taxnum, method, content string) (string, error) {
	if senid == "" {
		return "", errors.New("senid不能为空")
	}
	if appKey == "" {
		return "", errors.New("appKey不能为空")
	}
	if appSecret == "" {
		return "", errors.New("appSecret不能为空")
	}
	if token == "" {
		return "", errors.New("token不能为空")
	}
	if taxnum == "" {
		return "", errors.New("taxnum不能为空")
	}
	if method == "" {
		return "", errors.New("method不能为空")
	}
	if content == "" {
		return "", errors.New("content不能为空")
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	rand.Seed(time.Now().Unix())
	nonce := strconv.FormatInt(rand.Int63n(1000000), 10)
	sb := fmt.Sprintf("%s?senid=%s&nonce=%s&timestamp=%s&appkey=%s",
		url, senid, nonce, timestamp, appKey)
	req, err := http.NewRequest("POST", sb, strings.NewReader(content))
	if err != nil {
		return "", err
	}
	u, err := urlpkg.ParseRequestURI(url)
	if err != nil {
		return "", err
	}

	xnns := MakeSign(u.RequestURI(), appSecret, appKey, senid, nonce, content, timestamp)
	req.Header["Content-Type"] = []string{"application/json"}
	req.Header["X-Nuonuo-Sign"] = []string{xnns}
	req.Header["accessToken"] = []string{token}
	req.Header["userTax"] = []string{taxnum}
	req.Header["method"] = []string{method}

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
