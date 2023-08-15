package mappers

import (
	"github.com/zerefwayne/rooots/server/dto/strava"
	"github.com/zerefwayne/rooots/server/models"
)

func MapSummaryAthleteToUser(summaryAthlete *strava.SummaryAthlete) *models.User {
	return &models.User{
		StravaId:      summaryAthlete.Id,
		FirstName:     summaryAthlete.FirstName,
		LastName:      summaryAthlete.LastName,
		ProfileMedium: summaryAthlete.ProfileMedium,
		Profile:       summaryAthlete.Profile,
		City:          summaryAthlete.City,
		State:         summaryAthlete.State,
		Country:       summaryAthlete.Country,
		Sex:           summaryAthlete.Sex,
		Summit:        summaryAthlete.Summit,
		CreatedAt:     summaryAthlete.CreatedAt,
		UpdatedAt:     summaryAthlete.UpdatedAt,
	}
}
