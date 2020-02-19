package pay_center_client

import (
	"encoding/json"
	"errors"
	"github.com/FengWuTech/pay-center-client/util"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"net/http"
)

const (
	URL_RECHARGE_GOPAY = "http://pms-api.gmtech.top/pay-center/minip-user/recharge/gopay"
	RUL_BILL_GOPAY     = "http://pms-api.gmtech.top/pay-center/minip-user/billpay/gopay"
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
	sign := util.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(URL_RECHARGE_GOPAY)

	var response RechargeGoPayResponse
	gorequest.New().Post(url).
		AppendHeader("Content-Type", "application/json").
		SendStruct(request).
		EndStruct(&response)
	return &response
}

// 账单支付
func (client *PayClient) BillGoPay(apiKey string, request BillGoPayRequest) *BillGoPayResponse {
	var sendBody, _ = json.Marshal(request)
	sign := util.NewSign(apiKey)
	sign.SetBody(string(sendBody))
	url := sign.GenSignURL(URL_RECHARGE_GOPAY)

	var response BillGoPayResponse
	gorequest.New().Post(url).
		AppendHeader("Content-Type", "application/json").
		SendStruct(request).
		EndStruct(&response)
	return &response
}

type NotifyRequestFunc func(request NotifyRequest)

// 支付回调请求分析
func (client *PayClient) PayNotify(httpReq *http.Request) (error, *NotifyRequest) {
	reqBody, err := ioutil.ReadAll(httpReq.Body)
	if err != nil {
		return errors.New(err.Error()), nil
	}

	sign := util.NewSign(client.ApiKey)
	isMatch := sign.VerifySign(httpReq.RequestURI, string(reqBody))
	if isMatch {
		var notify NotifyRequest
		json.Unmarshal(reqBody, &notify)
		return nil, &notify
	} else {
		return errors.New("签名不一致"), nil
	}
}
