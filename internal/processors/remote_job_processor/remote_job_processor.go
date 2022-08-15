package remote_job_processor

import (
	"workflows/internal/job_queue"
	"workflows/internal/workflows"
)

type RemoteJobProcessor struct {
	eventStream *workflows.EventStream
	jobQueue    *job_queue.JobQueue
}
