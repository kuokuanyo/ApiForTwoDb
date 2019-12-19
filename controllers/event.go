package controllers

import (
	"ApiForTwoDb/driver"
	"ApiForTwoDb/repository"
	"ApiForTwoDb/utils"

	"encoding/json"
	"net/http"

	models "ApiForTwoDb/model"

	"github.com/gorilla/mux"
)

//@Summary add value to events
//@Tags Event
//@Description 加入數值至events
//@Accept json
//@Produce json
//@Param information body models.Event true "add data"
//@Success 200 {object} models.Event "add value"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /v1/mssql/addvalue [post]
func (c Controller) MssqlAddValue(MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			event    models.Event
			error    models.Error
			userRepo repository.UserRepository
		)

		//decode
		json.NewDecoder(r.Body).Decode(&event)

		//插入數值
		if err := userRepo.MssqlInsertValue(MsSqlDb, event); err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		//成功返回
		utils.SendSuccess(w, event)
	}
}

//@Summary get all data from events
//@Tags Event
//@Description 從events取得所有資料
//@Accept json
//@Produce json
//@Success 200 {object} models.Event "get all data"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /v1/mssql/getall [get]
func (c Controller) MssqlGetAll(MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			events   []models.Event
			error    models.Error
			userRepo repository.UserRepository
		)

		//取得所有資料
		events, err := userRepo.MssqlQueryAllData(MsSqlDb, events)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		//成功返回
		utils.SendSuccess(w, events)
	}
}

//@Summary get some data from events
//@Tags Event
//@Description 從events取得部分資料
//@Accept json
//@Produce json
//@Param key1 path int true "Key1"
//@Param key2 path int true "Key2"
//@Param key3 path int true "Key3"
//@Success 200 {object} models.Event "data"
//@Failure 400 {object} models.Error "The user does not exist!"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /v1/mssql/getsome/{key1}/{key2}/{key3} [get]
func (c Controller) MssqlGetSome(MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			event    models.Event
			error    models.Error
			events   []models.Event
			userRepo repository.UserRepository
		)

		//return map
		//func Vars(r *http.Request) map[string]string
		params := mux.Vars(r)
		event.Key1 = params["key1"]
		event.Key2 = params["key2"]
		event.Key3 = params["key3"]

		events, err := userRepo.MssqlQuerySomeData(MsSqlDb, events, event)
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

//@Summary update value
//@Tags Event
//@Description 更新events資料
//@Accept json
//@Produce json
//@Param information body models.mssqlupdate true "update data"
//@Success 200 {string} string "Successful update!"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /v1/mssql/update [put]
func (c Controller) MssqlUpdate(MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			error    models.Error
			event    models.Event
			userRepo repository.UserRepository
		)

		//decode
		json.NewDecoder(r.Body).Decode(&event)

		if err := userRepo.MssqlUpdateData(MsSqlDb, event); err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		utils.SendSuccess(w, "Success!")
	}
}

//@Summary delete value
//@Tags Event
//@Description 刪除events資料
//@Accept json
//@Produce json
//@Param key1 path int true "Key1"
//@Param key2 path int true "Key2"
//@Param key3 path int true "Key3"
//@Success 200 {string} string "Successful delete!"
//@Failure 500 {object} models.Error "Serve(database) error"
//@Router /v1/mssql/delete/{key1}/{key2}/{key3} [delete]
func (c Controller) MssqlDelete(MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			error    models.Error
			event    models.Event
			userRepo repository.UserRepository
		)

		//return map
		//func Vars(r *http.Request) map[string]string
		params := mux.Vars(r)
		event.Key1 = params["key1"]
		event.Key2 = params["key2"]
		event.Key3 = params["key3"]

		if err := userRepo.MssqlDeleteData(MsSqlDb, event); err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		utils.SendSuccess(w, "Success!")
	}
}
