package nuonuo

import "testing"

func TestRequestBillingNew(t *testing.T) {
	url := "https://sdk.nuonuo.com/open/v1/services"
	senid := ""
	appKey := ""
	appSecret := ""
	token := ""
	taxnum := ""
	method := "nuonuo.ElectronInvoice.requestBillingNew"
	content := ``

	s, err := SendPostSyncRequest(url, senid, appKey, appSecret, token, taxnum, method, content)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}

func TestQueryInvoiceResult(t *testing.T) {
	url := "https://sdk.nuonuo.com/open/v1/services"
	senid := ""
	appKey := ""
	appSecret := ""
	token := ""
	taxnum := ""
	method := "nuonuo.ElectronInvoice.queryInvoiceResult"
	content := `{"orderNos":[""]}`

	s, err := SendPostSyncRequest(url, senid, appKey, appSecret, token, taxnum, method, content)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}
