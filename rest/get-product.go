package rest

import (
	"encoding/json"
	"fmt"
	"github.com/Silverhammertech/mit-location-svc/mongodb"
	"net/http"
)

func HandleGetAllProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetAllProducts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting products. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	jsonOutput, err := json.Marshal(products)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling products. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	// Always set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonOutput))
}

func HandleGetHybridProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetHybridProducts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting products. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	jsonOutput, err := json.Marshal(products)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling products. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	// Always set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonOutput))
}

func HandleGetIndicaProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetIndicaProducts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting products. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	jsonOutput, err := json.Marshal(products)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling products. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	// Always set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonOutput))
}

func HandleGetSativaProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetSativaProducts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting products. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	jsonOutput, err := json.Marshal(products)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling products. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	// Always set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonOutput))
}
