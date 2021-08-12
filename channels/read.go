package channels

import (
	"fmt"
	"time"
)

func processAnimal(animal string) string {
	time.Sleep(time.Second * 1) // Sleep to simulate an expensive task.
	return fmt.Sprintf("%s is processed", animal)
}

// readFromChannel shows an example how to read from a channel in a goroutine
func ReadFromChannel() {
	animals := []string{"dog", "cat", "bird"}
	inStream := make(chan string)

	go func() {
		for animal := range inStream {
			log.Infof("Processing %s ...", animal)
			processAnimal(animal)
		}
	}()

	log.Info("Goroutine get ready ...")
	for _, animal := range animals {
		time.Sleep(time.Second * 1)
		log.Infof("Sending %s to get processed", animal)
		inStream <- animal
	}
}
