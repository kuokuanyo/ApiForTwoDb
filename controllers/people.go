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
//mysql
func (c Controller) MysqlAddValue(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var people models.People
		var error models.Error
		//decode
		json.NewDecoder(r.Body).Decode(&people)

		userRepo := repository.UserRepository{}
		err := userRepo.MysqlInsertValue(MySqlDb, people)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		utils.SendSuccess(w, people)
	}
}

//get all data
//mysql
func (c Controller) MysqlGetAll(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var peoples []models.People
		var error models.Error

		userRepo := repository.UserRepository{}
		peoples, err := userRepo.MysqlQueryAllData(MySqlDb, peoples)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		utils.SendSuccess(w, peoples)
	}
}

//get some data
//mysql
func (c Controller) MysqlGetSome(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var error models.Error
		var people models.People
		var peoples []models.People

		//decode
		json.NewDecoder(r.Body).Decode(&people)

		userRepo := repository.UserRepository{}
		peoples, err = userRepo.MysqlQuerySomeData(MySqlDb, peoples, people)
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
		utils.SendSuccess(w, peoples)
	}
}

//update value
//mysql
func (c Controller) MysqlUpdate(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var error models.Error
		var people models.People

		//decode
		json.NewDecoder(r.Body).Decode(&people)

		userRepo := repository.UserRepository{}
		err = userRepo.MysqlUpdateData(MySqlDb, people)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		utils.SendSuccess(w, "Success!")
	}
}

//delete value
//mysql
func (c Controller) MysqlDelete(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var error models.Error
		var people models.People

		//decode
		json.NewDecoder(r.Body).Decode(&people)

		userRepo := repository.UserRepository{}
		err = userRepo.MysqlDeleteData(MySqlDb, people)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		utils.SendSuccess(w, "Success!")
	}
}
