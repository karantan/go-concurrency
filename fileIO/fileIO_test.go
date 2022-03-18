package fileIO

import (
	"io/ioutil"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addToTooManyRequestsList(t *testing.T) {
	listPath := "TooManyRequestsList.domains"
	host := "foo.com"
	defer os.Remove(listPath)

	var m sync.RWMutex
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			exists, _ := existsInFile(&m, listPath, host)
			if !exists {
				writeToFile(&m, listPath, host)
			}
		}(&wg)
	}
	wg.Wait()

	listContent, err := ioutil.ReadFile(listPath)
	if err != nil {
		log.Fatal(err)
	}
	content := string(listContent)
	assert.Equal(
		t, content, "foo.com\n",
	)
}
