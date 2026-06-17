package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/i-amare/rest-api/middleware"
	"github.com/i-amare/rest-api/models"
	"github.com/i-amare/rest-api/utils"
)

func createVendor(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)

	v, err := parseVendorData(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error parsing data", err)
		return
	}

	v.OwnerID = userID
	if err = v.Save(); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error saving vendor", err)
		return
	}

	utils.WriteSuccess(w, http.StatusCreated, "Vendor created", v)
}

func getVendor(w http.ResponseWriter, r *http.Request) {
	id, err := parseVendorID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error parsing vendor ID", err)
		return
	}

	v, err := fetchVendor(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error finding vendor", err)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "Vendor found", v)
}

func getAllVendors(w http.ResponseWriter, r *http.Request) {
	v, err := models.GetAllVendors()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error fetching vendors", err)
		return
	}
	utils.WriteSuccess(w, http.StatusOK, "Vendors fetched successfully", v)
}

func updateVendor(w http.ResponseWriter, r *http.Request) {
	id, err := parseVendorID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error parsing vendor ID", err)
		return
	}

	v, err := fetchVendor(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error finding vendor", err)
		return
	}

	userID := middleware.GetUserID(r)
	if v.OwnerID != userID && userID != utils.ADMIN_ID {
		utils.WriteError(w, http.StatusUnauthorized, "Unauthorised to edit vendor", nil)
		return
	}

	inputVendor, err := parseVendorData(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error parsing data", err)
		return
	}

	v.Name = inputVendor.Name
	v.Description = inputVendor.Description
	_, err = v.Update()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error updating vendor", err)
		return
	}

	utils.WriteSuccess(w, http.StatusAccepted, "Successfully updated vendor", v)
}

func deleteVendor(w http.ResponseWriter, r *http.Request) {
	id, err := parseVendorID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error parsing vendor ID", err)
		return
	}

	v, err := fetchVendor(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error finding vendor", err)
		return
	}

	userID := middleware.GetUserID(r)
	if v.OwnerID != userID && userID != utils.ADMIN_ID {
		utils.WriteError(w, http.StatusUnauthorized, "Unauthorised to delete vendor", nil)
		return
	}

	if err = models.DeleteVendor(id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error deleting database entry", err)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "Vendor deleted", nil)
}

func parseVendorID(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func fetchVendor(id int64) (models.Vendor, error) {
	return models.GetVendor(id)
}

func parseVendorData(r *http.Request) (models.Vendor, error) {
	var v models.Vendor
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return v, err
	}
	return v, nil
}
