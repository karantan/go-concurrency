package async_workers

import (
	"crypto/sha1"
	"sync"
)

// AnimalInspector type is used for DI
type AnimalInspector func(string) Animal

const JOBS = 3

type Animal struct {
	hash     []byte
	category string
}

// InspectAnimals accepts an `AnimalInspector` function and a list of animals which
// needed to be inspected with the passed `AnimalInspector` function.
// It does this concurently with max `JOBS` workers so that we don't consume too much
// resources.
func InspectAnimals(ai AnimalInspector, animals []string) (results []Animal) {
	animalStream := make(chan string, len(animals))
	resultStream := make(chan Animal, len(animals))

	wg := new(sync.WaitGroup)
	for i := 0; i < JOBS; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for animal := range animalStream {
				resultStream <- ai(animal)
			}
		}()
	}

	for _, a := range animals {
		animalStream <- a
	}
	close(animalStream)
	wg.Wait()

	for i := 0; i < len(animals); i++ {
		result := <-resultStream
		results = append(results, result)
	}

	return
}

// example of an AnimalInspector implementation
func simpleInspector(animal string) Animal {
	h := sha1.New()
	h.Write([]byte(animal))
	bs := h.Sum(nil)

	return Animal{bs, animal}
}
