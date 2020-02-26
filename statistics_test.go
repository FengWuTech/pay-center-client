package pay_center_client

import (
	"fmt"
	"testing"
)

func TestPayClient_GetPrepayStatistics(t *testing.T) {
	client := NewPayClient(APPID, APIKEY)
	res := client.GetPrepayStatistics(1, 2020, 2)
	fmt.Printf("%v", *res)
}
