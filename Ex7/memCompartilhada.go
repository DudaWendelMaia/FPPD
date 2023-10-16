package main

import "fmt"

var sharedTest int = 0 
var ch_fim chan struct{} = make(chan struct{})

func MyFunc() {
	for k :=0; k < 100; k++ {
		sharedTest = sharedTest + 1
	}
	ch_fim <- struct{}{}
}

func main() {
	for i := 0; i<100; i++ {
	go MyFunc()
	}
	fmt.Println("Criei 100 processos")

	for i := 0; i<100; i++ {
		<-ch_fim
	}
	fmt.Println("Todos os processos terminaram!Resultado " sharedTest)

}