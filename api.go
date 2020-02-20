package pay_center_client

import (
	"encoding/json"
	"errors"
	"github.com/FengWuTech/pay-center-client/util/httputil"
	"github.com/FengWuTech/pay-center-client/util/signutil"
	"io/ioutil"
	"net/http"
)

const (
	URL_CASH_PREPAY    = "http://pms-api.gmtech.top/pay-center/minip-user/recharge/cash/pay"
	URL_RECHARGE_GOPAY = "http://pms-api.gmtech.top/pay-center/minip-user/recharge/gopay"
	URL_BILL_GOPAY     = "http://pms-api.gmtech.top/pay-center/minip-user/billpay/gopay"
	URL_ACCOUNT_CREATE = "http://pms-api.gmtech.top/pay-center/minip-user/account/add"
	URL_ACCOUNT_UPDATE = "http://pms-api.gmtech.top/pay-center/minip-user/account/update"
	URL_ACCOUNT_GET    = "http://pms-api.gmtech.top/pay-center/minip-user/account/get"
)

type PayClient struct {
	AppID  string
	ApiKey string
}

// 实例化请求端
func NewPayClient(appID string, apiKey string) *PayClient {
	var client PayClient
	client.AppID = appID
	client.ApiKey = apiKey
	return &client
}

// 充值
func (client *PayClient) RechargeGoPay(request RechargeGoPayRequest) *RechargeGoPayResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(URL_RECHARGE_GOPAY)

	var response RechargeGoPayResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}

// 账单支付
func (client *PayClient) BillGoPay(request BillGoPayRequest) *BillGoPayResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(URL_BILL_GOPAY)

	var response BillGoPayResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}

type NotifyRequestFunc func(request NotifyRequest)

// 支付回调请求分析
func (client *PayClient) PayNotify(httpReq *http.Request) (error, *NotifyRequest) {
	reqBody, err := ioutil.ReadAll(httpReq.Body)
	if err != nil {
		return errors.New(err.Error()), nil
	}

	sign := signutil.NewSign(client.ApiKey)
	isMatch := sign.VerifySign(httpReq.RequestURI, string(reqBody))
	if isMatch {
		var notify NotifyRequest
		json.Unmarshal(reqBody, &notify)
		return nil, &notify
	} else {
		return errors.New("签名不一致"), nil
	}
}

func (client *PayClient) CashRecharge(request CashRechargeRequest) *CashRechargeResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(URL_CASH_PREPAY)

	var response CashRechargeResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}

func (client *PayClient) CreateAccount(request CreateAccountRequest) *CreateAccountResponse {
	var sendBody, _ = json.Marshal(request)

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(URL_ACCOUNT_CREATE)

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
	url := sign.GenSignURL(URL_ACCOUNT_UPDATE)

	var response EditAccountResponse
	_, respBody := httputil.PostRawJson(url, string(sendBody))
	json.Unmarshal(respBody, &response)
	return &response
}

func (client *PayClient) GetAccount(id int) *GetAccountResponse {

	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.AddQuery("id", id)
	url := sign.GenSignURL(URL_ACCOUNT_GET)

	var response GetAccountResponse
	_, respBody := httputil.Get(url)
	json.Unmarshal(respBody, &response)
	return &response
}
