package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/zerefwayne/rooots/server/config"
	"github.com/zerefwayne/rooots/server/dto/strava"
	"github.com/zerefwayne/rooots/server/middleware"
	"github.com/zerefwayne/rooots/server/repository"
	"github.com/zerefwayne/rooots/server/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	authData := r.Context().Value(middleware.AuthorizationContextKey{}).(*middleware.AuthorizationData)
	if authData.AccessToken == "" {
		utils.HandleHttpError(fmt.Errorf("unauthorized"), w)
		return
	}

	user, err := repository.FindUserById(config.DB, authData.UserId)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	stravaId := user.StravaId

	stravaRequestUri := fmt.Sprintf("https://www.strava.com/api/v3/athletes/%d/stats", stravaId)

	request, err := http.NewRequest(http.MethodGet, stravaRequestUri, nil)
	request.Header.Add("Authorization", authData.AccessToken)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}
	bodyString := string(bodyBytes)

	var responseObject strava.ActivityStats

	err = json.Unmarshal([]byte(bodyString), &responseObject)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	jsonResponse, err := json.Marshal(&responseObject)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
