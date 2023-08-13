package repository

import (
	"log"

	"github.com/google/uuid"
	strava "github.com/zerefwayne/rooots/server/dto/strava"
	"github.com/zerefwayne/rooots/server/models"
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

func FindOrCreateUserByStrava(DB *gorm.DB, athlete *strava.SummaryAthlete) (*models.User, error) {
	foundUser, err := FindUserByStravaId(DB, athlete.Id)
	if err != nil {
		// Cannot find user
		createdUser, createUserErr := CreateUserByStravaLogin(DB, athlete)
		return createdUser, createUserErr
	}

	return foundUser, err
}
