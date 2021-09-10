package main

import (
	"fmt"
	ssh "go-concurrency/async-workers"
	"go-concurrency/logger"
	"time"
)

var log = logger.New("main")

var COMMAND = "uname -a"

func main() {
	defer timeTrack(time.Now(), "main")

	serverHostnames := []string{"foo", "bar"}
	var jobs []ssh.Job

	for i, hostname := range serverHostnames {
		// Run SSH command:
		jobs = append(jobs, ssh.Job{
			JobNumber: i,
			Hostname:  hostname,
			Command:   fmt.Sprintf("%s %s", hostname, COMMAND),
			App:       "ssh",
		})
	}

	results := ssh.Run(ssh.GenericJobRunner, jobs)
	for _, r := range results {
		log.Infof("[%d/%d %s] %s", r.JobNumber+1, len(results), r.Hostname, r.Result)
	}

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Infof("%s took %s", name, elapsed)
}
