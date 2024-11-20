package routes

import (
	"github.com/google/uuid"
	"github.com/rhc07/simple-go-service/chains"
)

func generateVacationIdea(r GenerateVacationIdeaRequest) GenerateVacationIdeaResponse {
	id := uuid.New()

	go chains.GenerateVacationIdeaChange(id, r.Budget, r.FavoriteSeason, r.Hobbies)

	return GenerateVacationIdeaResponse{ID: id, Completed: false}
}

func getVacation(id uuid.UUID) (GenerateVacationIdeaResponse, error) {
	// get the vaction idea from the db
	vacation, err := chains.GetVacationFromDb(id)

	if err != nil {
		return GenerateVacationIdeaResponse{}, err
	}

	return GenerateVacationIdeaResponse{ID: vacation.ID, Idea: vacation.Idea, Completed: vacation.Completed}, nil
}
