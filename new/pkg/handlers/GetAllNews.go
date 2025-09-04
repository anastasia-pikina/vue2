package handlers

import (
	"encoding/json"
	"fmt"
	"main/new/pkg/models"
	"net/http"
	"strconv"
)

func (h handler) GetAllNews(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil || page < 1 {
		page = 1 // Default to page 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	var news []models.New

	if err != nil || limit < 1 {
		if result := h.DB.Order("id desc").Where("is_active = ?", true).Find(&news); result.Error != nil {
			fmt.Println(result.Error)
		}
	} else {
		offset := (page - 1) * limit

		if result := h.DB.Order("id desc").Where("is_active = ?", true).Limit(limit).Offset(offset).Find(&news); result.Error != nil {
			fmt.Println(result.Error)
		}
	}

	type resultModel struct {
		List  []models.New `json:"list"`
		Count int64        `json:"count"`
	}

	var counts int64
	h.DB.Model(&news).Count(&counts)

	var result resultModel
	result.Count = counts
	result.List = news

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
