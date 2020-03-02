package pay_center_client

import (
	"fmt"
	"testing"
)

func TestAddAccount(t *testing.T) {
	response := NewClient().CreateAccount(CreateAccountRequest{
		CompanyID: 1,
		UserID:    2,
		Type:      0,
		Name:      "通用账户",
	})
	fmt.Printf("%v", response)
}

func TestEditAccount(t *testing.T) {
	response := NewClient().EditAccount(EditAccountRequest{
		ID:     4,
		Name:   "通用资金账户",
		Status: 0,
	})
	fmt.Printf("%v", response)
}

func TestGetAccount(t *testing.T) {
	response := NewClient().GetAccount(4)
	fmt.Printf("%v", response)
}

func TestPayClient_GetUserAccount(t *testing.T) {
	response := NewClient().GetUserAccount(17)
	fmt.Printf("%v", response)
}
