package controller

import (
	"encoding/json"
	"go-chi-rest-api-boilerplate/pkg/types"
	"go-chi-rest-api-boilerplate/pkg/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	utils.Mu.Lock()
	defer utils.Mu.Unlock()

	var req types.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseGenerator(w, 403, true, "invalidRequestPayload", nil, "Some information missing while validating your request.")
		return
	}

	if req.FirstName == "" {
		utils.ResponseGenerator(w, 403, true, "inputParamsValidationFailed", nil, "Some information missing while validating your request.")
		return
	}

	userID := utils.GenerateNanoIdWithLength(15)
	fullName := req.FirstName + " " + req.LastName
	username := utils.RemoveSpacesAndSpecialChars(fullName) + "_" + userID

	if !utils.IsUsernameUnique(userID, username) {
		utils.ResponseGenerator(w, 400, true, "usernameAlreadyExists", nil, "username already exists.")
		return
	}

	user := &types.User{
		UserID:    userID,
		Username:  username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	utils.UsersData = append(utils.UsersData, user)

	utils.ResponseGenerator(w, 200, false, "", map[string]interface{}{
		"userId": userID,
	}, "User created successfully")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	usersData := utils.UsersData
	utils.ResponseGenerator(w, 200, false, "", map[string]interface{}{
		"users": usersData,
	}, "Users fetched successfully")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	var req types.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseGenerator(w, 403, true, "invalidRequestPayload", nil, "Some information missing while validating your request.")
		return
	}

	isUsernameExists := req.Username != ""
	isFirstNameExists := req.FirstName != ""
	isLastNameExists := req.LastName != ""

	if userId == "" || (!isUsernameExists && !isFirstNameExists && !isLastNameExists) {
		utils.ResponseGenerator(w, 403, true, "inputParamsValidationFailed", nil, "Some information missing while validating your request")
		return
	}

	utils.Mu.Lock()
	defer utils.Mu.Unlock()

	if isUsernameExists && !utils.IsUsernameUnique(userId, req.Username) {
		utils.ResponseGenerator(w, 400, true, "usernameAlreadyExists", nil, "username already exists.")
		return
	}

	for i, user := range utils.UsersData {
		if user.UserID == userId {
			_user := utils.UsersData[i]
			if isUsernameExists {
				_user.Username = req.Username
			}
			if isFirstNameExists {
				_user.FirstName = req.FirstName
			}
			if isLastNameExists {
				_user.LastName = req.LastName
			}
			utils.ResponseGenerator(w, 200, false, "", nil, "User updated successfully.")
			return
		}
	}
	utils.ResponseGenerator(w, 404, true, "userNotFound", nil, "User does not exists.")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	if userId == "" {
		utils.ResponseGenerator(w, 403, true, "inputParamsValidationFailed", nil, "Some information missing while validating your request")
		return
	}

	utils.Mu.Lock()
	defer utils.Mu.Unlock()

	for i, user := range utils.UsersData {
		if user.UserID == userId {
			utils.UsersData = append(utils.UsersData[:i], utils.UsersData[i+1:]...)
			utils.ResponseGenerator(w, 200, false, "", nil, "User deleted successfully.")
			return
		}
	}

	utils.ResponseGenerator(w, 404, true, "userNotFound", nil, "User does not exists.")
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	if userId == "" {
		utils.ResponseGenerator(w, 403, true, "inputParamsValidationFailed", nil, "Some information missing while validating your request")
		return
	}

	usersData := utils.UsersData

	for _, user := range usersData {
		if user.UserID == userId {
			utils.ResponseGenerator(w, 200, false, "", map[string]interface{}{
				"user": user,
			}, "User details fetched successfully.")
			return
		}
	}

	utils.ResponseGenerator(w, 404, true, "userNotFound", nil, "User does not exists.")
}

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	if username == "" {
		utils.ResponseGenerator(w, 403, true, "inputParamsValidationFailed", nil, "Some information missing while validating your request")
		return
	}

	usersData := utils.UsersData

	for _, user := range usersData {
		if user.Username == username {
			utils.ResponseGenerator(w, 200, false, "", map[string]interface{}{
				"user": user,
			}, "User details fetched successfully.")
			return
		}
	}

	utils.ResponseGenerator(w, 404, true, "userNotFound", nil, "User does not exists.")
}
