package pipline

func LaunchPipeline(amount int) int {
	firstCh := generator(amount)
	secondCh := power(firstCh)
	thirdCh := sum(secondCh)
	result := <-thirdCh
	return result
}

func generator(max int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for i := 1; i <= max; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func power(in <-chan int) <-chan int {
	out := make(chan int, len(in))
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, len(in))
	go func() {
		var sum int
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
