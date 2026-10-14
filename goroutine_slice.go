package main


import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)

	const grs = 4
	var wg sync.WaitGroup
	wg.Add(grs)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	parte1 := nums[0:3]
	parte2 := nums[3:6]
	parte3 := nums[6:9]

	go func() {
		defer wg.Done()
		soma := 0
		for _, n := range parte1 {
			soma += n
		}
		ch <- soma
	}()

	go func() {
		defer wg.Done()
		soma := 0
		for _, n := range parte2 {
			soma += n
		}
		ch <- soma
	}()

	go func() {
		defer wg.Done()
		soma := 0
		for _, n := range parte3 {
			soma += n
		}
		ch <- soma
	}()

	go func() {
		defer wg.Done()
		soma := 0
		for _, n := range nums {
			soma += n
		}
		ch <- soma
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for parcial := range ch {
		fmt.Println("Recebi", parcial)

	}

}

