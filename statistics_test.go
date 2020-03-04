package pay_center_client

import (
	"fmt"
	"testing"
)

func TestPayClient_GetDailyRechargeStatistics(t *testing.T) {
	res := NewClient().GetDailyRechargeStatistics(1, 2020, 2)
	fmt.Printf("%v", *res)
}

func TestPayClient_GetMonthRechargeStatistics(t *testing.T) {
	res := NewClient().GetMonthRechargeStatistics(1, 2020)
	fmt.Printf("%v", *res)
}

func TestPayClient_GetMonthIncomeStatistics(t *testing.T) {
	res := NewClient().GetMonthIncomeStatistics(1, 2020)
	fmt.Printf("%v", *res)
}

func TestPayClient_GetDailyIncomeStatistics(t *testing.T) {
	res := NewClient().GetDailyIncomeStatistics(1, 2020, 3)
	fmt.Printf("%v", *res)
}
