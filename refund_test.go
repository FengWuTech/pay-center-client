package pay_center_client

import (
	"fmt"
	"testing"
)

func TestPayClient_RefundToAccount(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	var request = RefundToAccountRequest{
		BillID: 1,
	}
	response := client.RefundToAccount(request)
	fmt.Printf("%v", response)
}
