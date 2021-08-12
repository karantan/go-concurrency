package channels

import "time"

// WriteToChannelUnbuffered shows an example how to write to a channel in a goroutine
func WriteToChannelUnbuffered() []string {
	animals := []string{"dog", "cat", "bird"}
	results := []string{}
	resultStream := make(chan string)

	for _, animal := range animals {
		go func(a string) {
			log.Infof("Processing %s ...", a)
			resultStream <- processAnimal(a) // only 1 item can be send to the stream at a time
			log.Infof("Done with %s", a)
		}(animal)
	}

	log.Infof("Getting results from the stream ...")
	time.Sleep(time.Millisecond * 1500)
	for i := 0; i < len(animals); i++ {
		// this like is blocking the goroutines from writing in the `resultStream`
		// because we are using unbuffered channel
		result := <-resultStream
		log.Info(result)
		results = append(results, result)
		time.Sleep(time.Millisecond * 500) // sleep 1sec to show the blocking part
	}

	return results
}

// WriteToChannelBuffered shows an example how to write to a channel in a goroutine
func WriteToChannelBuffered() []string {
	animals := []string{"dog", "cat", "bird"}
	results := []string{}
	resultStream := make(chan string, 3)

	for _, animal := range animals {
		go func(a string) {
			log.Infof("Processing %s ...", a)
			resultStream <- processAnimal(a) // 3 items can be send to the stream at a time
			log.Infof("Done with %s", a)
		}(animal)
	}

	log.Infof("Getting results from the stream ...")
	time.Sleep(time.Millisecond * 1500)
	for i := 0; i < len(animals); i++ {
		// not blocking anymore because all the items are already processed because we
		// used buffered channel this time.
		result := <-resultStream
		log.Info(result)
		results = append(results, result)
		time.Sleep(time.Millisecond * 500) // sleep 1sec to show the blocking part
	}

	return results
}
