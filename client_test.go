package cronicleapi

import (
	"log"
	"testing"
	"encoding/json"
)

func PrintObject(obj interface{}) {
	bytes,err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(string(bytes))
	}
}

func Test_GetEvent(t *testing.T) {
	cli := NicleClient {
		ApiUrl: "http://193.112.78.109:3012",
		ApiKey: "",
	}
	event,err := cli.GetEvent("ejhfjkj8a06","")
	PrintObject(event)

	// todo 
	event.Id = ""
	eventId,err := cli.CreateEvent(event)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("event Id:",eventId)

	err = cli.UpdateEvent(eventId,map[string]interface{}{
		"title": "Test updated",
	})
	if err != nil {
		log.Fatal(err)
	}

	jobIds,err := cli.RunEvent(eventId,"",nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(jobIds)

	status,err := cli.GetEventStatus(eventId)
	if err != nil {
		log.Fatal(err)
	}
	PrintObject(status)

	err = cli.DeleteEvent(eventId)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(err)
	log.Println(event)
}

func Test_Status(t *testing.T) {
	cli := NicleClient {
		ApiUrl: "http://193.112.78.109:3012",
		ApiKey: "",
	}

	eventId := "ejhfn5uh608"
	jobId := "jjhfn5unv09"
	status,err := cli.GetEventStatus(jobId)
	if err != nil {
		log.Fatal(err)
	}
	PrintObject(status)

	err = cli.DeleteEvent(eventId)
	if err != nil {
		log.Fatal(err)
	}
}