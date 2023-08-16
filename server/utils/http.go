package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func HandleHttpError(err error, w http.ResponseWriter, args ...int) {
	defaultStatusCode := http.StatusInternalServerError

	if len(args) > 0 {
		defaultStatusCode = args[0]
	}

	fmt.Println("Error: ", err.Error())
	http.Error(w, err.Error(), defaultStatusCode)
}

func GetCookie(name string, value string, expires time.Time) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expires,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
}

func RemoveCookie(cookie *http.Cookie) *http.Cookie {
	return &http.Cookie{
		Name:     cookie.Name,
		Value:    cookie.Name,
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
}

func DecodeJson(source io.ReadCloser, v interface{}) error {
	return json.NewDecoder(source).Decode(v)
}

func RespondWithJson(w http.ResponseWriter, v interface{}, statusCode int) error {
	jsonResponse, err := json.Marshal(&v)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
	return nil
}
