package handlers

import (
	"encoding/json"
	"fmt"
	"main/new/pkg/models"
	"net/http"
)

func (h handler) GetContacts(w http.ResponseWriter, r *http.Request) {
	var contacts []models.Contact

	if result := h.DB.Find(&contacts); result.Error != nil {
		fmt.Println(result.Error)
	}

	fmt.Println(contacts)
	contactIds := []int{}

	for i := len(contacts) - 1; i >= 0; i-- {
		if !contains(contactIds, contacts[i].Id) {
			contactIds = append(contactIds, contacts[i].Id)
		}
	}

	var phones []models.Phone

	h.DB.Where("addressid IN ?", contactIds).Find(&phones)

	var myMap map[int][]string

	// Инициализируем map
	myMap = make(map[int][]string)

	for i := len(phones) - 1; i >= 0; i-- {
		myMap[phones[i].Addressid] = append(myMap[phones[i].Addressid], phones[i].Value)
	}

	fmt.Println(myMap)

	type result struct {
		Id      int      `json:"id"`
		Address string   `json:"address"`
		Phone   []string `json:"phone"`
	}

	results := []result{}

	for i := len(contacts) - 1; i >= 0; i-- {
		var contactItem result
		contactItem.Id = contacts[i].Id
		contactItem.Address = contacts[i].Address
		contactItem.Phone = myMap[contacts[i].Id]
		results = append(results, contactItem)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
