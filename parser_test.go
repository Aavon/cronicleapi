package cronicleapi

import (
	"log"
	"time"
	"testing"
)

func Test_parse(t *testing.T) {
	st := time.Now()
	log.Println(GenTimming(Once,st))
	log.Println(GenTimming(Day,st))
	log.Println(GenTimming(Month,st))
	log.Println(GenTimming(Year,st))
	log.Println(GenTimming(Week,st))
	log.Println(GenTimming(WorkDay,st))
}

