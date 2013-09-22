package around

import (
	"net"
	"log"
	"fmt"
	"encoding/json"
)

func StartListener() {
	l, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {

			//io.Copy(c, c)
			dec := json.NewDecoder(c)
			var job Job
			err = dec.Decode(&job)

			if err != nil {
				log.Fatal(err)
			}

			if job.Build.Phase == "STARTED" {
				Speak(fmt.Sprintf("Jenkins : The %s build has started", job.Name))
			} else if job.Build.Phase == "COMPLETED" {
				Speak(fmt.Sprintf("Jenkins : The %s build has %s", job.Name, job.Build.Status))
			}
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}


type Job struct {
	Name string
	Url string
	Build Build

}


type Build struct {
	Number int
	Phase string
	Status string
	Url string
}
