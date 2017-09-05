package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/pquerna/ffjson/ffjson"
)

func main() {
	const jsonStream = `
  		{"Name": "Ed", "Text": "Knock knock."}
  	`
	type Message struct {
		Name, Text string
	}

	dec := ffjson.NewDecoder()

	//dec := json.NewDecoder(strings.NewReader(jsonStream))
	var m Message

	//if err := dec.Decode(&m); err == io.EOF {
	if err := dec.DecodeReader(strings.NewReader(jsonStream), &m); err == io.EOF {
	} else if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s: %s\n", m.Name, m.Text)
}
