package pay_center_client

import (
	"encoding/json"
	"gitlab.gmtech.top/golang/pay-center-client/util/httputil"
	"gitlab.gmtech.top/golang/pay-center-client/util/signutil"
)

func (client *PayClient) RefundToAccount(request RefundToAccountRequest) *RefundToAccountResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(client.Host + URL_REFUND_TO_ACCOUNT)

	var response RefundToAccountResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}

func (client *PayClient) RefundToUser(request RefundToUserRequest) *RefundToUserResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(client.Host + URL_REFUND_TO_USER)

	var response RefundToUserResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}
