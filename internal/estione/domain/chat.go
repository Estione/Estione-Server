package domain

import (
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	ChatIdentifier uuid.UUID

	Members []uuid.UUID

	Created time.Time
}
