# Soma com Channels em Go

Exercicio: somar um slice de inteiros dividindo o
trabalho entre várias goroutines, comunicando os resultados através de um channel.

## O que o programa faz

Dado um slice `[1, 2, 3, 4, 5, 6, 7, 8, 9]`, o programa:

1. Divide o slice em 3 partes: `[1 2 3]`, `[4 5 6]`, `[7 8 9]`
2. Lança 3 goroutines, cada uma soma uma parte
3. Lança uma 4ª goroutine que soma o slice inteiro (pra confirmar que as partes batem)
4. Cada goroutine envia sua soma parcial pro channel
5. O `main` recebe os valores e imprime cada um

## Para rodar:

```bash
go run main.go
