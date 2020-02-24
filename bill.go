package pay_center_client

import (
	"encoding/json"
	"github.com/FengWuTech/pay-center-client/util/httputil"
	"github.com/FengWuTech/pay-center-client/util/signutil"
)

func (client *PayClient) BillDeductAuto(request BillDeductAutoRequest) *BillDeductAutoResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(URL_BILL_DEDUCT_AUTO)

	var response BillDeductAutoResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}

// 账单支付
func (client *PayClient) BillWeixinGoPay(request BillGoPayRequest) *BillGoPayResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(URL_BILL_WEIXIN_GOPAY)

	var response BillGoPayResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}

func (client *PayClient) BillCashPay(request BillCashPayRequest) *BillCashPayResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(URL_BILL_CASH_PAY)

	var response BillCashPayResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}
