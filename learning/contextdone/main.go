package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	ctx := context.Background()

	// go test1(ctx)
	// go test2(ctx)
	// cancel()
	// time.Sleep(500 * time.Millisecond)
	done := make(chan struct{})
	done1 := make(chan struct{})
	done3 := make(chan struct{})

	go test1(ctx, done, done1)
	go test3(ctx, done, done3)

	done1 <- struct{}{}
	done3 <- struct{}{}
	test(done)
	time.Sleep(1000 * time.Millisecond)
}

func test(done chan struct{}) {
	c := 0

LOOP:
	for {
		select {
		case <-done:
			c++
			fmt.Println(c)
		case <-time.After(1 * time.Second):
			break LOOP
		}
	}
	fmt.Println("out")
}

func test1(ctx context.Context, done chan struct{}, done1 chan struct{}) {
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done 1")
			break LOOP

		case <-done1:
			fmt.Println("done 1")
			break LOOP
		}
	}
	done <- struct{}{}
}

func test3(ctx context.Context, done chan struct{}, done3 chan struct{}) {
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done 1")
			break LOOP

		case <-done3:
			fmt.Println("done 3")
			break LOOP
		}
	}
	done <- struct{}{}
}

func test2(ctx context.Context, done chan struct{}) {
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done 2")
			break LOOP
		}
	}
}
