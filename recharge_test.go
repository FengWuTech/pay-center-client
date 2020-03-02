package pay_center_client

import (
	"fmt"
	"testing"
)

func TestPayClient_RechargeCashPay(t *testing.T) {
	response := NewClient().RechargeCashPay(RechargeCashPayRequest{
		AccountID:    1,
		Amount:       1,
		UserID:       17,
		CompanyID:    1,
		DeductAmount: 10,
	})
	fmt.Printf("%v", response)
}

func TestPayClient_RechargeWeixinGoPay(t *testing.T) {
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
	response := NewClient().RechargeWeixinGoPay(request)
	fmt.Printf("%v", response)
}

func TestPayClient_RechargeHistory(t *testing.T) {
	response := NewClient().RechargeHistory(17, 1, 10)
	fmt.Printf("%v", response)
}
