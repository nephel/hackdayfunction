package hackdayfunction

import (
	"fmt"
	"time"
)

func Handle(payload map[string]interface{}) {
	fmt.Printf("I received an event with payload: %+v\n", payload)
	time.Sleep(5)
}
