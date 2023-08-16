package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zerefwayne/rooots/server/dto/strava"
	"github.com/zerefwayne/rooots/server/middleware"
	"github.com/zerefwayne/rooots/server/utils"
)

func GetActivity(w http.ResponseWriter, r *http.Request) {
	authData := r.Context().Value(middleware.AuthorizationContextKey{}).(*middleware.AuthorizationData)

	activityId := mux.Vars(r)["activityId"]

	stravaRequestUri := fmt.Sprintf("https://www.strava.com/api/v3/activities/%s}", activityId)
	request, err := http.NewRequest(http.MethodGet, stravaRequestUri, nil)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	request.Header.Add("Authorization", authData.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil || resp.StatusCode != http.StatusOK {
		utils.HandleHttpError(err, w, resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	var responseObject strava.SummaryActivity

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
