package client

import (
	"errors"
)

type Client struct {
	Workers map[JobType]JobHandler
}

func NewClient() (Client, error) {
	workers := make(map[JobType]JobHandler)

	client := Client{
		Workers: workers,
	}

	injectWorker := NewJobWorker("Inject", handleInjectNodeJob)
	httpWorker := NewJobWorker("HTTP", handleHttpNodeJob)
	debugWorker := NewJobWorker("Debug", handleDebugNodeJob)

	client.AddWorker(injectWorker)
	client.AddWorker(httpWorker)
	client.AddWorker(debugWorker)

	return client, nil
}

func (client *Client) AddWorker(worker JobWorker) error {
	if _, exists := client.Workers[worker.JobType]; exists {
		return errors.New("worker for jobtype already added")
	}

	client.Workers[worker.JobType] = worker.Handler

	return nil
}

func (client *Client) DoJob(job any) (JobResults, error) {
	jobEntity := job.(Job)

	handler, exists := client.Workers[jobEntity.Type]
	if !exists {
		return JobResults{}, errors.New("no handler for jobtype registered")
	}

	return handler(jobEntity), nil
}

func handleInjectNodeJob(job Job) JobResults {

	input := NewNodeInput(job.Input)
	output := NewNodeOutput()
	ctx := NewNodeContext()

	ProcessInject(&input, &output, &ctx)

	results := JobResults{
		Output: output.Values,
		Logs:   ctx.Logs,
	}

	return results
}

func handleHttpNodeJob(job Job) JobResults {
	input := NewNodeInput(job.Input)
	output := NewNodeOutput()
	ctx := NewNodeContext()

	ProcessHttp(&input, &output, &ctx)

	results := JobResults{
		Output: output.Values,
		Logs:   ctx.Logs,
	}

	return results
}

func handleDebugNodeJob(job Job) JobResults {
	input := NewNodeInput(job.Input)
	output := NewNodeOutput()
	ctx := NewNodeContext()

	ProcessDebug(&input, &output, &ctx)

	results := JobResults{
		Output: output.Values,
		Logs:   ctx.Logs,
	}

	return results
}
