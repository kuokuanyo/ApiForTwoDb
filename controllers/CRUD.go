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

//AddValue insert value
//@Summary add value to database
//@Tags Data
//@Description 插入數值至資料庫
//@Accept json
//@Produce json
//@Param sql path string true "資料庫引擎"
//@Param information body models.People false "add data"
//@Param information body models.Event false "add data"
//@Success 200 {string} string "Successfully"
//@Failure 500 {object} models.Error "Internal Server Error"
//@Router /v1/addvalue/{sql} [post]
func (c Controller) AddValue(MySQLDb *driver.MySQLDb, MsSQLDb *driver.MsSQLDb) http.HandlerFunc {
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
			if err := userRepo.MysqlInsertValue(MySQLDb, people); err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			utils.SendSuccess(w, people)
		case "mssql":
			//decode
			json.NewDecoder(r.Body).Decode(&event)
			if err := userRepo.MssqlInsertValue(MsSQLDb, event); err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			utils.SendSuccess(w, event)
		}
	}
}

//GetAll get all
//@Summary get all data from database
//@Tags Data
//@Description 從資料庫取得所有資料
//@Accept json
//@Produce json
//@Param sql path string true "資料庫引擎"
//@Success 200 {object} models.People "Successfully"
//@Success 200 {object} models.Event "Successfully"
//@Failure 500 {object} models.Error "Internal Server Error"
//@Router /v1/getall/{sql} [get]
func (c Controller) GetAll(MySQLDb *driver.MySQLDb, MsSQLDb *driver.MsSQLDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			peoples  []models.People
			events   []models.Event
			error    models.Error
			userRepo repository.UserRepository
		)

		//印出url參數
		params := mux.Vars(r)

		switch strings.ToLower(params["sql"]) {
		case "mysql":
			//取得所有資料
			peoples, err := userRepo.MysqlQueryAllData(MySQLDb, peoples)
			if err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			//成功返回
			utils.SendSuccess(w, peoples)
		case "mssql":
			//取得所有資料
			events, err := userRepo.MssqlQueryAllData(MsSQLDb, events)
			if err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			//成功返回
			utils.SendSuccess(w, events)
		}
	}
}

//GetSome get some
//@Summary get some data from database
//@Tags Data
//@Description 從資料庫取得部分資料
//@Accept json
//@Produce json
//@Param sql path string true "資料庫引擎"
//@Param key1 query string false "Key1"
//@Param key2 query string false "Key2"
//@Param key3 query string false "Key3"
//@Param number query int false "Number"
//@Param gender query int false "Gender"
//@Param birth query int false "Birth"
//@Param injury_degree query string false "Injury_degree"
//@Param injury_position query int false "Injury_position"
//@Param protection query int false "Protection"
//@Param phone query int false "Phone"
//@Param person query string false "Person"
//@Param car query string false "Car"
//@Param action_status query int false "Action_status"
//@Param qualification query int false "Qualification"
//@Param license query int false "License"
//@Param drinking query int false "Drinking"
//@Param hit query int false "Hit"
//@Param city query string false "City"
//@Param position query string false "Position"
//@Param lane query string false "Lane"
//@Param death query string false "Death"
//@Param injured query string false "Injured"
//@Param death_exceed query string false "Death_exceed"
//@Param weather query string false "Weather"
//@Param light query string false "Light"
//@Param time_year query int false "Time_year"
//@Param time_month query string false "Time_month"
//@Param accident_chinese query string false "Accident_chinese"
//@Param anecdote_chinese query string false "Anecdote_chinese"
//@Success 200 {object} models.People "Successfully"
//@Failure 400 {object} models.Error "Bad Request"
//@Failure 500 {object} models.Error "Internal Server Error"
//@Router /v1/getsome/{sql} [get]
func (c Controller) GetSome(MySQLDb *driver.MySQLDb, MsSQLDb *driver.MsSQLDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			conditions = make(map[string]interface{})
			error      models.Error
			events     []models.Event
			peoples    []models.People
			userRepo   repository.UserRepository
		)

		//return map
		//func Vars(r *http.Request) map[string]string
		params := mux.Vars(r)

		//所有欄位
		MysqlAllCol := []string{"key1", "key2", "key3",
			"number", "gender", "birth", "injury_degree", "injury_position",
			"protection", "phone", "person", "car", "action_status",
			"qualification", "license", "drinking", "hit"}
		MssqlAllCol := []string{"key1", "key2", "key3",
			"city", "position", "lane", "death", "injured", "death_exceed",
			"weather", "light", "time_year", "time_month",
			"accident_chinese", "anecdote_chinese"}

		switch strings.ToLower(params["sql"]) {
		case "mysql":
			//處理query參數
			for _, Colname := range MysqlAllCol {
				value := r.URL.Query()[Colname]
				if len(value) > 0 {
					for i := 0; i < len(value); i++ {
						conditions[Colname] = value[i]
					}
				}
			}

			peoples, err := userRepo.MysqlQuerySomeData(MySQLDb, peoples, conditions)
			if err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			if len(peoples) == 0 {
				error.Message = "The user does not exist!"
				utils.SendError(w, http.StatusBadRequest, error)
				return
			}
			utils.SendSuccess(w, peoples)

		case "mssql":
			//處理query參數
			for _, Colname := range MssqlAllCol {
				value := r.URL.Query()[Colname]
				if len(value) > 0 {
					for i := 0; i < len(value); i++ {
						conditions[Colname] = value[i]
					}
				}
			}

			events, err := userRepo.MssqlQuerySomeData(MsSQLDb, events, conditions)
			if err != nil {
				//找不到資料
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			if len(events) == 0 {
				error.Message = "The user does not exist!"
				utils.SendError(w, http.StatusBadRequest, error)
				return
			}
			utils.SendSuccess(w, events)
		}
	}
}

//Update update
//@Summary update value
//@Tags Data
//@Description 更新資料庫數值
//@Accept json
//@Produce json
//@Param sql path string true "資料庫引擎"
//@Param where_key1 query string false "where_Key1"
//@Param where_key2 query string false "where_Key2"
//@Param where_key3 query string false "where_Key3"
//@Param where_number query int false "where_Number"
//@Param where_gender query int false "where_Gender"
//@Param where_birth query int false "where_Birth"
//@Param where_injury_degree query string false "where_Injury_degree"
//@Param where_injury_position query int false "where_Injury_position"
//@Param where_protection query int false "where_Protection"
//@Param where_phone query int false "where_Phone"
//@Param where_person query string false "where_Person"
//@Param where_car query string false "where_Car"
//@Param where_action_status query int false "where_Action_status"
//@Param where_qualification query int false "where_Qualification"
//@Param where_license query int false "where_License"
//@Param where_drinking query int false "where_Drinking"
//@Param where_hit query int false "where_Hit"
//@Param where_city query string false "where_City"
//@Param where_position query string false "where_Position"
//@Param where_lane query string false "where_Lane"
//@Param where_death query string false "where_Death"
//@Param where_injured query string false "where_Injured"
//@Param where_death_exceed query string false "where_Death_exceed"
//@Param where_weather query string false "where_Weather"
//@Param where_light query string false "where_Light"
//@Param where_time_year query int false "where_Time_year"
//@Param where_time_month query string false "where_Time_month"
//@Param where_accident_chinese query string false "where_Accident_chinese"
//@Param where_anecdote_chinese query string false "where_Anecdote_chinese"
//@Param update_key1 query string false "update_Key1"
//@Param update_key2 query string false "update_Key2"
//@Param update_key3 query string false "update_Key3"
//@Param update_number query int false "update_Number"
//@Param update_gender query int false "update_Gender"
//@Param update_birth query int false "update_Birth"
//@Param update_injury_degree query string false "update_Injury_degree"
//@Param update_injury_position query int false "update_Injury_position"
//@Param update_protection query int false "update_Protection"
//@Param update_phone query int false "update_Phone"
//@Param update_person query string false "update_Person"
//@Param update_car query string false "update_Car"
//@Param update_action_status query int false "update_Action_status"
//@Param update_qualification query int false "update_Qualification"
//@Param update_license query int false "update_License"
//@Param update_drinking query int false "update_Drinking"
//@Param update_hit query int false "update_Hit"
//@Param update_city query string false "update_City"
//@Param update_position query string false "update_Position"
//@Param update_lane query string false "update_Lane"
//@Param update_death query string false "update_Death"
//@Param update_injured query string false "update_Injured"
//@Param update_death_exceed query string false "update_Death_exceed"
//@Param update_weather query string false "update_Weather"
//@Param update_light query string false "update_Light"
//@Param update_time_year query int false "update_Time_year"
//@Param update_time_month query string false "update_Time_month"
//@Param update_accident_chinese query string false "update_Accident_chinese"
//@Param update_anecdote_chinese query string false "update_Anecdote_chinese"
//@Success 200 {string} string "Successfully"
//@Failure 500 {object} models.Error "Internal Server Error"
//@Router /v1/update/{sql} [put]
func (c Controller) Update(MySQLDb *driver.MySQLDb, MsSQLDb *driver.MsSQLDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			error      models.Error
			people     models.People
			event      models.Event
			userRepo   repository.UserRepository
			conditions = make(map[string]interface{})
			update     = make(map[string]interface{})
		)

		//return map
		//func Vars(r *http.Request) map[string]string
		params := mux.Vars(r)

		//所有欄位
		MysqlWhereCol := []string{"where_key1", "where_key2", "where_key3",
			"where_number", "where_gender", "where_birth", "where_injury_degree", "where_injury_position",
			"where_protection", "where_phone", "where_person", "where_car", "where_action_status",
			"where_qualification", "where_license", "where_drinking", "where_hit"}
		MssqlWhereCol := []string{"where_key1", "where_key2", "where_key3",
			"where_city", "where_position", "where_lane", "where_death", "where_injured", "where_death_exceed",
			"where_weather", "where_light", "where_time_year", "where_time_month",
			"where_accident_chinese", "where_anecdote_chinese"}

		MysqlUpdateCol := []string{"key1", "key2", "key3",
			"number", "gender", "birth", "injury_degree", "injury_position",
			"protection", "phone", "person", "car", "action_status",
			"qualification", "license", "drinking", "hit"}
		MssqlUpdateCol := []string{"update_key1", "update_key2", "update_key3",
			"update_city", "update_position", "update_lane", "update_death", "update_injured", "update_death_exceed",
			"update_weather", "update_light", "update_time_year", "update_time_month",
			"update_accident_chinese", "update_anecdote_chinese"}

		switch strings.ToLower(params["sql"]) {
		case "mysql":
			//處理query參數
			for _, Colname := range MysqlWhereCol {
				value := r.URL.Query()[Colname]
				if len(value) > 0 {
					for i := 0; i < len(value); i++ {
						conditions[Colname] = value[i]
					}
				}
			}
			for _, Colname := range MysqlUpdateCol {
				value := r.URL.Query()[Colname]
				if len(value) > 0 {
					for i := 0; i < len(value); i++ {
						update[Colname] = value[i]
					}
				}
			}

			if err := userRepo.MysqlUpdateData(MySQLDb, people, conditions, update); err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			utils.SendSuccess(w, "Success!")

		case "mssql":
			//處理query參數
			for _, Colname := range MssqlWhereCol {
				value := r.URL.Query()[Colname]
				if len(value) > 0 {
					for i := 0; i < len(value); i++ {
						conditions[Colname] = value[i]
					}
				}
			}
			for _, Colname := range MssqlUpdateCol {
				value := r.URL.Query()[Colname]
				if len(value) > 0 {
					for i := 0; i < len(value); i++ {
						update[Colname] = value[i]
					}
				}
			}

			if err := userRepo.MssqlUpdateData(MsSQLDb, event, conditions, update); err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			utils.SendSuccess(w, "Success!")
		}
	}
}

//Delete delete
//@Summary delete value from database
//@Tags Data
//@Description 刪除資料庫數值
//@Accept json
//@Produce json
//@Param key1 query string false "Key1"
//@Param key2 query string false "Key2"
//@Param key3 query string false "Key3"
//@Param number query int false "Number"
//@Param gender query int false "Gender"
//@Param birth query int false "Birth"
//@Param injury_degree query string false "Injury_degree"
//@Param injury_position query int false "Injury_position"
//@Param protection query int false "Protection"
//@Param phone query int false "Phone"
//@Param person query string false "Person"
//@Param car query string false "Car"
//@Param action_status query int false "Action_status"
//@Param qualification query int false "Qualification"
//@Param license query int false "License"
//@Param drinking query int false "Drinking"
//@Param hit query int false "Hit"
//@Param city query string false "City"
//@Param position query string false "Position"
//@Param lane query string false "Lane"
//@Param death query string false "Death"
//@Param injured query string false "Injured"
//@Param death_exceed query string false "Death_exceed"
//@Param weather query string false "Weather"
//@Param light query string false "Light"
//@Param time_year query int false "Time_year"
//@Param time_month query string false "Time_month"
//@Param accident_chinese query string false "Accident_chinese"
//@Param anecdote_chinese query string false "Anecdote_chinese"
//@Success 200 {string} string "Successfully"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /v1/mysql/delete/{sql} [delete]
func (c Controller) Delete(MySQLDb *driver.MySQLDb, MsSQLDb *driver.MsSQLDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			conditions = make(map[string]interface{})
			error      models.Error
			people     models.People
			event      models.Event
			userRepo   repository.UserRepository
		)

		//return map
		//func Vars(r *http.Request) map[string]string
		params := mux.Vars(r)

		//所有欄位
		MysqlAllCol := []string{"key1", "key2", "key3",
			"number", "gender", "birth", "injury_degree", "injury_position",
			"protection", "phone", "person", "car", "action_status",
			"qualification", "license", "drinking", "hit"}
		MssqlAllCol := []string{"key1", "key2", "key3",
			"city", "position", "lane", "death", "injured", "death_exceed",
			"weather", "light", "time_year", "time_month",
			"accident_chinese", "anecdote_chinese"}

		switch strings.ToLower(params["sql"]) {
		case "mysql":
			//處理query參數
			for _, Colname := range MysqlAllCol {
				value := r.URL.Query()[Colname]
				if len(value) > 0 {
					for i := 0; i < len(value); i++ {
						conditions[Colname] = value[i]
					}
				}
			}

			if err := userRepo.MysqlDeleteData(MySQLDb, people, conditions); err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			utils.SendSuccess(w, "Success!")
		case "mssql":
			for _, Colname := range MssqlAllCol {
				value := r.URL.Query()[Colname]
				if len(value) > 0 {
					for i := 0; i < len(value); i++ {
						conditions[Colname] = value[i]
					}
				}
			}

			if err := userRepo.MssqlDeleteData(MsSQLDb, event, conditions); err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			utils.SendSuccess(w, "Success!")
		}
	}
}
