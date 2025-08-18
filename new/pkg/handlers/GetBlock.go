package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"main/new/pkg/models"
	"net/http"
)

func (h handler) GetBlocks(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["code"]

	var blocks []models.Block

	//if result := h.DB.Find(&blocks); result.Error != nil {
	//	fmt.Println(result.Error)
	//}

	h.DB.Where("code = ? AND is_active = ?", code, true).Order("id asc").Find(&blocks)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blocks)
}
