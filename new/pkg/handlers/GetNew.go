package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main/new/pkg/models"
)

func (h handler) GetNew(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find book by Id
	var new models.New

	if result := h.DB.Where("id = ? AND is_active = ?", id, true).Limit(1).Find(&new); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	if new.Id > 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(new)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}
