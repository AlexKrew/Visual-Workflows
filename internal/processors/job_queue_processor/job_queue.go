package job_queue_processor

import (
	"errors"
	"sync"
	"workflows/internal/client"
	"workflows/internal/workflows"
)

type JobQueue struct {
	lockMutex sync.Mutex

	// list of jobs that have to be processed.
	// the jobs are stored inside a inverted index with the job-type as a key
	jobs   map[string][]workflows.Job
	client *client.Client
}

func ConstructJobQueue() (*JobQueue, error) {
	queue := &JobQueue{
		jobs: make(map[string][]workflows.Job),
	}

	return queue, nil
}

func (queue *JobQueue) AddJob(job workflows.Job) {
	queue.lockMutex.Lock()
	defer queue.lockMutex.Unlock()

	if _, ok := queue.jobs[job.NodeType]; !ok {
		queue.jobs[job.NodeType] = []workflows.Job{}
	}

	queue.jobs[job.NodeType] = append(queue.jobs[job.NodeType], job)
}

func (queue *JobQueue) ExecuteJob(jobId workflows.JobID) (client.JobResults, error) {

	job, exists := queue.RemoveJob(jobId)
	if !exists {
		return client.JobResults{}, errors.New("job with this id does not exist")
	}

	messages := make(map[string]client.Message)
	for key, msg := range job.Input {
		messages[key] = client.Message{
			Datatype: workflows.MessageTypeToString(msg.DataType),
			Value:    msg.Value,
		}
	}

	clientJob := client.Job{
		ID:    job.ID,
		Type:  job.NodeType,
		Input: messages,
	}
	return queue.client.DoJob(clientJob)
}

func (queue *JobQueue) LockJob(jobId workflows.JobID) (bool, error) {
	queue.lockMutex.Lock()
	defer queue.lockMutex.Unlock()

	job, exists := queue.JobById(jobId)
	if !exists {
		return false, errors.New("job with jobid does not exist")
	}

	if job.Locked {
		return false, nil
	}

	job.Locked = true
	// TODO: Start unlock interval

	// TODO: Emit Job Locked Event?

	return true, nil
}

func (queue *JobQueue) RemoveJob(jobId workflows.JobID) (workflows.Job, bool) {
	queue.lockMutex.Lock()
	defer queue.lockMutex.Unlock()

	job, exists := queue.JobById(jobId)
	if !exists {
		return workflows.Job{}, false
	}

	// find index of job
	for index := range queue.jobs[job.NodeType] {

		if queue.jobs[job.NodeType][index].ID == jobId {

			job := queue.jobs[job.NodeType][index]

			// job found -> remove job
			jobs := []workflows.Job{}
			jobs = append(jobs, queue.jobs[job.NodeType][:index]...)
			jobs = append(jobs, queue.jobs[job.NodeType][index+1:]...)
			queue.jobs[job.NodeType] = jobs

			return job, true
		}
	}

	return workflows.Job{}, false
}

func (queue *JobQueue) JobById(id workflows.JobID) (*workflows.Job, bool) {
	for jobType := range queue.jobs {
		for _, job := range queue.jobs[jobType] {
			if job.ID == id {
				return &job, true
			}
		}
	}

	return nil, false
}
