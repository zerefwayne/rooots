package repository

import (
	"log"

	"github.com/google/uuid"
	strava "github.com/zerefwayne/rooots/server/dto/strava"
	"github.com/zerefwayne/rooots/server/models"
	"gorm.io/gorm"
)

func FindUserById(DB *gorm.DB, id uuid.UUID) (*models.User, error) {
	var foundUser models.User

	result := DB.Model(&models.User{}).Where(&models.User{Id: id}).First(&foundUser)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	return &foundUser, nil
}

func FindUserByStravaId(DB *gorm.DB, id uint64) (*models.User, error) {
	var foundUser models.User

	result := DB.Model(&models.User{}).Where(&models.User{StravaId: id}).First(&foundUser)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	}

	return &foundUser, nil
}

func CreateUserByStravaLogin(DB *gorm.DB, exchangeTokenBody *strava.ExchangeTokenResponseBody) (*models.User, error) {
	newUser := models.User{
		Id:           uuid.New(),
		StravaId:     exchangeTokenBody.Athlete.Id,
		FirstName:    exchangeTokenBody.Athlete.FirstName,
		LastName:     exchangeTokenBody.Athlete.LastName,
		RefreshToken: exchangeTokenBody.RefreshToken,
	}

	result := DB.Model(&models.User{}).Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}

func FindOrCreateUserByStrava(DB *gorm.DB, exchangeTokenBody *strava.ExchangeTokenResponseBody) (*models.User, error) {
	foundUser, err := FindUserByStravaId(DB, exchangeTokenBody.Athlete.Id)
	if err != nil {
		// Cannot find user
		createdUser, createUserErr := CreateUserByStravaLogin(DB, exchangeTokenBody)
		return createdUser, createUserErr
	}

	return foundUser, err
}

func UpdateRefreshToken(DB *gorm.DB, oldRefreshToken string, newRefreshToken string) (*models.User, error) {
	var foundUser models.User
	if result := DB.Model(&models.User{}).Where(&models.User{RefreshToken: oldRefreshToken}).First(&foundUser); result.Error != nil {
		return nil, result.Error
	}
	foundUser.RefreshToken = newRefreshToken
	DB.Save(foundUser)
	return &foundUser, nil
}
