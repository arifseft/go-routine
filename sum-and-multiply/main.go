package main

import (
	"fmt"
	"time"
)

func sum(first int, second int, sumChan chan int) {
	sum := first + second
	time.Sleep(1000 * time.Millisecond)
	sumChan <- sum
}

func multiply(first int, second int, multiplyChan chan int) {
	time.Sleep(2000 * time.Millisecond)
	multiplyChan <- first * second
}

func main() {
	first := 23
	second := 2
	fmt.Printf("First= %d\n", first)
	fmt.Printf("Second= %d\n", second)

	sumChan := make(chan int)
	multiplyChan := make(chan int)

	go sum(first, second, sumChan)
	go multiply(first, second, multiplyChan)

	fmt.Println("Calculating...")

	s := <-sumChan
	fmt.Printf("Sum of %d and %d is %d\n", first, second, s)

	m := <-multiplyChan
	fmt.Printf("Multiply of %d and %d is %d\n", first, second, m)
}
