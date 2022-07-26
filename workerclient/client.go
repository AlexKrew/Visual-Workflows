package workerclient

import (
	"errors"
	"visualWorkflows/shared/entities"
	"visualWorkflows/shared/utils"
)

type WorkerClient struct {
	ID utils.UUID
}

func ConstructWorker() *WorkerClient {
	return &WorkerClient{
		ID: utils.GetNewUUID(),
	}
}

func (wc *WorkerClient) ProcessJob(job entities.Job, back chan entities.JobResult) {

	var results interface{}
	var err error

	if job.Type == entities.JTProcess {
		results, err = executeNode(job)
	}

	if err != nil {
		back <- entities.JobResultError(job, err)
		return
	}

	// TODO: Cleanup -> JobResult should have a isEmpty method
	if results != nil {
		back <- results.(entities.JobResult)
		return
	}

	back <- entities.JobResult{
		Errors: []error{errors.New("unknown job type")},
	}
}

func executeNode(job entities.Job) (entities.JobResult, error) {
	return executeNodeJobs(job)
}
