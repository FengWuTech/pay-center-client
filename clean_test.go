package pay_center_client

import (
	"fmt"
	"testing"
)

func TestPayClient_CleanCompanyData(t *testing.T) {
	response := NewClient().CleanCompanyData(20)
	fmt.Printf("%v", response)
}
