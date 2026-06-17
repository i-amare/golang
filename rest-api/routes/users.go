package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/i-amare/rest-api/models"
	"github.com/i-amare/rest-api/utils"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	u, err := parseUserData(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error parsing data", err)
		return
	}

	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error parsing password", err)
		return
	}

	if err = u.Save(); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error saving user", err)
		return
	}

	authToken, err := utils.GenerateAuthToken(u.Email, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "User created, error generating auth token", err)
		return
	}

	utils.WriteSuccess(w, http.StatusCreated, "User created", map[string]interface{}{
		"user":      u,
		"authToken": authToken,
	})
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	usersArr, err := models.GetAllUsers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error fetching users", err)
		return
	}

	fmt.Println("Printing")
	utils.WriteSuccess(w, http.StatusOK, "Users fetched successfully", usersArr)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	u, err := parseUserData(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error parsing data", err)
		return
	}

	u.ID, err = u.GetUserID()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error fetching user", err)
		return
	}

	isValid, err := u.ValidateCredentials()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error validating credentials", err)
		return
	}

	if !isValid {
		utils.WriteError(w, http.StatusBadRequest, "Password is invalid", u)
		return
	}

	authToken, err := utils.GenerateAuthToken(u.Email, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error generating auth token", err)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "Password is valid", map[string]interface{}{
		"user":      u,
		"authToken": authToken,
	})
}

func parseUserData(r *http.Request) (models.User, error) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}
