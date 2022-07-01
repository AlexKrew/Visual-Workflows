package runtime

import "github.com/google/uuid"

func getNewUUID() uuid.UUID {
	return uuid.New()
}
