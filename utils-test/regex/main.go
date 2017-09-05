package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	str := "-url=http://172.26.158.201:17050/api/v2/api-docs -target=/v1/products/{productId} -method=get -payload=payload.json -rate=1 -requests=5 -server=true"

	reg := regexp.MustCompile(`-requests=\d`)

	newReq := fmt.Sprintf("-requests=%v", 1000)
	str = reg.ReplaceAllString(str, newReq)
	fmt.Println(str)

	now := time.Now()
	fmt.Println(now.Format("20060102150405"))
}
