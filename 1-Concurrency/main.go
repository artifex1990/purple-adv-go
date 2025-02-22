package main

import (
	"fmt"
	"go/adv-dz1/generator"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func main() {
	maxBuffer := 10
	maxRandomNumber := 101
	firstCh := make(chan int, maxBuffer)
	secondCh := make(chan int, maxBuffer)

	rand.Intn(maxRandomNumber)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, num := range generator.Generate(maxBuffer) {
			firstCh <- num
		}
		close(firstCh)
	}()

	go func() {
		defer wg.Done()
		for num := range firstCh {
			secondCh <- num * num
		}
		close(secondCh)
	}()

	wg.Wait()

	for ch := range secondCh {
		fmt.Print(ch, " ")
	}
}
