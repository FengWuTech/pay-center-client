package pay_center_client

import (
	"encoding/json"
	"gitlab.gmtech.top/golang/pay-center-client/util/httputil"
	"gitlab.gmtech.top/golang/pay-center-client/util/signutil"
)

// 微信充值
func (client *PayClient) RechargeWeixinGoPay(request RechargeGoPayRequest) *RechargeGoPayResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(client.Host + URL_RECHARGE_WEIXIN_GOPAY)

	var response RechargeGoPayResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}

// 现金充值
func (client *PayClient) RechargeCashPay(request RechargeCashPayRequest) *RechargeCashPayResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(client.Host + URL_RECHARGE_CASH_PAY)

	var response RechargeCashPayResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}

// 充值历史
func (client *PayClient) RechargeHistory(userID int, page int, pageSize int) *RechargeHistoryResponse {
	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.AddQuery("user_id", userID)
	sign.AddQuery("page", page)
	sign.AddQuery("page_size", pageSize)
	url := sign.GenSignURL(client.Host + URL_RECHARGE_HISTORY)

	var response RechargeHistoryResponse
	_, respBody := httputil.Get(url)
	json.Unmarshal(respBody, &response)
	return &response
}
