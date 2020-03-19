package todo

import (
	"time"

	"github.com/google/uuid"
)

type TODO struct {
	ID          uuid.UUID
	Title, Text string
	CreateAt    time.Time
	_           struct{}
	_           [0]func()
}
