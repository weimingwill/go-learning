package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"time"
)

type Test struct {
	Running bool
	Name    string `json:"name"`
	ch      chan bool

	num int
}

func main() {
	err := fmt.Errorf("this is error")

	fmt.Printf("%s\n", err)
	// a := [][]string{[]string{"z"}, []string{"a", "asd"}}
	// b := "asdf"
	// fmt.Println(a, string(b[1]))

	//ctx, cancelFn := context.WithCancel(context.Background())
	//testCh := make(chan struct{})
	//go func() {
	//	looptest(ctx, testCh)
	//	fmt.Println("break out loop")
	//}()
	//cancelFn()
	//
	//time.Sleep(time.Second)
	//fmt.Println("exit")

	//cmd := docker.RunCmd("ares", "ares", "-url=http://1111", 3000)
	//fmt.Println(cmd)
}

func looptest(ctx context.Context, testCh chan struct{}) {
	for {
		select {
		case <-testCh:
			fmt.Println("test")
		case <-ctx.Done():
			fmt.Println("return")
			return
		}
	}
}

func testJSON() {
	test := &Test{
		Running: true,
	}
	fmt.Println("before", test)
	changeTest(test)
	fmt.Println("after", test)
}

func changeTest(test *Test) {
	test = &Test{
		Running: false,
	}
}

func testContext() {
	parentCtx, _ := context.WithCancel(context.Background())

	childCtx, childCancel := context.WithCancel(parentCtx)

	childCancel()

loop:
	for {
		select {
		case <-parentCtx.Done():
			fmt.Println("parent done")
			break loop
		case <-childCtx.Done():
			fmt.Println("child done")
		}
	}
}

func testJson() {
	test := Test{
		Name: "test",
		ch:   make(chan bool),
	}

	b, err := json.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}

	var t Test
	json.Unmarshal(b, &t)
	fmt.Println(t)
}

func happy(test chan string) {
	//fmt.Printf("happt")
	test <- "test"
}

func random(val int64) int64 {
	src := rand.NewSource(time.Now().UnixNano() * val)
	r := rand.New(src).Intn(10)
	return int64(r)
}

// GetLocalIP returns IP of local machine
func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println("error closing connection", err)
		}
	}()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String(), nil
}
