package fileIO

/*
Let's say we have a file where we keep a list of IPs we want to block if they make too
many requests (i.e. DDoS).

The part of the code that is listening for the traffic runs in goroutines so adding IPs
to the list will be triggered by another goroutine.

We need to make sure to create a exclusive acces to this shared resource. We can do this
with mutex (mutual exclusion).
*/
import (
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

// observer
func existsInFile(m *sync.RWMutex, filePath, host string) (bool, error) {
	m.Lock()
	defer m.Unlock()
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false, err
	}
	content := string(fileContent)
	if strings.Contains(content, host) {
		return true, nil
	}
	return false, nil
}

// producer
func writeToFile(m *sync.RWMutex, filePath string, host string) {
	m.Lock()
	defer m.Unlock()
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		log.Errorf("Error opening %s: %s", filePath, err)
	}

	if _, err := file.WriteString(host + "\n"); err != nil {
		log.Errorf("Error writing %s: %s", filePath, err)
		return
	}

	file.Close()
}
