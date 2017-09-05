package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {

	for _, user := range users {
		for _, char := range user {
			coin := 0
			switch string(char) {
			case "a", "A":
				coin = 1
			case "e", "E":
				coin = 1
			case "i", "I":
				coin = 2
			case "o", "O":
				coin = 3
			case "u", "U":
				coin = 4
			}
			distribution[user] += coin
		}
		if distribution[user] > 10 {
			distribution[user] = 10
		}
		coins -= distribution[user]

	}

	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}
