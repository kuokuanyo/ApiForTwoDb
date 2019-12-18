package controllers

import (
	"ApiForTwoDb/driver"
	"ApiForTwoDb/repository"
	"ApiForTwoDb/utils"

	"encoding/json"
	"net/http"

	models "ApiForTwoDb/model"
)

//@Summary add value to peoples
//@Tags People
//@Description 插入數值至peoples
//@Accept json
//@Produce json
//@Param information body model.People true "add data"
//@Success 200 {object} models.People "add data"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /v1/mysql/addvalue [post]
func (c Controller) MysqlAddValue(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			people   models.People
			error    models.Error
			userRepo repository.UserRepository
		)

		//decode
		json.NewDecoder(r.Body).Decode(&people)

		if err := userRepo.MysqlInsertValue(MySqlDb, people); err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		utils.SendSuccess(w, people)
	}
}

//@Summary get all data from peoples
//@Tags People
//@Description 從peoples取得所有資料
//@Accept json
//@Produce json
//@Success 200 {object} models.People "get all data"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /v1/mysql/getall [get]
func (c Controller) MysqlGetAll(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			peoples  []models.People
			error    models.Error
			userRepo repository.UserRepository
		)

		peoples, err := userRepo.MysqlQueryAllData(MySqlDb, peoples)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		utils.SendSuccess(w, peoples)
	}
}

//@Summary get some data from events
//@Tags People
//@Description 從peoples取得部分資料
//@Accept json
//@Produce json
//@Param information body model.mysqlgetsome true "get some data from condition"
//@Success 200 {object} models.People "data"
//@Failure 400 {object} models.Error "The user does not exist!"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /v1/mysql/getsome [get]
func (c Controller) MysqlGetSome(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			error    models.Error
			people   models.People
			peoples  []models.People
			userRepo repository.UserRepository
		)

		//decode
		json.NewDecoder(r.Body).Decode(&people)

		peoples, err := userRepo.MysqlQuerySomeData(MySqlDb, peoples, people)
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

//@Summary update value
//@Tags People
//@Description 更新peoples數值
//@Accept json
//@Produce json
//@Param information body model.mysqlupdate true "update data"
//@Success 200 {string} string "Successful update!"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /v1/mysql/update [put]
func (c Controller) MysqlUpdate(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			error    models.Error
			people   models.People
			userRepo repository.UserRepository
		)

		//decode
		json.NewDecoder(r.Body).Decode(&people)

		if err := userRepo.MysqlUpdateData(MySqlDb, people); err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		utils.SendSuccess(w, "Success!")
	}
}

//@Summary delete value
//@Tags People
//@Description 刪除peoples數值
//@Accept json
//@Produce json
//@Param information body model.mysqldelete true "delete data"
//@Success 200 {string} string "Successful delete!"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /v1/mysql/delete [delete]
func (c Controller) MysqlDelete(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			error    models.Error
			people   models.People
			userRepo repository.UserRepository
		)

		//decode
		json.NewDecoder(r.Body).Decode(&people)

		if err := userRepo.MysqlDeleteData(MySqlDb, people); err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		utils.SendSuccess(w, "Success!")
	}
}
