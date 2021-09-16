package channels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockedInspector(a string) Animal {
	return Animal{[]byte(a), a}
}

func TestInspectAnimals(t *testing.T) {
	assert := assert.New(t)

	got := InspectAnimals(mockedInspector, []string{"dog", "cat"})

	// Because we process animals concurrently we can't rely on the order. Only that
	// the result contins what we want.
	assert.Contains(got, Animal{[]byte("dog"), "dog"})
	assert.Contains(got, Animal{[]byte("cat"), "cat"})

}
