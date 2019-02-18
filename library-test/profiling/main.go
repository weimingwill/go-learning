package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

// bigBytes allocates 10 sets of 100 megabytes
func bigBytes() *[]byte {
	s := make([]byte, 100000000)
	return &s
}

func main() {
	var wg sync.WaitGroup

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for i := 0; i < 10; i++ {
		s := bigBytes()
		if s == nil {
			log.Println("oh noes")
		}
	}

	wg.Add(1)
	wg.Wait() // this is for the benefit of the pprof server analysis
}

// func main() {
// 	// pprof.StartCPUProfile(os.Stdout)
// 	// defer pprof.StopCPUProfile()

// 	for i := 0; i < 10; i++ {
// 		s := bigBytes()
// 		if s == nil {
// 			log.Println("oh noes")
// 		}
// 	}

// 	pprof.WriteHeapProfile(os.Stdout)
// }
