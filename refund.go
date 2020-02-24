package pay_center_client

import (
	"encoding/json"
	"github.com/FengWuTech/pay-center-client/util/httputil"
	"github.com/FengWuTech/pay-center-client/util/signutil"
)

func (client *PayClient) RefundToAccount(request RefundToAccountRequest) *RefundToAccountResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(URL_REFUND_TO_ACCOUNT)

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
	url := sign.GenSignURL(URL_REFUND_TO_USER)

	var response RefundToUserResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}
