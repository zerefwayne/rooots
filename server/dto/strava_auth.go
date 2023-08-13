package dto

type ExchangeTokenBody struct {
	Code string
}

type LoginSuccessResponse struct {
	AccessToken string `json:"accessToken"`
	Name        string `json:"name"`
}
