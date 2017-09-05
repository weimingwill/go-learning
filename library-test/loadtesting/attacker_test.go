package loadtesting

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestAttacker(t *testing.T) {
	rate := uint64(1) // per second
	duration := 10 * time.Second
	target := Target{
		Method: "GET",
		URL:    "http://blog.zhuangweiming.me",
		Header: http.Header{
			"Accept":     []string{"application/json"},
			"User-agent": []string{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"},
		},
	}

	targeter := NewStaticTargeter(target, target)

	attacker := NewAttacker(
		Connections(1),
	)

	count := 0
	var metrics Metrics
	for res := range attacker.Attack2(targeter, rate, duration) {
		count += 1
		metrics.Add(res)
		fmt.Println(res)
	}
	metrics.Close()
	fmt.Println("Success Rate: ", metrics.Success)
	for key, val := range metrics.Errors {
		fmt.Println("Error: ", key)
		fmt.Println("Error Count: ", val)
	}
}
