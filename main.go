package main

// Add imports.
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
		fmt.Println("Recebi", parcial) // o main tem que fazer algo sobre os canais reproduzidos

	}

}

// 1. close só pode ser chamado UMA vez

// 2. close tem que vir DEPOIS de todos os envios (em uma função)

// 2.5 quem espera e fecha é uma goroutine separada

// 1. main cria ch e wg
// 2. main lança 4 goroutines produtoras (wg.Add(4))
// 3. main lança 1 goroutine fechadora (que dorme no wg.Wait)
// 4. main entra no for range ch (espera valores)
// 5. As 4 produtoras enviam valores pro canal
// 6. Cada uma chama wg.Done() ao terminar
// 7. Contador chega a 0 → fechadora chama close(ch)
// 8. for range vê canal fechado → termina
// 9. main termina → programa acaba
