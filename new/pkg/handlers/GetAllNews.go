package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"main/new/pkg/models"
)

func (h handler) GetAllNews(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1 // Default to page 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10 // Default to 10 items per page
	}

	offset := (page - 1) * limit

	var news []models.New

	if result := h.DB.Limit(limit).Offset(offset).Find(&news); result.Error != nil {
		fmt.Println(result.Error)
	}

	type result struct {
		List  []models.New `json:"list"`
		Count int64        `json:"count"`
	}

	var counts int64
	h.DB.Model(&news).Count(&counts)

	fmt.Println(counts)

	var result2 result
	result2.Count = counts
	result2.List = news

	fmt.Println(result2)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result2)
}
