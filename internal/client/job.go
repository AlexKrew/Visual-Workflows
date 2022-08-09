package client

type JobType = string

type Job struct {
	ID    string
	Type  JobType
	Input map[string]Message
}

type JobResults struct {
	Output map[string]Message
	Logs   []any
}

type JobHandler func(job Job) JobResults
