package routes

import "github.com/google/uuid"

type GenerateVacationIdeaRequest struct {
	FavoriteSeason string   `json:"favoriteSeason"`
	Hobbies        []string `json:"hobbies"`
	Budget         int      `json:"budget"`
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
