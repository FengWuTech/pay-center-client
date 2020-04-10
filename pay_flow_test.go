package pay_center_client

import (
	"fmt"
	"testing"
)

func TestPayClient_GetPayFlow(t *testing.T) {
	response := NewClient().GetPayFlow(5)
	fmt.Printf("%v", response)
}
