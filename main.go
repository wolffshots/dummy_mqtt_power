package main

import (
	"fmt"
	"github.com/wolffshots/phocus_mqtt" // comms with mqtt broker
	"log"                               // console io
	"time"                              // for sleeping
)

func main() {
	// mqtt
	err := phocus_mqtt.Setup("192.168.88.124", "go_dummy_client'")
	if err != nil {
		log.Fatalf("Failed to set up mqtt with err: %v", err)
	}
	count := 0
	for {
		payload := 0
		if (count % 10) > 5 {
			payload = 2000 + count*2
		}
		err = phocus_mqtt.Send("dummy/stats/power", 0, false, fmt.Sprintf("%d", payload), 10)
		if err != nil {
            log.Fatalf("MQTT send of %d failed with: %v", err, payload)
		}
		time.Sleep(10 * time.Second)
		count++
	}
}
