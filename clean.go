package pay_center_client

import (
	"encoding/json"
	"github.com/FengWuTech/pay-center-client/util/httputil"
	"github.com/FengWuTech/pay-center-client/util/signutil"
)

// 账单支付
func (client *PayClient) CleanCompanyData(companyID int) *Response {
	var sendBody, _ = json.Marshal(map[string]interface{}{
		"company_id": companyID,
	})

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(client.Host + URL_CLEAN_COMPANY_DATA)

	var response Response
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}
