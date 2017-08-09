package 

import (
	"net/http"
)

type Default struct {

}

//CreateSkuCommand created SKU via API putting it onto a queue
func CreateSkuCommand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
}

func PingGet(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

