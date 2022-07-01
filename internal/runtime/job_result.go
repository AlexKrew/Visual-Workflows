package runtime

import "github.com/google/uuid"

type JobResult struct {
	ID   uuid.UUID
	Body any
}
