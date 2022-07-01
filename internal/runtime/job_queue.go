package runtime

import "fmt"

type JobQueue struct {
	queue   []Job
	runtime *Runtime
}

func constructJobQueue(runtime *Runtime) (*JobQueue, error) {
	jq := JobQueue{
		runtime: runtime,
		queue:   make([]Job, 0),
	}

	return &jq, nil
}

func (jq *JobQueue) addJob(job Job) error {

	jq.queue = append(jq.queue, job)
	fmt.Println("Added job", job)

	return nil
}
