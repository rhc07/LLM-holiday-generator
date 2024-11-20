package chains

import "github.com/google/uuid"

type Vacation struct {
	ID        uuid.UUID `json:"Id"`
	Idea      string    `json:"idea"`
	Completed bool      `json:"completed"`
}
