package routes

import "github.com/google/uuid"

type GenerateVacationIdeaRequest struct {
	Weather         string   `json:"weather"`
	Hobbies         []string `json:"hobbies"`
	Budget          int      `json:"budget"`
	TravellingMonth string   `json:"travellingMonth"`
	FlyingFrom      string   `json:"flyingFrom"`
	FlyingTime      int      `json:"flyingTime"`
}

type GenerateVacationResponse struct {
	ID        uuid.UUID `json:"id"`
	Completed bool      `json:"completed"`
}

type GenerateVacationIdeaResponse struct {
	ID        uuid.UUID `json:"id"`
	Completed bool      `json:"completed"`
	Idea      string    `json:"idea"`
}
