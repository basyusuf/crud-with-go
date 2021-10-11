package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"main/app"
	"main/database"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgconn"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetAllUser(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var users []models.User
	var errorResponse app.ErrorResponse
	ormResult := database.DatabaseConnector.Find(&users)
	if ormResult.Error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		errorResponse.Error = "server error"
		json.NewEncoder(writer).Encode(errorResponse)
	} else if ormResult.RowsAffected == 0 {
		writer.WriteHeader(http.StatusNotFound)
		errorResponse.Error = "User with that id does not exist"
		json.NewEncoder(writer).Encode(errorResponse)
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(models.UserList.UserArrayToPublic(users))
	}
}

func GetUserById(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id := mux.Vars(req)["id"]
	var errorResponse app.ErrorResponse
	var user models.User
	//check if there is such a user in the database
	ormResult := database.DatabaseConnector.First(&user, id)
	if ormResult.RowsAffected != 0 && ormResult.Error == nil {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(user.ToPublic())
	} else {
		if errors.Is(ormResult.Error, gorm.ErrRecordNotFound) {
			writer.WriteHeader(http.StatusNotFound)
			errorResponse.Error = "User with that id does not exist"
			json.NewEncoder(writer).Encode(errorResponse)
		} else {
			writer.WriteHeader(http.StatusBadRequest)
			errorResponse.Error = "Bad request"
			json.NewEncoder(writer).Encode(errorResponse)
		}
	}
}

func CreateUser(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	requestBody, _ := ioutil.ReadAll(req.Body)
	var user models.User
	var errorResponse app.ErrorResponse
	json.Unmarshal(requestBody, &user)
	//Input validation for create
	validateStatus := user.ValidateFor(models.ValidationStatus.CREATE)
	if validateStatus != nil {
		errorResponse.Error = "Bad request"
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}
	//Check mail on DB
	checkEmailOnDB := database.DatabaseConnector.First(&user, "email = ?", user.Email)
	if checkEmailOnDB.RowsAffected != 0 {
		errorResponse.Error = "User with that email already exists"
		writer.WriteHeader(http.StatusForbidden)
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}
	createError := database.DatabaseConnector.Create(&user).Error
	if createError != nil {
		var perr *pgconn.PgError
		errors.As(createError, &perr)
		if perr.Code == "23505" {
			errorResponse.Error = "User with that email already exists"
			writer.WriteHeader(http.StatusForbidden)
		} else {
			errorResponse.Error = "Bad request"
			writer.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(writer).Encode(errorResponse)
	} else {
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(user.ToPublic())
	}
}

func UpdateUser(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var errorResponse app.ErrorResponse
	id := mux.Vars(req)["id"]
	var user models.User
	//check if there is such a user in the database
	ormResult := database.DatabaseConnector.First(&user, id)
	if ormResult.Error == nil {
		var bodyUser models.User
		requestBody, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(requestBody, &bodyUser)
		//check the request body have a any field for update
		updateValidationError := bodyUser.ValidateFor(models.ValidationStatus.UPDATE)
		if updateValidationError != nil {
			writer.WriteHeader(http.StatusBadRequest)
			errorResponse.Error = "Bad request"
			json.NewEncoder(writer).Encode(errorResponse)
			return
		}
		if bodyUser.Password != "" {
			password, _ := bcrypt.GenerateFromPassword([]byte(bodyUser.Password), 10)
			user.Password = string(password)
		}
		if bodyUser.Name != "" {
			user.Name = bodyUser.Name
		}
		updateResult := database.DatabaseConnector.Save(&user)
		if updateResult.Error != nil {
			writer.WriteHeader(http.StatusBadRequest)
			errorResponse.Error = "Bad request"
			json.NewEncoder(writer).Encode(errorResponse)
		} else {
			writer.WriteHeader(http.StatusCreated)
			json.NewEncoder(writer).Encode(user.ToPublic())
		}
	} else {
		if errors.Is(ormResult.Error, gorm.ErrRecordNotFound) {
			writer.WriteHeader(http.StatusNotFound)
			errorResponse.Error = "User with that id does not exist"
			json.NewEncoder(writer).Encode(errorResponse)
		} else {
			writer.WriteHeader(http.StatusBadRequest)
			errorResponse.Error = "Bad request"
			json.NewEncoder(writer).Encode(errorResponse)
		}
	}

}

func DeleteUser(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id := mux.Vars(req)["id"]
	var errorResponse app.ErrorResponse
	var user models.User
	//check if there is such a user in the database
	ormResult := database.DatabaseConnector.First(&user, id)
	if ormResult.Error == nil {
		deleteResult := database.DatabaseConnector.Delete(&user, id)
		if deleteResult.Error == nil && deleteResult.RowsAffected == 1 {
			writer.WriteHeader(http.StatusOK)
		} else {
			writer.WriteHeader(http.StatusBadRequest)
			errorResponse.Error = "Bad request"
			json.NewEncoder(writer).Encode(errorResponse)
		}
	} else {
		if errors.Is(ormResult.Error, gorm.ErrRecordNotFound) {
			writer.WriteHeader(http.StatusNotFound)
			errorResponse.Error = "User with that id does not exist"
			json.NewEncoder(writer).Encode(errorResponse)
		} else {
			writer.WriteHeader(http.StatusBadRequest)
			errorResponse.Error = "Bad request"
			json.NewEncoder(writer).Encode(errorResponse)
		}
	}

}
