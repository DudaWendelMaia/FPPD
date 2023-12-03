package main

import (
	"fmt"
	"time"
)

var (
	// Amount of elves currently waiting for Santa
	elves     uint8 = 0
	// Amount of reindeers currently waiting for Santa
	reindeers uint8 = 0
	// Semaphore(0) - Santa waits for either 3 elves or 9 reindeers
	santaSem = NewSemaphore(0)
	// Semaphore(0) - Reindeers wait for Santa to hitch them
	reindeerSem = NewSemaphore(0)
	// Semaphore(1) - Elves wait for other elves to be helped
	elfTex = NewSemaphore(1)
	// Semaphore(1) - Mutex for accessing reindeers and elves counters
	mutex = NewSemaphore(1)
)

func santa() {
	for {
		PrintSanta("Santa is sleeping")
		santaSem.Wait()
		PrintSanta("Santa is awake")
		PrintSanta("Santa is waiting to help elves or reindeers")
		mutex.Wait()
		PrintSanta("Santa is ready to help")
			if reindeers == 9 {
				prepareSleigh()
				for i := 0; i < 9; i++ {
					PrintSanta(fmt.Sprintf("Hitching reindeer %d", i))
					reindeerSem.Signal()
				}
			} else if elves == 3 {
				helpElves()
			}
		PrintSanta("Santa is done helping")
		mutex.Signal()
		PrintSanta("Santa is going to sleep")
	}
}

func prepareSleigh() {
	PrintSanta("Santa is preparing the sleigh")
	reindeers = 0
}

func helpElves() {
	PrintSanta("Santa is helping elves")
}

func reindeer(name string) {
	for {
		PrintReindeer(name + " is waiting")
		mutex.Wait()
		PrintReindeer(name + " is in the critical section")
			reindeers += 1
			if reindeers == 9 {
				PrintReindeer(name + " is waking santa")
				santaSem.Signal()
			}
		PrintReindeer(name + " is out the critical section")
		mutex.Signal()

		PrintReindeer(fmt.Sprintf("%s is waiting for hitch (%d/9)", name, reindeers))
		reindeerSem.Wait()
		getHitched(name)
	}
}

func getHitched(name string) {
	PrintReindeer(name + " is getting hitched")
}

func elf(id string) {
	for {
		elfTex.Wait()
		PrintElf(id + " is waiting")
		mutex.Wait()
		PrintElf(id + " is in the critical section")
			elves += 1
			if elves == 3 {
				PrintElf(fmt.Sprintf("%s is waking santa (%d/3)", id, elves))
				santaSem.Signal()
			} else {
				PrintElf(fmt.Sprintf("%s is waiting for other elves (%d/4)", id, elves))
				elfTex.Signal()
			}
		PrintElf(id + " is out the critical section")
		mutex.Signal()

		// Should wait for others
		getHelp(id)

		mutex.Wait()
		PrintElf(id + " is leaving")
			elves -= 1
			if elves == 0 {
				elfTex.Signal()
			}
		PrintElf(id + " is out")
		mutex.Signal()
	}
}

func getHelp(id string) {
	PrintElf(id + " is getting help")
}

func gSantaClaus(d time.Duration) {
	reindeerNames := []string{"Dasher", "Dancer", "Prancer", "Vixen", "Comet", "Cupid", "Donder", "Blitzen", "Rudolph"}
	elfAmount := 4
	fmt.Println("Christmas is coming")
	go santa()
	for i := 0; i < len(reindeerNames); i++ {
		fmt.Println("Making reindeer ", i)
		go reindeer(reindeerNames[i])
	}
	for i := 0; i < elfAmount; i++ {
		fmt.Println("Making elf ", i)
		go elf(fmt.Sprintf("%d", i))
	}
	<-time.After(d)
	goodbye := " Christmas is here "
	colors := []string{Red, Green, Yellow}
	for i := 0; i < len(goodbye); i++ {
		fmt.Print(colors[i%3] + "=")
	}
	fmt.Println()
	for i := 0; i < len(goodbye); i++ {
		fmt.Print(colors[i%3] + string(goodbye[i]))
	}
	fmt.Println()
	for i := 0; i < len(goodbye); i++ {
		fmt.Print(colors[i%3] + "=")
	}
	fmt.Println(Reset)
}

func PrintSanta(s string) {
	PrintRed(s)
}

func PrintReindeer(s string) {
	PrintYellow(s)
}

func PrintElf(s string) {
	PrintGreen(s)
}
