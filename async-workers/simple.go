package async_workers

import (
	"sync"
	"time"
)

const WORKERS = 3

func processAnimal(animal string, taskId int) string {
	log.Infof("Processing animal %s (taskID: %d)...\n", animal, taskId)
	time.Sleep(time.Second * 1) // Sleep to simulate an expensive task.
	return animal
}

func Simple1() {
	animals := []string{"dog", "cat", "bird", "monkey", "fish", "snake", "whale"}

	// chan for gathering data out of goroutines. it needs to be buffered otherwise it
	// will block inside for loop when writing to it because reading
	// from this chan isn't fired yet (because we do `wg.Wait` before doing the reading)
	resultStream := make(chan string, len(animals))
	// chan for sending the data to the goroutines
	// this can be unbufferd because we initialised the sender and receiver together.
	animalStream := make(chan string)

	var wg sync.WaitGroup
	for i := 1; i <= WORKERS; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done() // On return, notify the WaitGroup that we're done.
			// Get animals from the channel. Don't forget to close the `animalStream`
			// when you are done sending the data in or else you will have a deadlock
			// here.
			for a := range animalStream { // blocking operation the the chan is empty or if the chan is not closed
				log.Infof("Worker %d starting\n", id)
				resultStream <- processAnimal(a, id)
				log.Infof("Worker %d done\n", id)
			}
		}(i)
	}

	log.Infof("Goroutines get ready for incoming data ...")
	time.Sleep(time.Second * 2)
	for _, animal := range animals {
		// because `animalStream` is unbuffered channel we can only send 1 item in at
		// a time. But because we are using 3 workers we can kinda send 4 items in at
		// the start because each worker takes 1 item out of the stream and 1 can stay
		// in the unbuffered channel
		log.Infof("Sending %s for processing ...", animal)
		animalStream <- animal
	}

	// Close the `animalStream` chan to indicate that we are done sending the data.
	// It needs to be closed before we wait for all the workers to finish
	// (i.e. before `wg.Wait()`)
	close(animalStream)

	wg.Wait() // Wait for all goroutines to finish their work.

	// Gather the data
	for i := 0; i < len(animals); i++ {
		result := <-resultStream // blocking operation if the channel is empty
		log.Infof("We have the processed data ready: %s\n", result)
	}

	// Hey, close the `resultStream` chan (i.e. `close(resultStream)`)!
	// You needn't close every channel when you've finished with it. It's only necessary
	// to close a channel when it's important to tell the receiving goroutines that all
	// data have been sent.
}
