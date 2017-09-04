package blah

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//Sku is the base struct for SKU Api definition
type Sku struct {
	//SKU ID for JSON response
	SkuID               int64 `json:"sku_id"`
	SkuIDOdbms          int32
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
	lines []Sku
}

//GetSkuByID retieves a single sku by its ID
func GetSkuByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/html; charset=UTF-8") // switch this to application/json when read
	w.WriteHeader(http.StatusAccepted)
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Yo yo, it works - SKU Id: %s - %s", vars["skuId"], time.Now()) // this must respond with matching content-type.

}

func ListSkus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

//Message demo struct to explore data types
type Message struct {
	Name string
	Body string
	Time int64
}

//SKU data collection of skus
type Data struct {
	skus []Sku
}

// Implement json.Unmarshaller
func (d *Sku) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &d.SkuID)
}

type Post struct {
	ID int64 `json:"id"`
}
