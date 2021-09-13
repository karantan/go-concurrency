package async_workers

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type Job struct {
	JobNumber int
	Hostname  string
	App       string
	Command   string
}

type JobResult struct {
	JobNumber int
	Hostname  string
	Result    string
	ExitCode  int
}

// JobRunner type is used for DI
type JobRunner func(Job) JobResult

// The number of concurrent workers. More workers = more memory consumed & faster all
// jobs are finished.
const SSH_JOBS = 10

func Run(jr JobRunner, jobs []Job) (results []JobResult) {
	jobStream := make(chan Job, len(jobs))
	resultStream := make(chan JobResult, len(jobs))

	wg := new(sync.WaitGroup)
	for i := 0; i < SSH_JOBS; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for job := range jobStream {
				resultStream <- jr(job)
			}
		}()
	}

	for _, j := range jobs {
		jobStream <- j
	}
	close(jobStream)
	wg.Wait()

	for i := 0; i < len(jobs); i++ {
		result := <-resultStream
		results = append(results, result)
	}

	return
}

// GenericJobRunner executes a command described in a job.
func GenericJobRunner(j Job) JobResult {
	// First make sure we can connect to the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := exec.CommandContext(ctx, "ssh", j.Hostname).Run(); err != nil {
		return JobResult{
			JobNumber: j.JobNumber,
			Hostname:  j.Hostname,
			Result:    fmt.Sprintf("Can't connect to %s.", j.Hostname),
		}
	}

	command := strings.Split(j.Command, " ")
	cmd := exec.Command(j.App, command...)
	log.Infof("Executing command `%s` on %s ...", cmd.String(), j.Hostname)
	out, err := cmd.CombinedOutput()

	var ee *exec.ExitError
	exitCode := 0
	if errors.As(err, &ee) {
		exitCode = ee.ExitCode()
	}

	return JobResult{
		JobNumber: j.JobNumber,
		Hostname:  j.Hostname,
		Result:    strings.TrimRight(string(out), "\n"),
		ExitCode:  exitCode,
	}
}
