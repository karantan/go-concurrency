package sync

import (
	"sync"
	"time"
)

const WORKERS = 5

func runWorkers() {
	var wg sync.WaitGroup // initialise the goroutine "counter"

	for i := 1; i <= WORKERS; i++ {
		wg.Add(1) // add +1 to the counter
		go func(id int) {
			defer wg.Done() // add -1 to the counter
			log.Infof("Worker %d starting\n", id)

			time.Sleep(time.Second * 1) // Sleep to simulate an expensive task.
			log.Infof("Worker %d done\n", id)
		}(i)
	}
	wg.Wait() // wait for the counter to be 0 (i.e. wait for all goroutines to finish)
}
