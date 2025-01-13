// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
)



func numberServer(incrementChannel, decrementChannel, resultChan chan bool, results chan int) {
    i := 0
    for {
        select {
            case <- incrementChannel:
                i++
            case <- decrementChannel:
                i--
            case <-resultChan:
                results <- i
                return
        }
    }

}

func incrementing(incrementChannel chan bool, done chan bool) {
    //TODO: increment i 1000000 times
    for j:=0; j < 1000000; j++{
        incrementChannel <- true
    }
    done <- true
}

func decrementing(decrementChannel chan bool, done chan bool) {
    //TODO: decrement i 1000000 times

    for j:=0; j < 1000000; j++{
        decrementChannel <- true
    }
    done <- true
}

func main() {
    // What does GOMAXPROCS do? What happens if you set it to 1?
    //GOMAXPROCS limits the number of threads which can be executed simultanoulsy, if I set it to 1 only one thread can be executed a time. 
    runtime.GOMAXPROCS(2)    

    incrementChannel := make(chan bool)
    decrementChannel := make(chan bool)
    results := make(chan int)
    resultChan := make(chan bool)
    done := make(chan bool)

	
    // TODO: Spawn both functions as goroutines
    go numberServer(incrementChannel, decrementChannel, resultChan, results)
	go incrementing(incrementChannel, done)
    go decrementing(decrementChannel, done)
    // We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
    // We will do it properly with channels soon. For now: Sleep.
    <-done
    <-done
    resultChan <- true

    i := <-results
    Println("The magic number is:", i)
}
