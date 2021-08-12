package async_workers

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
	assert.Equal(got, []Animal{
		{[]byte("dog"), "dog"},
		{[]byte("cat"), "cat"},
	})
}
