package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zerefwayne/rooots/server/config"
	"github.com/zerefwayne/rooots/server/dto/strava"
	"github.com/zerefwayne/rooots/server/repository"
	"github.com/zerefwayne/rooots/server/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("Authorization")
	vars := mux.Vars(r)

	userId := vars["id"]
	log.Println("VARS[userId]", userId)

	userIdUuid, err := uuid.Parse(userId)
	log.Println("userUuid", userIdUuid)

	if err != nil || accessToken == "" || userId == "" {
		log.Println("Unauthorized!")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("Accessing strava data with accessToken", accessToken)

	user, err := repository.FindUserById(config.DB, userIdUuid)
	if err != nil {
		utils.HandleHttpError(err, w)
		return
	}

	stravaId := user.StravaId

	stravaRequestUri := fmt.Sprintf("https://www.strava.com/api/v3/athletes/%d/stats", stravaId)

	request, err := http.NewRequest(http.MethodGet, stravaRequestUri, nil)
	request.Header.Add("Authorization", accessToken)
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
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)

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
