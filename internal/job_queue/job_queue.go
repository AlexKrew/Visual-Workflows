package job_queue

import (
	"errors"
	"log"
	"sync"
	"time"
	"workflows/shared/shared_entities"
)

type JobQueue struct {
	lockMutex sync.Mutex

	// list of jobs that have to be processed.
	// the jobs are stored inside a inverted index with the job-type as a key
	jobs map[string][]shared_entities.Job
}

func NewJobQueue() (*JobQueue, error) {
	queue := &JobQueue{
		jobs: make(map[string][]shared_entities.Job),
	}

	return queue, nil
}

func (queue *JobQueue) AddJob(job shared_entities.Job) bool {
	queue.lockMutex.Lock()
	defer queue.lockMutex.Unlock()

	if _, ok := queue.jobs[job.Type]; !ok {
		queue.jobs[job.Type] = []shared_entities.Job{}
	}

	queue.jobs[job.Type] = append(queue.jobs[job.Type], job)
	log.Println("Added job", job)

	return true
}

func (queue *JobQueue) LockJob(jobId shared_entities.JobID) (bool, error) {
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
	go queue.UnlockJob(jobId, 10*time.Second)

	// TODO: Emit Job Locked Event?

	return true, nil
}

func (queue *JobQueue) UnlockJob(jobId shared_entities.JobID, timeout time.Duration) (bool, error) {

	// await the timeout duration
	time.Sleep(timeout * time.Second)

	queue.lockMutex.Lock()
	defer queue.lockMutex.Unlock()

	job, exists := queue.JobById(jobId)
	if !exists {
		return false, errors.New("job with id does not exist")
	}

	if !job.Locked {
		return false, nil
	}

	job.Locked = false
	return true, nil
}

func (queue *JobQueue) RemoveJob(jobId shared_entities.JobID) (shared_entities.Job, bool) {
	queue.lockMutex.Lock()
	defer queue.lockMutex.Unlock()

	var removedJob shared_entities.Job

	job, exists := queue.JobById(jobId)
	if !exists {
		return removedJob, false
	}

	// find index of job
	for index := range queue.jobs[job.Type] {

		if queue.jobs[job.Type][index].ID == jobId {

			job := queue.jobs[job.Type][index]

			// job found -> remove job
			jobs := []shared_entities.Job{}
			jobs = append(jobs, queue.jobs[job.Type][:index]...)
			jobs = append(jobs, queue.jobs[job.Type][index+1:]...)
			queue.jobs[job.Type] = jobs

			return job, true
		}
	}

	return removedJob, false
}

func (queue *JobQueue) JobById(id shared_entities.JobID) (*shared_entities.Job, bool) {
	for jobType := range queue.jobs {
		for _, job := range queue.jobs[jobType] {
			if job.ID == id {
				return &job, true
			}
		}
	}

	return nil, false
}
