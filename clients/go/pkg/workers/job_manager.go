package workers

import (
	"errors"
	"workflows/clients/go/pkg/entities"
	"workflows/clients/go/pkg/nodes"
	pb "workflows/gateway"
)

type JobManager struct {
	workers map[string]nodes.HandlerFunc
}

func NewJobManager() JobManager {
	return JobManager{
		workers: make(map[string]nodes.HandlerFunc),
	}
}

func (manager *JobManager) AddWorker(worker JobWorker) error {

	if _, exists := manager.workers[worker.JobType]; exists {
		return errors.New("worker for jobtype already exist")
	}

	manager.workers[worker.JobType] = worker.Handler

	return nil
}

func (manager *JobManager) ProcessJob(job *pb.ActivateJobResponse) (*pb.CompleteJobRequest, error) {

	handlerFunc, exist := manager.workers[job.Job.Type]
	if !exist {
		return nil, errors.New("no handler for jobtype registered")
	}

	activatedJob := entities.ActivatedJob{Job: job.Job}

	input := nodes.NodeInputFromJob(activatedJob)
	output := nodes.NewNodeOutput()

	// here the actual execution of the node logic is happening.
	// after calling the handlerFunc, output contains the results of the job
	handlerFunc(&input, &output)

	jobOutput := make(map[string]any)
	jobOutput["output"] = output

	// jobResult := pb.CompleteJobRequest{
	// 	JobId: job.Job.JobId,
	// 	Output: ,
	// }
	return &pb.CompleteJobRequest{
		JobId:  job.Job.JobId,
		Output: "",
	}, nil
}
