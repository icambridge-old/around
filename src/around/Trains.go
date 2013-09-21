package around

import (
	"fmt"
	"time"
	"log"
	"net/http"
	"encoding/json"
)

func CheckTrains() {

	for {
		timeNow := time.Now()

		hour := timeNow.Hour()

		if hour == 8 {
			checkTrainsCall("arriving")
		} else if hour == 17 {
			checkTrainsCall("departing")
		} else {

			minute := int(timeNow.Minute())
			sleepUntil := int(60 - minute);
			time.Sleep(time.Duration(sleepUntil)*time.Minute)

			continue
		}
		time.Sleep(time.Minute * 5)
	}

}

func checkTrainsCall(verb string) {

	var url = "http://trainapi.gopagoda.com/"+verb+".php"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(res.Body)
	var alert TrainAlert
	err = dec.Decode(&alert)

	if err != nil {
		log.Fatal(err)
	}

	msg := fmt.Sprintf("The train %s East Didsbury at %s current status is %s", verb, alert.Time, alert.Status)

	Speak(msg)

	fmt.Println("Done stuff")
}

type TrainAlert struct {
	Time string
	Status string
}
