package utils

import "github.com/google/uuid"

type UUID = string

func GetNewUUID() string {
	return uuid.New().String()
}
