package main

func main() {

	var numChan chan int = make(chan int, 10)

	for i := 0; i < 30; i++ {

		numChan <- i

	}

}

/*
fatal error: all goroutines are asleep - deadlock!
*/
