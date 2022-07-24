package utils

import "github.com/google/uuid"

type UUID = uuid.UUID

func GetNewUUID() uuid.UUID {
	return uuid.New()
}
