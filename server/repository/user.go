package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/zerefwayne/rooots/server/models"
	strava "github.com/zerefwayne/rooots/server/models/strava"
	"gorm.io/gorm"
)

func FindUserByStravaId(DB *gorm.DB, id uint64) (*models.User, error) {
	var foundUser models.User

	result := DB.Model(&models.User{}).Where(&models.User{StravaId: id}).First(&foundUser)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	return &foundUser, nil
}

func CreateUserByStravaLogin(DB *gorm.DB, summaryAthlete *strava.SummaryAthlete) (*models.User, error) {
	newUser := models.User{
		Id:        uuid.New(),
		StravaId:  summaryAthlete.Id,
		FirstName: summaryAthlete.FirstName,
		LastName:  summaryAthlete.LastName,
	}

	result := DB.Model(&models.User{}).Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}
