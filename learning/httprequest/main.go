package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/pquerna/ffjson/ffjson"
)

type Test struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type Value struct {
	value string `json:"value"`
}

func main() {
	fmt.Println(rand.Intn(100))
}

func sendToEtcd( ){
	m := make(map[string]string)
	m["value"] = "leaderAddr"
	b, err := ffjson.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	http.Post("http://127.0.0.1:2379/v2/keys/leaderAddr", "application-type/json", bytes.NewReader(b))
}

func sendRandomKeyVal() {
	for i := 0; i < 1000000; i++ {
		test := Test{Key: "set-" + RandStringRunes(5), Name: RandStringRunes(6)}
		testB, err := ffjson.Marshal(test)
		if err != nil {
			fmt.Println(err)
		}
		_, err = http.Post(fmt.Sprintf("http://172.26.158.167:11000/test"), "application-type/json", bytes.NewReader(testB))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
