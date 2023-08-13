package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func StravaLoginHandler(w http.ResponseWriter, r *http.Request) {
	STRAVA_CLIENT_ID := os.Getenv("STRAVA_CLIENT_ID")
	STRAVA_REDIRECT_URI := os.Getenv("STRAVA_REDIRECT_URI")
	STRAVA_SCOPE := os.Getenv("STRAVA_SCOPE")

	stravaLoginRedirectUrl := fmt.Sprintf("https://www.strava.com/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s&approval_prompt=force&scope=%s", STRAVA_CLIENT_ID, STRAVA_REDIRECT_URI, STRAVA_SCOPE)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, stravaLoginRedirectUrl)
}
