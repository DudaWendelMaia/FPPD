package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("Escolha uma das opções abaixo:")
	fmt.Println("1 - Filósofos")
	fmt.Println("2 - Leitores e Escritores")
	fmt.Println("3 - Papai Noel")
	fmt.Println("--------------------------------------------------")

	var option int
	fmt.Scanln(&option)

	fmt.Println("--------------------------------------------------")
	fmt.Println("Escolha o tempo de duração da simulação em segundos:")
	fmt.Println("--------------------------------------------------")

	var delayValue int
	fmt.Scanln(&delayValue)

	switch option {
	case 1:
		Init("filósofos", delayValue, aFilosofos)
	case 2:
		Init("leitores e escritores", delayValue, dLeitoresEscritores)
	case 3:
		Init("papai noel", delayValue, gSantaClaus)
	default:
		fmt.Println("Opção inválida")
	}
}

func Init(
	name string,
	delayValue int,
	fn func(time.Duration),
) {
	countdown := 3 * time.Second
	delay := time.Duration(delayValue) * time.Second
	PrintStartFunc(name, delayValue)
	Countdown(countdown)
	fn(delay)
	PrintEndFunc(name)
}

func PrintStartFunc(title string, delayValue int) {
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Iniciando simulação %s durando %d segundo(s): \n", title, delayValue)
	fmt.Println("--------------------------------------------------")
}

func Countdown(t time.Duration) {
	for t > 0 {
		fmt.Println(t)
		t -= time.Second
		time.Sleep(time.Second)
	}
}

func PrintEndFunc(title string) {
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Fim da simulação %s\n", title)
	fmt.Println("--------------------------------------------------")
}
