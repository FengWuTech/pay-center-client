package logutil

import (
	"fmt"
)

func ErrorF(format string, args ...interface{}) {
	fmt.Printf("pay-center-client ERROR "+format+"\n", args...)
}

func WarnF(format string, args ...interface{}) {
	fmt.Printf("pay-center-client WARN "+format+"\n", args...)
}

func InfoF(format string, args ...interface{}) {
	fmt.Printf("pay-center-client INFO "+format+"\n", args...)
}

func DebugF(format string, args ...interface{}) {
	fmt.Printf("pay-center-client DEBUG "+format+"\n", args...)
}
