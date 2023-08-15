package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/zerefwayne/rooots/server/config"
	"github.com/zerefwayne/rooots/server/middleware"
	"github.com/zerefwayne/rooots/server/models"
	"github.com/zerefwayne/rooots/server/repository"
	"github.com/zerefwayne/rooots/server/utils"
)

func cleanSensitiveInformation(user *models.User) {
	user.RefreshToken = ""
	user.Id = uuid.Nil
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	authData := r.Context().Value(middleware.AuthorizationContextKey{}).(*middleware.AuthorizationData)

	user, err := repository.FindUserById(config.DB, authData.UserId)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	cleanSensitiveInformation(user)

	jsonResponse, err := json.Marshal(&user)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
