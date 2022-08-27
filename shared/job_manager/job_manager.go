package job_manager

import (
	"errors"
	"workflows/shared/job_worker"
	"workflows/shared/shared_entities"
)

type JobManager struct {
	workers map[string]*job_worker.JobWorker
}

func NewJobManager() JobManager {
	return JobManager{
		workers: make(map[string]*job_worker.JobWorker),
	}
}

func (manager *JobManager) SupportedServices() []string {
	services := []string{}
	for serviceType := range manager.workers {
		services = append(services, serviceType)
	}

	return services
}

func (manager *JobManager) AddWorker(worker job_worker.JobWorker) {
	jobType := worker.JobType()
	manager.workers[jobType] = &worker
}

func (manager *JobManager) CanExecute(jobType string) bool {
	_, hasHandler := manager.workers[jobType]
	return hasHandler
}

func (manager *JobManager) Execute(job shared_entities.Job) (shared_entities.JobResult, error) {

	handler, ok := manager.workers[job.Type]
	if !ok {
		return shared_entities.JobResult{}, errors.New("manager missing matching handler")
	}

	result := (*handler).ProcessJob(job)
	return result, nil
}
