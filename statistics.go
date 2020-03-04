package pay_center_client

import (
	"encoding/json"
	"github.com/FengWuTech/pay-center-client/util/httputil"
	"github.com/FengWuTech/pay-center-client/util/signutil"
)

func (client *PayClient) GetDailyRechargeStatistics(companyID int, year int, month int) *Response {
	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.AddQuery("company_id", companyID)
	sign.AddQuery("year", year)
	sign.AddQuery("month", month)
	url := sign.GenSignURL(client.Host + URL_STATISTICS_RECHARGE_DAILY)

	_, res := httputil.Get(url)

	var response Response
	json.Unmarshal(res, &response)
	return &response
}

func (client *PayClient) GetMonthRechargeStatistics(companyID int, year int) *Response {
	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.AddQuery("company_id", companyID)
	sign.AddQuery("year", year)
	url := sign.GenSignURL(client.Host + URL_STATISTICS_RECHARGE_MONTH)

	_, res := httputil.Get(url)

	var response Response
	json.Unmarshal(res, &response)
	return &response
}

func (client *PayClient) GetMonthIncomeStatistics(companyID int, year int) *Response {
	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.AddQuery("company_id", companyID)
	sign.AddQuery("year", year)
	url := sign.GenSignURL(client.Host + URL_STATISTICS_INCOME_MONTH)

	_, res := httputil.Get(url)

	var response Response
	json.Unmarshal(res, &response)
	return &response
}

func (client *PayClient) GetDailyIncomeStatistics(companyID int, year int, month int) *Response {
	sign := signutil.NewSign(client.ApiKey)
	sign.AddQuery("appid", client.AppID)
	sign.AddQuery("company_id", companyID)
	sign.AddQuery("year", year)
	sign.AddQuery("month", month)
	url := sign.GenSignURL(client.Host + URL_STATISTICS_INCOME_DAILY)

	_, res := httputil.Get(url)

	var response Response
	json.Unmarshal(res, &response)
	return &response
}
