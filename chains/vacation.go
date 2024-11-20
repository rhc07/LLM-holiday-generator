package chains

import (
	"errors"
	"slices"

	"github.com/google/uuid"
)

var Vacations []*Vacation

func GetVacationFromDb(id uuid.UUID) (Vacation, error) {
	// Use the slices package to find the index of the object with
	// matching ID in the database. If it does not exist, this will return
	// -1
	idx := slices.IndexFunc(Vacations, func(v *Vacation) bool { return v.ID == id })

	// If the ID didn't exist, return an error and let the caller
	// handle it
	if idx < 0 {
		return Vacation{}, errors.New("ID Not Found")
	}

	return *Vacations[idx], nil
}

func GenerateVacationIdeaChange(id uuid.UUID, budget int, idea string, hobbies []string) {

}
