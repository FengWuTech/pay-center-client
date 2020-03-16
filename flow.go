package pay_center_client

import (
	"encoding/json"
	"github.com/FengWuTech/pay-center-client/util/httputil"
	"github.com/FengWuTech/pay-center-client/util/signutil"
)

func (client *PayClient) GetAccountFlow(request GetAccountFlowRequest) *GetAccountFlowResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(client.Host + URL_ACCOUNT_FLOW_LIST)

	var response GetAccountFlowResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}
