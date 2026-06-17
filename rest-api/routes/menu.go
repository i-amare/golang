package routes

import (
	"encoding/json"
	"net/http"

	"github.com/i-amare/rest-api/models"
	"github.com/i-amare/rest-api/utils"
)

func createMenuItem(w http.ResponseWriter, r *http.Request) {
	m, err := parseMenuData(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error parsing data", err)
		return
	}

	if err = m.Save(); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error saving menu item", err)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "Menu item created", m)
}

func parseMenuData(r *http.Request) (models.MenuItem, error) {
	var m models.MenuItem
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		return m, err
	}
	return m, nil
}
