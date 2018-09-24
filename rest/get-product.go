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

func HandleGetOtherProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetOtherProducts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting other products. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	jsonOutput, err := json.Marshal(products)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling other products. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	// Always set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonOutput))
}

func HandleGetOtherProductEdible(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetEdibleProducts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting edible products. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	jsonOutput, err := json.Marshal(products)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling edible products. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	// Always set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonOutput))
}

func HandleGetOtherProductConcentrate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetEdibleProducts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting concentrate products. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	jsonOutput, err := json.Marshal(products)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling concentrate products. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	// Always set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonOutput))
}

func HandleGetOtherProductConcentrateCartridge(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetConcentrateCartridgeProducts()
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

func HandleGetOtherProductConcentrateIngestible(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetConcentrateIngestibleProducts()
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

func HandleGetOtherProductConcentrateSolvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetConcentrateSolventProducts()
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

func HandleGetOtherProductConcentrateSolventless(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetConcentrateSolventlessProducts()
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

func HandleGetOtherProductConcentrateTerpene(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetConcentrateTerpeneProducts()
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

func HandleGetOtherProductEdibleBeverage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetEdibleBeverageProducts()
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

func HandleGetOtherProductEdibleBreakfast(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetEdibleBreakFastProducts()
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

func HandleGetOtherProductEdibleBrownie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetEdibleBrownieProducts()
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

func HandleGetOtherProductEdibleCandy(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetEdibleCandyProducts()
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

func HandleGetOtherProductEdibleCapsule(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetEdibleCapsuleProducts()
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

func HandleGetOtherProductEdibleChocolate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetEdibleChocolateProducts()
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

func HandleGetOtherProductEdibleCondiment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetEdibleCondimentProducts()
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

func HandleGetOtherProductEdibleCookie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetEdibleCookieProducts()
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

func HandleGetOtherProductEdibleCooking(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetEdibleCookingProducts()
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

func HandleGetOtherProductEdibleFrozen(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetEdibleFrozenProducts()
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

func HandleGetOtherProductEdibleSnack(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetEdibleSnackProducts()
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

func HandleGetOtherProductEdibleTincture(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// query mongodb
	products, err := mongodb.GetEdibleTinctureProducts()
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
