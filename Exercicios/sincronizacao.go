package main

import (
	"fmt"
)

const N = 40

var ch [N]chan struct{}

func funcaoA(id int, s string, in chan struct{}, out chan struct{}) {
	for {
		<-in
		fmt.Println(s, id)
		out <- struct{}{}
	}
}

func main() {
	for i := 0; i < N; i++ {
		ch[i] = make(chan struct{})
	}
	for i := 0; i < N; i++ {
		s := " "
		for j := 0; j < i; j++ {
			s = s + " "
		}
		go funcaoA(i, s, ch[i], ch[(i+1)%N])
	}
	ch[0] <- struct{}{}

	blq := make(chan struct{})
	<-blq
}
