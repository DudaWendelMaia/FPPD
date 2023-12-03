package main

import (
	"fmt"
)

func processoConcorrente(i int, c chan int) {
	fmt.Println(i)
	c <- i
}

func concorrente() {
	fmt.Println("inicio concorrente")

	var ch_fim chan int = make(chan int)

	for i := 0; i < 10; i++ {
		go processoConcorrente(i, ch_fim)

	}

	for i := 0; i < 10; i++ {
		fmt.Println(" ", <-ch_fim)
	}

	fmt.Println("fim concorrente")
}

func main() {
	concorrente()
}
