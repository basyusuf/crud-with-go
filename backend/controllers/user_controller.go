package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"main/database"
	"main/helper"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgconn"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetAllUser(writer http.ResponseWriter, req *http.Request) {
	var users []models.User
	var errorResponse helper.ErrorResponse
	ormResult := database.DatabaseConnector.Find(&users)
	writer.Header().Set("Content-Type", "application/json")
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
		json.NewEncoder(writer).Encode(users)
	}
}

func GetUserById(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var errorResponse helper.ErrorResponse
	var user models.User
	ormResult := database.DatabaseConnector.First(&user, id)
	writer.Header().Set("Content-Type", "application/json")
	if ormResult.RowsAffected != 0 && ormResult.Error == nil {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(user)
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

func CreateUser(writer http.ResponseWriter, res *http.Request) {
	requestBody, _ := ioutil.ReadAll(res.Body)
	writer.Header().Set("Content-Type", "application/json")
	var user models.User
	var errorResponse helper.ErrorResponse
	json.Unmarshal(requestBody, &user)
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(password)
	createError := database.DatabaseConnector.Create(&user).Error
	if createError != nil {
		var perr *pgconn.PgError
		errors.As(createError, &perr)
		if perr.Code == "23505" {
			errorResponse.Error = "User with that email already exists"
			writer.WriteHeader(403)
		} else {
			errorResponse.Error = "Bad request"
			writer.WriteHeader(400)
		}
		json.NewEncoder(writer).Encode(errorResponse)
	} else {
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(user)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
