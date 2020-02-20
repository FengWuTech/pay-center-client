package pay_center_client

import (
	"fmt"
	"testing"
)

func TestRechargeCashPay(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	response := client.RechargeCashPay(CashRechargeRequest{
		AccountID:    1,
		Amount:       1,
		UserID:       17,
		CompanyID:    1,
		DeductAmount: 10,
	})
	fmt.Printf("%v", response)
}

func TestRechargeGoPay(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	var request = RechargeGoPayRequest{
		AccountID:    1,
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
	response := client.RechargeWeixinGoPay(request)
	fmt.Printf("%v", response)
}
