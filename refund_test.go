package pay_center_client

import (
	"fmt"
	"testing"
)

func TestPayClient_RefundToAccount(t *testing.T) {
	var request = RefundToAccountRequest{
		BillID: 1,
	}
	response := NewClient().RefundToAccount(request)
	fmt.Printf("%v", response)
}
