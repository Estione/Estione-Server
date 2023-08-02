package domain

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Uuid uuid.UUID

	Sender   uuid.UUID
	Receiver uuid.UUID

	Text string

	Created time.Time
	Updated time.Time
}
