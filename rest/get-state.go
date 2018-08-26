package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"encoding/json"
	"github.com/Silverhammertech/mit-location-svc/model"
	"github.com/Silverhammertech/mit-location-svc/mongodb"
	"github.com/gorilla/mux"
)

func HandleGetState(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req model.RequestGeo
	var res model.ResponseGeo
	var err error

	vars := mux.Vars(r)
	req.State = vars["state-code"]
	if req.State == "" {
		http.Error(w, "Please supply state code.", http.StatusBadRequest)
		return
	}

	req.State = strings.ToUpper(req.State)

	// set process date
	res.Summary.Date = time.Now().UTC()

	// validate sort param
	sortParam := r.URL.Query().Get("sort")
	if len(sortParam) > 0 {
		switch sp := strings.ToLower(sortParam); sp {
		case strings.ToLower(model.SORT_PARAM_CITY):
			req.SortParam = model.SORT_PARAM_CITY
		case strings.ToLower(model.SORT_PARAM_LICENSETYPE):
			req.SortParam = model.SORT_PARAM_LICENSETYPE
		case strings.ToLower(model.SORT_PARAM_STATE):
			req.SortParam = model.SORT_PARAM_STATE
		case strings.ToLower(model.SORT_PARAM_TYPE):
			req.SortParam = model.SORT_PARAM_TYPE
		default:
			http.Error(w, "sort param is not valid.", http.StatusBadRequest)
			return
		}
	} else {
		req.SortParam = model.SORT_PARAM_NAME
	}

	// validate sort param
	sortOrder := r.URL.Query().Get("order")
	if len(sortOrder) > 0 {
		switch so := strings.ToLower(sortOrder); so {
		case model.SORT_ORDER_ASC, model.SORT_ORDER_DESC:
			req.SortOrder = so
		default:
			http.Error(w, "sort order is not valid.", http.StatusBadRequest)
			return
		}
	} else {
		req.SortOrder = model.SORT_ORDER_ASC
	}

	// validate limit
	limit := r.URL.Query().Get("limit")
	if len(limit) > 0 {
		req.Limit, err = strconv.Atoi(limit)
		if err != nil {
			http.Error(w, "limit is not valid.", http.StatusBadRequest)
			return
		}
	} else {
		req.Limit = 100
	}

	// get params
	res.Summary.Params = req

	// query mongodb
	err = mongodb.GetState(&res)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting locations. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	// get count
	res.Summary.Count = len(res.Locations)

	// get duration
	res.Summary.Duration = time.Now().Sub(res.Summary.Date).String()

	jsonOutput, err := json.Marshal(res)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling locations. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	// Always set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonOutput))
}
