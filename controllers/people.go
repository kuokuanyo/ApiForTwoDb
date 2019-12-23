package controllers

import (
	"ApiForTwoDb/driver"
	"ApiForTwoDb/repository"
	"ApiForTwoDb/utils"
	"strings"

	"encoding/json"
	"net/http"

	models "ApiForTwoDb/model"

	"github.com/gorilla/mux"
)

//@Summary add value to database
//@Tags Data
//@Description 插入數值至資料庫
//@Accept json
//@Produce json
//@Param sql path string true "資料庫引擎"
//@Param information body models.People true "add data"
//@Param information body models.Event true "add data"
//@Success 200 {string}} string "Successfully"
//@Failure 500 {object} models.Error "Internal Server Error"
//@Router /v1/mysql/addvalue [post]
func (c Controller) AddValue(MySqlDb *driver.MySqlDb, MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var (
			people   models.People
			event    models.Event
			error    models.Error
			userRepo repository.UserRepository
		)

		//印出url參數
		params := mux.Vars(r)

		switch strings.ToLower(params["sql"]) {
		case "mysql":
			//decode
			json.NewDecoder(r.Body).Decode(&people)
			if err := userRepo.MysqlInsertValue(MySqlDb, people); err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			utils.SendSuccess(w, people)
		case "mssql":
			//decode
			json.NewDecoder(r.Body).Decode(&event)
			if err := userRepo.MssqlInsertValue(MsSqlDb, event); err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			utils.SendSuccess(w, event)
		}
	}
}

//@Summary get some data from events
//@Tags People
//@Description 從peoples取得部分資料
//@Accept json
//@Produce json
//@Param key1 path int true "Key1"
//@Param key2 path int true "Key2"
//@Param key3 path int true "Key3"
//@Success 200 {object} models.People "data"
//@Failure 400 {object} models.Error "The user does not exist!"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /v1/mysql/getsome/{key1}/{key2}/{key3} [get]
func (c Controller) MysqlGetSome(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			error    models.Error
			people   models.People
			peoples  []models.People
			userRepo repository.UserRepository
		)

		//return map
		//func Vars(r *http.Request) map[string]string
		params := mux.Vars(r)
		people.Key1 = params["key1"]
		people.Key2 = params["key2"]
		people.Key3 = params["key3"]

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
//@Param information body models.mysqlupdate true "update data"
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
//@Param key1 path int true "Key1"
//@Param key2 path int true "Key2"
//@Param key3 path int true "Key3"
//@Success 200 {string} string "Successful delete!"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /v1/mysql/delete/{key1}/{key2}/{key3} [delete]
func (c Controller) MysqlDelete(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			error    models.Error
			people   models.People
			userRepo repository.UserRepository
		)

		//return map
		//func Vars(r *http.Request) map[string]string
		params := mux.Vars(r)
		people.Key1 = params["key1"]
		people.Key2 = params["key2"]
		people.Key3 = params["key3"]

		if err := userRepo.MysqlDeleteData(MySqlDb, people); err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		utils.SendSuccess(w, "Success!")
	}
}
