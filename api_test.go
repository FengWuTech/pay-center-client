package pay_center_client

import (
	"fmt"
	"testing"
)

const APPID = "test"
const APIKEY = "testtesttesttesttesttesttesttest"

func TestRechargeGoPay(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	var request = RechargeGoPayRequest{
		WxSubAppID:   "wxfa424fd01a813cce",
		WxSubMchID:   "1575841741",
		WxSubOpenID:  "oVsiGv3nClXr_fGtoQ1I7lgYZgU0",
		Title:        "测试",
		Amount:       1,
		NotifyURL:    "https://pms-api.gmtech.top/pay-center/minip-user/recharge/gopay/notify",
		NotifyAttach: "{}",
		UserID:       17,
		UserIP:       "127.0.0.2",
		CompanyID:    1,
		DeductAmount: 100,
	}
	response := client.RechargeGoPay(request)
	fmt.Printf("%v", response)
}

func TestBillGoPay(t *testing.T) {
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
	response := client.BillGoPay(request)
	fmt.Printf("%v", response)
}

func TestCashRecharge(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	response := client.CashRecharge(CashRechargeRequest{
		Amount:       1,
		UserID:       17,
		CompanyID:    1,
		DeductAmount: 10,
	})
	fmt.Printf("%v", response)
}

func TestAddAccount(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	response := client.CreateAccount(CreateAccountRequest{
		AppID:     "test",
		CompanyID: 1,
		UserID:    1,
		Type:      0,
		Name:      "通用账户",
	})
	fmt.Printf("%v", response)
}

func TestEditAccount(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	response := client.EditAccount(EditAccountRequest{
		ID:     4,
		Name:   "通用资金账户",
		Status: 0,
	})
	fmt.Printf("%v", response)
}

func TestGetAccount(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	response := client.GetAccount(4)
	fmt.Printf("%v", response)
}
