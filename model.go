package pay_center_client

type RechargeGoPayRequest struct {
	PayAppID     string `json:"pay_app_id"`
	WxSubAppID   string `json:"wx_sub_app_id"`
	WxSubMchID   string `json:"wx_sub_mch_id"`
	WxSubOpenID  string `json:"wx_sub_open_id"`
	Title        string `json:"title"`
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

type BillGoPayBillListItem struct {
	ID        int `json:"id"`
	PayAmount int `json:"pay_amount"`
}

type BillGoPayRequest struct {
	PayAppID     string                  `json:"pay_app_id"`
	WxSubAppID   string                  `json:"wx_sub_app_id"`
	WxSubMchID   string                  `json:"wx_sub_mch_id"`
	WxSubOpenID  string                  `json:"wx_sub_open_id"`
	Title        string                  `json:"title"`
	Amount       int                     `json:"amount"`
	NotifyURL    string                  `json:"notify_url"`
	NotifyAttach string                  `json:"notify_attach"`
	UserID       int                     `json:"user_id"`
	UserIP       string                  `json:"user_ip"`
	CompanyID    int                     `json:"company_id"`
	DeductAmount int                     `json:"deduct_amount"`
	BillList     []BillGoPayBillListItem `json:"bill_list"`
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
