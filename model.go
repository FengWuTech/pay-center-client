package pay_center_client

import "time"

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

type RechargeGoPayRequest struct {
	WxSubAppID   string `json:"wx_sub_app_id"`
	WxSubMchID   string `json:"wx_sub_mch_id"`
	WxSubOpenID  string `json:"wx_sub_open_id"`
	Title        string `json:"title"`
	AccountID    int    `json:"account_id"`
	Amount       int    `json:"amount"`
	NotifyURL    string `json:"notify_url"`
	NotifyAttach string `json:"notify_attach"`
	UserID       int    `json:"user_id"`
	UserIP       string `json:"user_ip"`
	CompanyID    int    `json:"company_id"`
	DeductAmount int    `json:"deduct_amount"`
}

type RechargeGoPayResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AppID     string `json:"appId"`
		TimeStamp string `json:"timeStamp"`
		NonceStr  string `json:"nonceStr"`
		Package   string `json:"package"`
		SignType  string `json:"signType"`
		Sign      string `json:"sign"`
	} `json:"data"`
}

type PayBillItem struct {
	ID        int `json:"id"`
	PayAmount int `json:"pay_amount"`
}

type BillGoPayRequest struct {
	WxSubAppID   string        `json:"wx_sub_app_id"`
	WxSubMchID   string        `json:"wx_sub_mch_id"`
	WxSubOpenID  string        `json:"wx_sub_open_id"`
	Title        string        `json:"title"`
	Amount       int           `json:"amount"`
	NotifyURL    string        `json:"notify_url"`
	NotifyAttach string        `json:"notify_attach"`
	UserID       int           `json:"user_id"`
	UserIP       string        `json:"user_ip"`
	CompanyID    int           `json:"company_id"`
	DeductAmount int           `json:"deduct_amount"`
	BillList     []PayBillItem `json:"bill_list"`
}

type BillGoPayResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AppID     string `json:"appId"`
		TimeStamp string `json:"timeStamp"`
		NonceStr  string `json:"nonceStr"`
		Package   string `json:"package"`
		SignType  string `json:"signType"`
		Sign      string `json:"sign"`
	} `json:"data"`
}

type NotifyRequest struct {
	Attach string `json:"attach"`
}

type NotifyResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RechargeCashPayRequest struct {
	AccountID    int `json:"account_id"`
	Amount       int `json:"amount"`
	UserID       int `json:"user_id"`
	CompanyID    int `json:"company_id"`
	DeductAmount int `json:"deduct_amount"`
}

type RechargeCashPayResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CreateAccountRequest struct {
	CompanyID int    `json:"company_id" valid:"Required"`
	UserID    int    `json:"user_id" valid:"Required"`
	Type      int    `json:"type" valid:"Required"`
	Name      string `json:"name" valid:"Required"`
}

type CreateAccountResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AccountID int `json:"account_id"`
	} `json:"data"`
}

type EditAccountRequest struct {
	ID     int    `json:"id" valid:"Required"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type EditAccountResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type GetAccountResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ID              int       `json:"id"`
		Appid           string    `json:"appid"`
		CompanyId       int       `json:"company_id"`
		UserId          int       `json:"user_id"`
		Type            int       `json:"type"`
		Name            string    `json:"name"`
		Status          int       `json:"status"`
		AmountAll       int       `json:"amount_all"`
		AmountAvailable int       `json:"amount_available"`
		AmountBlocked   int       `json:"amount_blocked"`
		CreateTime      time.Time `json:"create_time"`
		UpdateTime      time.Time `json:"update_time"`
	} `json:"data"`
}

type RefundToAccountRequest struct {
	AccountID int `json:"account_id"`
	Amount    int `json:"amount"`
}

type RefundToAccountResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type BillCashPayRequest struct {
	Amount    int           `json:"amount" valid:"Required"`
	UserID    int           `json:"user_id" valid:"Required"`
	CompanyID int           `json:"company_id" valid:"Required"`
	BillList  []PayBillItem `json:"bill_list" valid:"Required"`
}

type BillCashPayResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
