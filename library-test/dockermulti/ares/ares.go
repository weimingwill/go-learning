package main

import (
	"fmt"

	"go-learning/dockermulti/ares/util"
	"go-learning/dockermulti/common"
)

func main() {
	util.HelloPrint()
	common.WorldPrint()
	fmt.Println("input", util.Flag().Input)
}
