package client

type JobWorker struct {
	JobType JobType
	Handler JobHandler
}

func NewJobWorker(jobType JobType, handler JobHandler) JobWorker {
	return JobWorker{
		JobType: jobType,
		Handler: handler,
	}
}
