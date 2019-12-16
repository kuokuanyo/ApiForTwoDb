package controllers

import (
	"ApiForTwoDb/driver"
	models "ApiForTwoDb/model"
	"ApiForTwoDb/repository"
	"ApiForTwoDb/utils"
	"encoding/json"
	"net/http"
)

//add value
//mssql
func (c Controller) MssqlAddValue(MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event models.Event
		var error models.Error
		//decode
		json.NewDecoder(r.Body).Decode(&event)

		userRepo := repository.UserRepository{}
		err := userRepo.MssqlInsertValue(MsSqlDb, event)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		utils.SendSuccess(w, event)
	}
}

//get all data
//mssql
func (c Controller) MssqlGetAll(MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var events []models.Event
		var error models.Error

		userRepo := repository.UserRepository{}
		events, err := userRepo.MssqlQueryAllData(MsSqlDb, events)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		utils.SendSuccess(w, events)
	}
}

//get some data
//mssql
func (c Controller) MssqlGetSome(MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var event models.Event
		var error models.Error
		var events []models.Event

		//decode
		json.NewDecoder(r.Body).Decode(&event)

		userRepo := repository.UserRepository{}
		events, err = userRepo.MssqlQuerySomeData(MsSqlDb, events, event)
		if err != nil {
			//找不到資料
			if err.Error() == "record not found" {
				error.Message = "The user does not exist!"
				utils.SendError(w, http.StatusBadRequest, error)
				return
			} else {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}
		utils.SendSuccess(w, events)
	}
}

//update value
//mssql
func (c Controller) MssqlUpdate(MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var error models.Error
		var event models.Event

		//decode
		json.NewDecoder(r.Body).Decode(&event)

		userRepo := repository.UserRepository{}
		err = userRepo.MssqlUpdateData(MsSqlDb, event)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		utils.SendSuccess(w, "Success!")
	}
}

//delete value
//mssql
func (c Controller) MssqlDelete(MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var error models.Error
		var event models.Event

		//decode
		json.NewDecoder(r.Body).Decode(&event)

		userRepo := repository.UserRepository{}
		err = userRepo.MssqlDeleteData(MsSqlDb, event)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		utils.SendSuccess(w, "Success!")
	}
}
