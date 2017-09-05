package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

type Test struct {
	A string
	B string
	I int
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	scmd := client.LRange("a", 0, 9)
	locations, err := scmd.Result()
	if err != nil {
		log.Println(err)
	}
	log.Println(locations)

}

func ping() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	ts := []Test{}

	for x := 0; x <= 10; x++ {
		t := Test{"a", "b", x}
		ts = append(ts, t)
		// b, _ := json.Marshal(t)

		// i := client.RPush("trial1", b)
		// if i.Err() != nil {
		// 	fmt.Println(i.Err())
		// }

		// res, err := i.Result()
		// if err != nil {
		// 	fmt.Println(res, err)
		// }

		// if res > 10 {
		// 	i := client.LPop("trial")
		// 	if i.Err() != nil {
		// 		fmt.Println(i.Err())
		// 	}
		// }
		// fmt.Println(res, err)

		// fmt.Printf("%+v\n", i)
	}

	scmd := client.HGet("tmap", "user1")
	s, err := scmd.Result()
	fmt.Println(s)

	b, _ := json.Marshal(ts)

	if err != nil && err == redis.Nil {
		bcmd := client.HSet("tmap", "user1", b)
		success, err := bcmd.Result()
		fmt.Println(success)
		fmt.Println(err)
	}

}
