package blah

import (
	"fmt"
	"net/http"
)

type Default struct {
}

func CreateSkuCommand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PingGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("X-Apiping", "Alive")
	fmt.Fprintf(w, "Service Alive")
	w.WriteHeader(http.StatusOK)
}
