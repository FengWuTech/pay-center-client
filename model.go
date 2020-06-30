package pay_center_client

import (
	"time"
)

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

type BillGoPayRequest struct {
	CompanyID    int           `json:"company_id"`
	GroupID      int           `json:"group_id"`
	WxSubAppID   string        `json:"wx_sub_app_id"`
	WxSubOpenID  string        `json:"wx_sub_open_id"`
	Title        string        `json:"title"`
	Amount       int           `json:"amount"`
	NotifyURL    string        `json:"notify_url"`
	NotifyAttach string        `json:"notify_attach"`
	UserID       int           `json:"user_id"`
	UserIP       string        `json:"user_ip"`
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

type BillCashPayRequest struct {
	UserID     int              `json:"user_id"`
	CompanyID  int              `json:"company_id" valid:"Required"`
	GroupID    int              `json:"group_id" valid:"Required"`
	Remark     string           `json:"remark"`
	PayChannel []PayChannelItem `json:"pay_channel_list"`
	BillList   []PayBillItem    `json:"bill_list" valid:"Required"`
	PayTime    time.Time        `json:"pay_time"`
}

type BillCashPayResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type BillPayInfoResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		DeductList []struct {
			BillID          int       `json:"bill_id"`
			AccountID       int       `json:"account_id"`
			AccountName     string    `json:"account_name"`
			AmountBefore    int       `json:"amount_before"`
			AmountAfter     int       `json:"amount_after"`
			AmountChange    int       `json:"amount_change"`
			AccountFlowType int       `json:"account_flow_type"`
			CreateTime      time.Time `json:"create_time"`
			UpdateTime      time.Time `json:"update_time"`
		} `json:"deduct_list"`
		PayList []struct {
			PayShareID     int       `json:"pay_share_id"`
			PayShareAmount int       `json:"pay_share_amount"`
			PayFlowID      int       `json:"pay_flow_id"`
			PayChannel     int       `json:"pay_channel"`
			PayTime        time.Time `json:"pay_time"`
		} `json:"pay_list"`
	} `json:"data"`
}

type GetPayFlowResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ID           *int       `json:"id"`             // ID
		AppID        *string    `json:"app_id"`         // 应用ID
		BatchNo      *string    `json:"batch_no"`       // 批量号，用来标记支付批次
		CompanyId    *int       `json:"company_id"`     // 公司ID
		BizType      *int       `json:"biz_type"`       // 业务类型 1.预交费 2.账单支付 3.押金支付
		LocalTradeNo *string    `json:"local_trade_no"` // 本地流水单号
		ThirdTradeNo *string    `json:"third_trade_no"` // 三方交易单号
		PayChannel   *int       `json:"pay_channel"`    // 支付渠道
		PayAmount    *int       `json:"pay_amount"`     // 支付金额，单位分
		PayTime      *time.Time `json:"pay_time"`       // 支付时间
		PayStatus    *int       `json:"pay_status"`     // 支付状态 1.待支付 2已支付
		UserId       *int       `json:"user_id"`        // 支付用户ID
		DeductAmount *int       `json:"deduct_amount"`  // 抵扣金额
		AccountID    *int       `json:"account_id"`     // 资金账户ID
		Remark       *string    `json:"remark"`         // 备注
		CreateTime   *time.Time `json:"create_time"`    // 创建时间
		UpdateTime   *time.Time `json:"update_time"`    // 更新时间
	}
}
