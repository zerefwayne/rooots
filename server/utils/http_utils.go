package utils

import (
	"fmt"
	"net/http"
)

func HandleHttpError(err error, w http.ResponseWriter) {
	fmt.Println("Error: ", err.Error())
	http.Error(w, err.Error(), http.StatusBadRequest)
}
