package around

import (
	"fmt"
	"net/http"
	"log"
	"bytes"
	"strings"
	"encoding/json"
)

func Connect(quit chan int) {
	res, err := http.Get("https://"+getApiKey()+"@streaming.campfirenow.com/room/568525/live.json")
	if err != nil {
		log.Fatal(err)
	}
	var buffer bytes.Buffer

	ln := []byte{13}

	stringLn := string(ln)

	for {
		b := []byte{10}
		_, err := res.Body.Read(b)

		if err != nil {
			log.Fatal(err)
		}

		s := string(b)

		if s == stringLn {
			line := buffer.String()

			if line != " " {
				dec := json.NewDecoder(strings.NewReader(line))

				var message Message

				err = dec.Decode(&message)

				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(message.Body)
			}
			buffer.Reset()
		} else {
			buffer.WriteString(s)
		}

	}
	quit <- 1
}


func Speak(words string) {
	xmlMessage := "<message><type>TextMessage</type><body>" + words +  "</body></message>"
	xmlReader := strings.NewReader(xmlMessage)


	_, err := http.Post("https://"+getApiKey()+"@workstars.campfirenow.com/room/568525/speak.xml", "text/xml", xmlReader)
	if err != nil {
		log.Fatal(err)
	}
}



type Message struct {
	Starred bool
	UserId int
	RoomId int
	Id int
	Body string
	Time string
}
