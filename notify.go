package pay_center_client

import (
	"encoding/json"
	"errors"
	"github.com/FengWuTech/pay-center-client/util/signutil"
	"io/ioutil"
	"net/http"
)

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
