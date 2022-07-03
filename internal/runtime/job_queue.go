package runtime

import (
	"fmt"
	"visualWorkflows/shared/entities"
	wc "visualWorkflows/workerclient"
)

type JobQueue struct {
	queue   []entities.Job
	runtime *Runtime

	jobResults chan entities.JobResult
}

func constructJobQueue(runtime *Runtime) (*JobQueue, error) {
	jq := JobQueue{
		runtime:    runtime,
		queue:      make([]entities.Job, 0),
		jobResults: make(chan entities.JobResult),
	}

	go jq.handleJobResult()

	return &jq, nil
}

func (jq *JobQueue) addJob(job entities.Job) error {

	jq.queue = append(jq.queue, job)
	fmt.Println("Added job", job)

	go jq.runJobs()

	return nil
}

func (jq *JobQueue) runJobs() {
	for _, job := range jq.queue {
		worker := jq.selectWorker()

		go worker.ProcessJob(job, jq.jobResults)
	}
}

func (jq *JobQueue) handleJobResult() {
	result := <-jq.jobResults
	event := JobResultReceivedBody{
		Result: result,
	}
	jq.runtime.EventStreamer.invokeEvent(
		buildJobResultReceivedEvent(event),
	)
}

// TODO: Implement a valid strategy
func (jq *JobQueue) selectWorker() wc.WorkerClient {
	return *jq.runtime.Workers[0]
}
