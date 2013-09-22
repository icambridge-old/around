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

	<-routineQuit // blocks until quit is written to
}



