package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rhc07/simple-go-service/chains"
)

func generateVacationIdea(r GenerateVacationIdeaRequest) GenerateVacationIdeaResponse {
	id := uuid.New()

	go chains.GenerateVacationIdeaChange(id, r.Budget, r.Weather, r.Hobbies, r.TravellingMonth, r.FlyingFrom, r.FlyingTime)

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

func GetVacationRouter(router *gin.Engine) *gin.Engine {

	// add the "vactions" group to the router
	registrationRoutes := router.Group("/vacations")

	// handle the POST request to generate a vacation idea
	registrationRoutes.POST("/create", func(c *gin.Context) {
		var req GenerateVacationIdeaRequest
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		} else {
			c.JSON(http.StatusOK, generateVacationIdea(req))
		}
	})

	// handle the GET request to retrieve a vacation idea
	registrationRoutes.GET("/:id", func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		} else {
			resp, err := getVacation(id)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid parameter"})
			} else {
				c.JSON(http.StatusOK, resp)
			}

		}
	})
	// return the update router
	return router
}
