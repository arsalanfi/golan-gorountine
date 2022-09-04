package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(wg *sync.WaitGroup, pingChannel chan string, pongChannel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		pingChannel <- "World!"
		time.Sleep(1 * time.Second)
		<-pongChannel
	}
	wg.Done()

}
func sayWorld(wg *sync.WaitGroup, pingChannel chan string, pongChannel chan string) {

	for i := 0; i < 5; i++ {
		v := <-pingChannel
		fmt.Println(v)
		time.Sleep(1 * time.Second)
		pongChannel <- "pong"
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	pingChannel := make(chan string, 0)
	pongChannel := make(chan string)
	wg.Add(1)
	go sayHello(&wg, pingChannel, pongChannel)

	wg.Add(1)
	go sayWorld(&wg, pingChannel, pongChannel)

	wg.Wait()

	fmt.Println("Finished")
}
