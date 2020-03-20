package todo

import (
	"time"

	"github.com/google/uuid"
)

type TODO struct {
	_         struct{}
	_         [0]func()
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Create_at time.Time `json:"created_at"`
}
