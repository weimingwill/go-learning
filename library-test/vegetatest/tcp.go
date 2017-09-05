package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
	"net/http"
)

func main() {
	rate := uint64(1) // per second
	duration := 1 * time.Second
	target := vegeta.Target{
		Method: "GET",
		URL:    "http://172.26.158.201:17050/api/v1/cart/basket/items",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
	}
	target.Header.Add("client-id", "tenant001")
	target.Header.Add("access-token", "eyJhbGciOiJSUzI1NiJ9.eyJjbGllbnRJZCI6InRlbmFudDAwMSIsImV4cGlyYXRpb25EYXRlIjoiMjAxNy0wOS0xOVQxNTowOTozOS43NzhaW0V0Yy9VVENdIiwiZW5kcG9pbnRzIjpbIipAKiJdfQ.HRDQO5of4DK7JPipgISDKygvUvxg_acZ-UQCDooqNoFFgcCJJbgfcwvZjMrPI-_bQzhxbvirLet4Y_eytdhENjZET2HBkmKHUkxnW1tt1G9jZG8biu397kZ2qnPSBf5y3buGm1ulB6jIDRmoevnL4xtWMjdqylkWdPQ92-OC4pg")
	target.Header.Add("storeId", "eccefec0-87d2-11e7-a3e1-4d3a169b023e")
	target.Header.Add("siteId", "31626270-87d3-11e7-a3e1-4d3a169b023e")
	target.Header.Add("X-Requested-With", "XMLHttpRequest")

	targeter := vegeta.NewStaticTargeter(target)

	attacker := vegeta.NewAttacker()

	count := 0
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration) {
		count += 1
		metrics.Add(res)
		fmt.Println(res)
	}
	metrics.Close()
	fmt.Println(count)
	fmt.Println("Success: ", metrics.Success)
	fmt.Println("Errors: ", len(metrics.Errors))
	for _, err := range metrics.Errors {
		fmt.Println("Error: " + err)
	}
}
