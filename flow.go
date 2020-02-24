package pay_center_client

import (
	"encoding/json"
	"github.com/FengWuTech/pay-center-client/util/httputil"
	"github.com/FengWuTech/pay-center-client/util/signutil"
)

func (client *PayClient) GetAccountFlow(userID int, page int, pageSize int) *GetAccountFlowResponse {
	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.AddQuery("user_id", userID)
	sign.AddQuery("page", page)
	sign.AddQuery("page_size", pageSize)
	url := sign.GenSignURL(URL_ACCOUNT_FLOW_LIST)

	var response GetAccountFlowResponse
	_, respBody := httputil.Get(url)
	json.Unmarshal(respBody, &response)
	return &response
}
