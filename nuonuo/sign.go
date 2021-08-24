package nuonuo

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
)

/**
* 计算签名
*
* @param path      请求地址
* @param senid     流水号
* @param appKey    appKey
* @param appSecret appSecret
* @param nonce     随机码
* @param body      请求包体
* @param timestamp    时间戳
* @return 返回签名
 */
func MakeSign(path, appSecret, appKey, senid, nonce, body, timestamp string) string {
	split := strings.Split(path, "/")
	signStr := fmt.Sprintf("a=%s&l=%s&p=%s&k=%s&i=%s&n=%s&t=%s&f=%s",
		split[3], split[2], split[1], appKey, senid, nonce, timestamp, body)
	h := hmac.New(sha1.New, []byte(appSecret))
	h.Write([]byte(signStr))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
