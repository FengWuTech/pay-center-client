package pay_center_client

import (
	"fmt"
	"testing"
)

func TestPayClient_BillDeductAuto(t *testing.T) {
	response := NewClient().BillDeductAuto(BillDeductAutoRequest{
		AccountID: 1,
		BillID:    1,
		Amount:    10,
	})
	fmt.Printf("%v", response)
}

func TestPayClient_BillWeixinGoPay(t *testing.T) {
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
		Amount:    1,
		UserID:    17,
		CompanyID: 1,
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
	response := NewClient().BillCashPay(request)
	fmt.Printf("%v", response)
}
