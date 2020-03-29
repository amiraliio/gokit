package helper

import (
	"github.com/google/uuid"
)

type UUID uuid.UUID

func NewUUID() uuid.UUID {
	return uuid.New()
}
