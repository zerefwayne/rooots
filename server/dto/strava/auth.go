package strava

type ExchangeTokenResponseBody struct {
	TokenType    string         `json:"token_type"`
	ExpiresAt    int64          `json:"expires_at"`
	ExpiresIn    int64          `json:"expires_in"`
	RefreshToken string         `json:"refresh_token"`
	AccessToken  string         `json:"access_token"`
	Athlete      SummaryAthlete `json:"athlete"`
}

type ExchangeTokenBody struct {
	Code string
}

type LoginSuccessResponse struct {
	AccessToken string `json:"accessToken"`
	Name        string `json:"name"`
}
