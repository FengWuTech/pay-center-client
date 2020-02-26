package pay_center_client

import (
	"fmt"
	"testing"
)

func TestPayClient_GetDailyRechargeStatistics(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	res := client.GetDailyRechargeStatistics(1, 2020, 2)
	fmt.Printf("%v", *res)
}

func TestPayClient_GetMonthRechargeStatistics(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	res := client.GetMonthRechargeStatistics(1, 2020)
	fmt.Printf("%v", *res)
}
