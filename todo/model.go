package todo

import (
	"time"

	"github.com/google/uuid"
)

type TODO struct {
	_           struct{}
	_           [0]func()
	ID          uuid.UUID
	Title, Text string
	CreateAt    time.Time
}
