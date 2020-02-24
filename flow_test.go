package pay_center_client

import (
	"fmt"
	"testing"
)

func TestPayClient_GetAccountFlow(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	response := client.GetAccountFlow(17, 1, 10)
	fmt.Printf("%v", response)
}
