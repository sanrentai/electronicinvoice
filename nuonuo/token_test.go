package nuonuo

import "testing"

func TestGetToken(t *testing.T) {
	appkey := ""
	appSecret := ""
	got, err := GetToken(appkey, appSecret)
	if err != nil {
		t.Errorf("GetToken() error = %v, ", err)
		return
	}
	t.Log(got)
}
