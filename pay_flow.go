package pay_center_client

import (
	"encoding/json"
	"github.com/FengWuTech/pay-center-client/util/httputil"
	"github.com/FengWuTech/pay-center-client/util/signutil"
)

func (client *PayClient) GetPayFlow(id int) *GetPayFlowResponse {

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.AddQuery("id", id)
	url := sign.GenSignURL(client.Host + URL_PAY_FLOW_GET)

	var response GetPayFlowResponse
	_, respBody := httputil.Get(url)
	json.Unmarshal(respBody, &response)
	return &response
}
