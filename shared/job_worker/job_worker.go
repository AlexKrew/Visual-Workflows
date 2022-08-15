package job_worker

import "workflows/shared/shared_entities"

type JobWorker interface {
	JobType() string
	ProcessJob(job shared_entities.Job) shared_entities.JobResult
}
