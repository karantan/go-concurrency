package channels

import "crypto/sha1"

type Animal struct {
	hash     []byte
	category string
}

// AnimalInspector type is used for DI
type AnimalInspector func(string) Animal

// InspectAnimals accepts an `AnimalInspector` function and a list of animals which
// needed to be inspected with the passed `AnimalInspector` function.
// This spawns as many goroutines as needed so be careful with it's usage.
// If you are worry about memory you should use `InspectAnimals` in
// `go-concurrency/async-workers/production.go`
func InspectAnimals(ai AnimalInspector, animals []string) (results []Animal) {
	resultStream := make(chan Animal, len(animals))

	for _, animal := range animals {
		go func(a string) {
			resultStream <- ai(a)
		}(animal)
	}

	// We could return chan and move this code to the e.g. main function, but then
	// we would have to deal with channels in the main function which IMO is not
	// that nice.
	for i := 0; i < len(animals); i++ {
		result := <-resultStream
		log.Infof("%s processed.", result.category)
		results = append(results, result)
	}

	return
}

// example of an AnimalInspector implementation
func simpleInspector(animal string) Animal {
	h := sha1.New()
	h.Write([]byte(animal))
	bs := h.Sum(nil)

	return Animal{
		hash:     bs,
		category: animal,
	}
}
