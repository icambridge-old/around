/**
 *
 *
 */

package main

import (
	"fmt"
	"around"
)


func main() {
	fmt.Println("Starting up the campfire")

	routineQuit := make(chan int)
	go around.Connect(routineQuit)
	fmt.Println("We're now sitting around the campfire")

   	go around.CheckTrains()
	go around.StartListener()

	//around.NotifyOtherPeople("Hello this is a test\nSo this should be Good\n")

	<-routineQuit // blocks until quit is written to
}



