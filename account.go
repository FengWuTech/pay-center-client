package pay_center_client

import (
	"encoding/json"
	"github.com/FengWuTech/pay-center-client/util/httputil"
	"github.com/FengWuTech/pay-center-client/util/signutil"
)

func (client *PayClient) CreateAccount(request CreateAccountRequest) *CreateAccountResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(client.Host + URL_ACCOUNT_CREATE)

	var response CreateAccountResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}

func (client *PayClient) EditAccount(request EditAccountRequest) *EditAccountResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(client.Host + URL_ACCOUNT_UPDATE)

	var response EditAccountResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}

func (client *PayClient) GetAccount(id int) *GetAccountResponse {

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.AddQuery("id", id)
	url := sign.GenSignURL(client.Host + URL_ACCOUNT_GET)

	var response GetAccountResponse
	_, respBody := httputil.Get(url)
	json.Unmarshal(respBody, &response)
	return &response
}

func (client *PayClient) GetUserAccount(userID int) *GetUserAccountResponse {

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.AddQuery("user_id", userID)
	url := sign.GenSignURL(client.Host + URL_USER_ACCOUNT)

	var response GetUserAccountResponse
	_, respBody := httputil.Get(url)
	json.Unmarshal(respBody, &response)
	return &response
}
