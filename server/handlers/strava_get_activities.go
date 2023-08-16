package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zerefwayne/rooots/server/dto/strava"
	"github.com/zerefwayne/rooots/server/middleware"
	"github.com/zerefwayne/rooots/server/utils"
)

func GetActivities(w http.ResponseWriter, r *http.Request) {
	authData := r.Context().Value(middleware.AuthorizationContextKey{}).(*middleware.AuthorizationData)

	stravaRequestUri := "https://www.strava.com/api/v3/athlete/activities?per_page=100"
	request, err := http.NewRequest(http.MethodGet, stravaRequestUri, nil)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	request.Header.Add("Authorization", authData.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil || resp.StatusCode != http.StatusOK {
		utils.HandleHttpError(err, w)
		return
	}
	defer resp.Body.Close()

	responseObject := make(strava.SummaryActivityList, 0)

	err = utils.DecodeJson(resp.Body, &responseObject)
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
