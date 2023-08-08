package models

type ExchangeTokenResponseBody struct {
	TokenType    string         `json:"token_type"`
	ExpiresAt    uint64         `json:"expires_at"`
	ExpiresIn    uint64         `json:"expires_in"`
	RefreshToken string         `json:"refresh_token"`
	AccessToken  string         `json:"access_token"`
	Athlete      SummaryAthlete `json:"athlete"`
}
