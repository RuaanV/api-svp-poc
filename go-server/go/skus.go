package blah

import (
	"fmt"
	"net/http"
)

type Sku struct {
	//SKU ID for JSON response
	skuID               int64
	skuIDOdbms          int32
	productStyleID      string
	productStyleIDOdbms int32
	productStyleName    string
	variantTypeID       string
	variantTypeName     string
	skuDescription      string
	size                string
}

//Skus Need to discuss how we declare an array/slices of the SKU type as a collection
type Skus struct {
}

//GetSkuByID retieves a single sku by its ID
func GetSkuByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Getting SKU "+r.URL.String())
	//m := Message{"Alice", "Hello", 1294706395881547000}
	//b, err := json.Marshal(m)
}

func ListSkus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

type Message struct {
	Name string
	Body string
	Time int64
}
