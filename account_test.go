package pay_center_client

import (
	"fmt"
	"testing"
)

func TestAddAccount(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	response := client.CreateAccount(CreateAccountRequest{
		CompanyID: 1,
		UserID:    2,
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
