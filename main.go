package main

import (
	"fmt"
	"math/rand"
	"time"
)

var quit = make(chan bool)

func fibonacci(number float64, ch chan float64) {
	x, y := 1.0, 1.0

	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
	ch <- x
}

func fib2(c chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}

func main() {
	start := time.Now()
	ch := make(chan float64)
	r := rand.Intn(30)
	for i := 0; i < r; i++ {
		go fibonacci(float64(i), ch)
	}

	for i := 0; i < r; i++ {
		fmt.Printf("Fib(%v): %v\n", i, <-ch)
	}

	command := ""
	data := make(chan int)

	go fib2(data)

	for {
		num := <-data
		fmt.Println(num)
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}

	time.Sleep(1 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Done! it took %v seconds!\n", elapsed.Seconds())
}
