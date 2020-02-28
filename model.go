package pay_center_client

import "time"

type PayClient struct {
	Host   string
	AppID  string
	ApiKey string
}

// 实例化请求端
func NewPayClient(appID string, apiKey string, host string) *PayClient {
	var client PayClient
	client.AppID = appID
	client.ApiKey = apiKey
	client.Host = host
	return &client
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PayBillItem struct {
	ID        int `json:"id"`
	PayAmount int `json:"pay_amount"`
}

type PayChannelItem struct {
	PayChannelID int `json:"pay_channel_id"`
	Amount       int `json:"amount"`
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
	AccountID    int    `json:"account_id"`
	Amount       int    `json:"amount"`
	UserID       int    `json:"user_id"`
	CompanyID    int    `json:"company_id"`
	DeductAmount int    `json:"deduct_amount"`
	PayChannel   int    `json:"pay_channel"`
	Remark       string `json:"remark"`
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
	BillID int `json:"bill_id"`
}

type RefundToAccountResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RefundToUserRequest struct {
	AccountID int `json:"account_id"`
	Amount    int `json:"amount"`
}

type RefundToUserResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type BillCashPayRequest struct {
	UserID     int              `json:"user_id"`
	CompanyID  int              `json:"company_id" valid:"Required"`
	Remark     string           `json:"remark"`
	PayChannel []PayChannelItem `json:"pay_channel_list"`
	BillList   []PayBillItem    `json:"bill_list" valid:"Required"`
}

type BillCashPayResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type BillDeductAutoRequest struct {
	AccountID int `json:"account_id"`
	BillID    int `json:"bill_id"`
	Amount    int `json:"amount"`
}

type BillDeductAutoResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type GetAccountFlowRequest struct {
	CompanyID   int `json:"company_id"`
	FlowType    int `json:"flow_type"`
	AccountType int `json:"account_type"`
	UserID      int `json:"user_id"`
	Page        int `json:"page"`
	PageSize    int `json:"page_size"`
}

type GetAccountFlowResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Total int `json:"total"`
		List  []struct {
			ID            *int       `json:"id"`            // 充值列表ID
			AccountId     *int       `json:"account_id"`    // 资金账户ID
			FlowType      *int       `json:"flow_type"`     // 变动类型：1.用户充值 2.物业公司退费 3.自动划扣
			AmountBefore  *int       `json:"amount_before"` // 变动前资金账户金额
			AmountChange  *int       `json:"amount_change"` // 变动金额
			AmountAfter   *int       `json:"amount_after"`  // 变动后金额
			PayFlowId     *int       `json:"pay_flow_id"`   // 支付流水ID
			BillID        *int       `json:"bill_id"`       // 扣款账单ID
			CreateTime    *time.Time `json:"create_time"`
			UpdateTime    *time.Time `json:"update_time"`
			PayRealAmount *int       `json:"pay_real_amount"`
			AccountName   *string    `json:"account_name"`
			AccountType   *int       `json:"account_type"`
			UserID        *int       `json:"user_id"`
			Remark        *string    `json:"remark"`
		}
		StatisticsAll struct {
			ShouldFee int `json:"should_fee"`
			RealFee   int `json:"real_fee"`
			Expend    int `json:"expend"`
		} `json:"statistics_all"`
		StatisticsPage struct {
			ShouldFee int `json:"should_fee"`
			RealFee   int `json:"real_fee"`
			Expend    int `json:"expend"`
		} `json:"statistics_page"`
	} `json:"data"`
}
