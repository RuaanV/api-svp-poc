package blah

import (
	"fmt"
	"net/http"
)

type Sku struct {
	//SKU ID for JSON response
	skuID               int64
	skuIDOdbms          int32
	productStyleId      string
	productStyleIdOdbms int32
	productStyleName    string
	variantTypeId       string
	variantTypeName     string
	skuDescription      string
	size                string
}

type Skus struct {
}

func GetSkuById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Getting SKU "+r.URL.String())
}

func ListSkus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
