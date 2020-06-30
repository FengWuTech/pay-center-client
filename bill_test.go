package pay_center_client

import (
	"fmt"
	"testing"
)

func TestPayClient_BillWeixinGoPay(t *testing.T) {
	var request = BillGoPayRequest{
		CompanyID:    1,
		GroupID:      20,
		WxSubAppID:   "wxfa424fd01a813cce",
		WxSubOpenID:  "oVsiGv3nClXr_fGtoQ1I7lgYZgU0",
		Title:        "测试",
		Amount:       1,
		NotifyURL:    "https://pms-api.gmtech.top/pay-center/minip-user/billpay/gopay/notify",
		NotifyAttach: "{}",
		UserID:       17,
		UserIP:       "127.0.0.2",
		DeductAmount: 100,
		BillList: []PayBillItem{
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
	response := NewClient().BillWeixinGoPay(request)
	fmt.Printf("%v", response)
}

func TestPayClient_BillCashPay(t *testing.T) {
	var request = BillCashPayRequest{
		UserID:    1,
		CompanyID: 1,
		GroupID:   20,
		Remark:    "sdk测试",
		PayChannel: []PayChannelItem{
			{
				PayChannelID: 1,
				Amount:       100000,
			},
			{
				PayChannelID: 2,
				Amount:       1,
			},
			{
				PayChannelID: 3,
				Amount:       1,
			},
		},
		BillList: []PayBillItem{
			{
				ID:        1,
				PayAmount: 1,
			},
			{
				ID:        2,
				PayAmount: 1,
			},
		},
	}
	response := NewClient().BillCashPay(request)
	fmt.Printf("%v", response)
}

func TestPayClient_BillPayInfo(t *testing.T) {
	res := NewClient().BillPayInfo(2904)
	fmt.Printf("%v", res)
}
