package pay_center_client

import (
	"fmt"
	"testing"
)

func TestPayClient_GetAccountFlow(t *testing.T) {
	response := NewClient().GetAccountFlow(GetAccountFlowRequest{
		CompanyID:   1,
		FlowType:    1,
		AccountType: 0,
		UserID:      17,
		Page:        1,
		PageSize:    10,
	})
	fmt.Printf("%v", response)
}
