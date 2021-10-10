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
		json.NewEncoder(writer).Encode(models.UserList.UserArrayToPublic(users))
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
	requestBody, _ := ioutil.ReadAll(req.Body)
	writer.Header().Set("Content-Type", "application/json")
	var user models.User
	var errorResponse helper.ErrorResponse
	json.Unmarshal(requestBody, &user)
	validateStatus := user.ValidateFor("create")
	if validateStatus != nil {
		errorResponse.Error = "Bad request"
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(errorResponse)
	} else {
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
			json.NewEncoder(writer).Encode(user.ToPublic())
		}
	}

}

func UpdateUser(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	var errorResponse helper.ErrorResponse
	var user models.User
	ormResult := database.DatabaseConnector.First(&user, id)
	writer.Header().Set("Content-Type", "application/json")
	if ormResult.Error == nil {
		//We don't use User Entity because update service have optional field
		m := make(map[string]string)
		requestBody, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(requestBody, &m)
		if m["password"] != "" {
			password, _ := bcrypt.GenerateFromPassword([]byte(m["password"]), 10)
			user.Password = string(password)
		}
		if m["name"] != "" {
			user.Name = m["name"]
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
	id := mux.Vars(req)["id"]
	var errorResponse helper.ErrorResponse
	var user models.User
	writer.Header().Set("Content-Type", "application/json")
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
