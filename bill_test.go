package pay_center_client

import (
	"fmt"
	"testing"
)

func TestPayClient_BillWeixinGoPay(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	var request = BillGoPayRequest{
		WxSubAppID:   "wxfa424fd01a813cce",
		WxSubMchID:   "1575841741",
		WxSubOpenID:  "oVsiGv3nClXr_fGtoQ1I7lgYZgU0",
		Title:        "测试",
		Amount:       1,
		NotifyURL:    "https://pms-api.gmtech.top/pay-center/minip-user/billpay/gopay/notify",
		NotifyAttach: "{}",
		UserID:       17,
		UserIP:       "127.0.0.2",
		CompanyID:    1,
		DeductAmount: 100,
		BillList: []BillGoPayBillListItem{
			{
				ID:        218,
				PayAmount: 10,
			},
			{
				ID:        219,
				PayAmount: 10,
			},
		},
	}
	response := client.BillWeixinGoPay(request)
	fmt.Printf("%v", response)
}

func TestPayClient_BillCashPay(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	var request = BillCashPayRequest{
		Amount:    1,
		UserID:    17,
		CompanyID: 1,
		BillList: []BillGoPayBillListItem{
			{
				ID:        218,
				PayAmount: 10,
			},
			{
				ID:        219,
				PayAmount: 10,
			},
		},
	}
	response := client.BillCashPay(request)
	fmt.Printf("%v", response)
}
