package controllers

import (
	"ApiForTwoDb/driver"
	"ApiForTwoDb/repository"
	"ApiForTwoDb/utils"
	"strconv"
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
//@Success 200 {string} string "Successfully"
//@Failure 500 {object} models.Error "Internal Server Error"
//@Router /v1/addvalue/{sql} [post]
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

//@Summary get all data from database
//@Tags Data
//@Description 從資料庫取得所有資料
//@Accept json
//@Produce json
//@Param sql path string true "資料庫引擎"
//@Success 200 {string} string "Successfully"
//@Failure 500 {object} models.Error "Internal Server Error"
//@Router /v1/getall/{sql} [post]
func (c Controller) GetAll(MySqlDb *driver.MySqlDb, MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
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
			peoples, err := userRepo.MysqlQueryAllData(MySqlDb, peoples)
			if err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			//成功返回
			utils.SendSuccess(w, peoples)
		case "mssql":
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
}

//@Summary get some data from database
//@Tags Data
//@Description 從資料庫取得部分資料
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
//@Success 200 {object} models.People "Successfully"
//@Failure 400 {object} models.Error "Bad Request"
//@Failure 500 {object} models.Error "Internal Server Error"
//@Router /v1/getsome/{sql} [get]
func (c Controller) GetSome(MySqlDb *driver.MySqlDb, MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
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

		switch strings.ToLower(params["sql"]) {
		case "mysql":
			//處理query參數
			key1 := r.URL.Query()["key1"]
			if len(key1) > 0 {
				conditions["key1"] = key1[0]
			}
			key2 := r.URL.Query()["key2"]
			if len(key2) > 0 {
				conditions["key2"] = key2[0]
			}
			key3 := r.URL.Query()["key3"]
			if len(key3) > 0 {
				conditions["key3"] = key3[0]
			}
			number := r.URL.Query()["number"]
			if len(number) > 0 {
				conditions["number"], _ = strconv.Atoi(number[0])
			}
			gender := r.URL.Query()["gender"]
			if len(gender) > 0 {
				conditions["gender"], _ = strconv.Atoi(gender[0])
			}
			birth := r.URL.Query()["birth"]
			if len(birth) > 0 {
				conditions["birth"], _ = strconv.Atoi(birth[0])
			}
			injury_degree := r.URL.Query()["injury_degree"]
			if len(injury_degree) > 0 {
				conditions["injury_degree"] = injury_degree[0]
			}
			injury_position := r.URL.Query()["injury_position"]
			if len(injury_position) > 0 {
				conditions["injury_position"], _ = strconv.Atoi(injury_position[0])
			}
			protection := r.URL.Query()["protection"]
			if len(protection) > 0 {
				conditions["protection"], _ = strconv.Atoi(protection[0])
			}
			phone := r.URL.Query()["phone"]
			if len(phone) > 0 {
				conditions["phone"], _ = strconv.Atoi(phone[0])
			}
			person := r.URL.Query()["person"]
			if len(person) > 0 {
				conditions["person"] = person[0]
			}
			car := r.URL.Query()["car"]
			if len(car) > 0 {
				conditions["car"] = car[0]
			}
			action_status := r.URL.Query()["action_status"]
			if len(action_status) > 0 {
				conditions["action_status"], _ = strconv.Atoi(action_status[0])
			}
			qualification := r.URL.Query()["qualification"]
			if len(qualification) > 0 {
				conditions["qualification"], _ = strconv.Atoi(qualification[0])
			}
			license := r.URL.Query()["license"]
			if len(license) > 0 {
				conditions["license"], _ = strconv.Atoi(license[0])
			}
			drinking := r.URL.Query()["drinking"]
			if len(drinking) > 0 {
				conditions["drinking"], _ = strconv.Atoi(drinking[0])
			}
			hit := r.URL.Query()["hit"]
			if len(hit) > 0 {
				conditions["hit"], _ = strconv.Atoi(hit[0])
			}

			peoples, err := userRepo.MysqlQuerySomeData(MySqlDb, peoples, conditions)
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
			key1 := r.URL.Query()["key1"]
			if len(key1) > 0 {
				conditions["key1"] = key1[0]
			}
			key2 := r.URL.Query()["key2"]
			if len(key2) > 0 {
				conditions["key2"] = key2[0]
			}
			key3 := r.URL.Query()["key3"]
			if len(key3) > 0 {
				conditions["key3"] = key3[0]
			}
			city := r.URL.Query()["city"]
			if len(city) > 0 {
				conditions["city"] = city[0]
			}
			position := r.URL.Query()["position"]
			if len(position) > 0 {
				conditions["position"] = position[0]
			}
			lane := r.URL.Query()["lane"]
			if len(lane) > 0 {
				conditions["lane"] = lane[0]
			}
			death := r.URL.Query()["death"]
			if len(death) > 0 {
				conditions["death"] = death[0]
			}
			injured := r.URL.Query()["injured"]
			if len(injured) > 0 {
				conditions["injured"] = injured[0]
			}
			death_exceed := r.URL.Query()["death_exceed"]
			if len(death_exceed) > 0 {
				conditions["death_exceed"] = death_exceed[0]
			}
			weather := r.URL.Query()["weather"]
			if len(weather) > 0 {
				conditions["weather"] = weather[0]
			}
			light := r.URL.Query()["light"]
			if len(light) > 0 {
				conditions["light"] = light[0]
			}
			time_year := r.URL.Query()["time_year"]
			if len(time_year) > 0 {
				conditions["time_year"], _ = strconv.Atoi(time_year[0])
			}
			time_month := r.URL.Query()["time_month"]
			if len(time_month) > 0 {
				conditions["time_month"] = time_month[0]
			}
			accident_chinese := r.URL.Query()["accident_chinese"]
			if len(accident_chinese) > 0 {
				conditions["accident_chinese"] = accident_chinese[0]
			}
			anecdote_chinese := r.URL.Query()["anecdote_chinese"]
			if len(anecdote_chinese) > 0 {
				conditions["anecdote_chinese"] = anecdote_chinese[0]
			}

			events, err := userRepo.MssqlQuerySomeData(MsSqlDb, events, conditions)
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

//@Summary update value
//@Tags Data
//@Description 更新資料庫數值
//@Accept json
//@Produce json
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
func (c Controller) Update(MySqlDb *driver.MySqlDb, MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
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

		switch strings.ToLower(params["sql"]) {
		case "mysql":
			//處理query參數
			where_key1 := r.URL.Query()["where_key1"]
			if len(where_key1) > 0 {
				conditions["key1"] = where_key1[0]
			}
			where_key2 := r.URL.Query()["where_key2"]
			if len(where_key2) > 0 {
				conditions["key2"] = where_key2[0]
			}
			where_key3 := r.URL.Query()["where_key3"]
			if len(where_key3) > 0 {
				conditions["key3"] = where_key3[0]
			}
			where_number := r.URL.Query()["where_number"]
			if len(where_number) > 0 {
				conditions["number"], _ = strconv.Atoi(where_number[0])
			}
			where_gender := r.URL.Query()["where_gender"]
			if len(where_gender) > 0 {
				conditions["gender"], _ = strconv.Atoi(where_gender[0])
			}
			where_birth := r.URL.Query()["where_birth"]
			if len(where_birth) > 0 {
				conditions["birth"], _ = strconv.Atoi(where_birth[0])
			}
			where_injury_degree := r.URL.Query()["where_injury_degree"]
			if len(where_injury_degree) > 0 {
				conditions["injury_degree"] = where_injury_degree[0]
			}
			where_injury_position := r.URL.Query()["where_injury_position"]
			if len(where_injury_position) > 0 {
				conditions["injury_position"], _ = strconv.Atoi(where_injury_position[0])
			}
			where_protection := r.URL.Query()["where_protection"]
			if len(where_protection) > 0 {
				conditions["protection"], _ = strconv.Atoi(where_protection[0])
			}
			where_phone := r.URL.Query()["where_phone"]
			if len(where_phone) > 0 {
				conditions["phone"], _ = strconv.Atoi(where_phone[0])
			}
			where_person := r.URL.Query()["where_person"]
			if len(where_person) > 0 {
				conditions["person"] = where_person[0]
			}
			where_car := r.URL.Query()["where_car"]
			if len(where_car) > 0 {
				conditions["car"] = where_car[0]
			}
			where_action_status := r.URL.Query()["where_action_status"]
			if len(where_action_status) > 0 {
				conditions["action_status"], _ = strconv.Atoi(where_action_status[0])
			}
			where_qualification := r.URL.Query()["where_qualification"]
			if len(where_qualification) > 0 {
				conditions["qualification"], _ = strconv.Atoi(where_qualification[0])
			}
			where_license := r.URL.Query()["where_license"]
			if len(where_license) > 0 {
				conditions["license"], _ = strconv.Atoi(where_license[0])
			}
			where_drinking := r.URL.Query()["where_drinking"]
			if len(where_drinking) > 0 {
				conditions["drinking"], _ = strconv.Atoi(where_drinking[0])
			}
			where_hit := r.URL.Query()["where_hit"]
			if len(where_hit) > 0 {
				conditions["hit"], _ = strconv.Atoi(where_hit[0])
			}

			update_key1 := r.URL.Query()["update_key1"]
			if len(update_key1) > 0 {
				update["key1"] = update_key1[0]
			}
			update_key2 := r.URL.Query()["update_key2"]
			if len(update_key2) > 0 {
				update["key2"] = update_key2[0]
			}
			update_key3 := r.URL.Query()["update_key3"]
			if len(update_key3) > 0 {
				update["key3"] = update_key3[0]
			}
			update_number := r.URL.Query()["update_number"]
			if len(update_number) > 0 {
				update["number"], _ = strconv.Atoi(update_number[0])
			}
			update_gender := r.URL.Query()["update_gender"]
			if len(update_gender) > 0 {
				update["gender"], _ = strconv.Atoi(update_gender[0])
			}
			update_birth := r.URL.Query()["update_birth"]
			if len(update_birth) > 0 {
				update["birth"], _ = strconv.Atoi(update_birth[0])
			}
			update_injury_degree := r.URL.Query()["update_injury_degree"]
			if len(update_injury_degree) > 0 {
				update["injury_degree"] = update_injury_degree[0]
			}
			update_injury_position := r.URL.Query()["update_injury_position"]
			if len(update_injury_position) > 0 {
				update["injury_position"], _ = strconv.Atoi(update_injury_position[0])
			}
			update_protection := r.URL.Query()["update_protection"]
			if len(update_protection) > 0 {
				update["protection"], _ = strconv.Atoi(update_protection[0])
			}
			update_phone := r.URL.Query()["update_phone"]
			if len(update_phone) > 0 {
				update["phone"], _ = strconv.Atoi(update_phone[0])
			}
			update_person := r.URL.Query()["update_person"]
			if len(update_person) > 0 {
				update["person"] = update_person[0]
			}
			update_car := r.URL.Query()["update_car"]
			if len(update_car) > 0 {
				update["car"] = update_car[0]
			}
			update_action_status := r.URL.Query()["update_action_status"]
			if len(update_action_status) > 0 {
				update["action_status"], _ = strconv.Atoi(update_action_status[0])
			}
			update_qualification := r.URL.Query()["update_qualification"]
			if len(update_qualification) > 0 {
				update["qualification"], _ = strconv.Atoi(update_qualification[0])
			}
			update_license := r.URL.Query()["update_license"]
			if len(update_license) > 0 {
				update["license"], _ = strconv.Atoi(update_license[0])
			}
			update_drinking := r.URL.Query()["update_drinking"]
			if len(update_drinking) > 0 {
				update["drinking"], _ = strconv.Atoi(update_drinking[0])
			}
			update_hit := r.URL.Query()["update_hit"]
			if len(update_hit) > 0 {
				update["hit"], _ = strconv.Atoi(update_hit[0])
			}

			if err := userRepo.MysqlUpdateData(MySqlDb, people, conditions, update); err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			utils.SendSuccess(w, "Success!")

		case "mssql":
			//處理query參數
			where_key1 := r.URL.Query()["where_key1"]
			if len(where_key1) > 0 {
				conditions["key1"] = where_key1[0]
			}
			where_key2 := r.URL.Query()["where_key2"]
			if len(where_key2) > 0 {
				conditions["key2"] = where_key2[0]
			}
			where_key3 := r.URL.Query()["where_key3"]
			if len(where_key3) > 0 {
				conditions["key3"] = where_key3[0]
			}
			where_city := r.URL.Query()["where_city"]
			if len(where_city) > 0 {
				conditions["city"] = where_city[0]
			}
			where_position := r.URL.Query()["where_position"]
			if len(where_position) > 0 {
				conditions["position"] = where_position[0]
			}
			where_lane := r.URL.Query()["where_lane"]
			if len(where_lane) > 0 {
				conditions["lane"] = where_lane[0]
			}
			where_death := r.URL.Query()["where_death"]
			if len(where_death) > 0 {
				conditions["death"] = where_death[0]
			}
			where_injured := r.URL.Query()["where_injured"]
			if len(where_injured) > 0 {
				conditions["injured"] = where_injured[0]
			}
			where_death_exceed := r.URL.Query()["where_death_exceed"]
			if len(where_death_exceed) > 0 {
				conditions["death_exceed"] = where_death_exceed[0]
			}
			where_weather := r.URL.Query()["where_weather"]
			if len(where_weather) > 0 {
				conditions["weather"] = where_weather[0]
			}
			where_light := r.URL.Query()["where_light"]
			if len(where_light) > 0 {
				conditions["light"] = where_light[0]
			}
			where_time_year := r.URL.Query()["where_time_year"]
			if len(where_time_year) > 0 {
				conditions["time_year"], _ = strconv.Atoi(where_time_year[0])
			}
			where_time_month := r.URL.Query()["where_time_month"]
			if len(where_time_month) > 0 {
				conditions["time_month"] = where_time_month[0]
			}
			where_accident_chinese := r.URL.Query()["where_accident_chinese"]
			if len(where_accident_chinese) > 0 {
				conditions["accident_chinese"] = where_accident_chinese[0]
			}
			where_anecdote_chinese := r.URL.Query()["where_anecdote_chinese"]
			if len(where_anecdote_chinese) > 0 {
				conditions["anecdote_chinese"] = where_anecdote_chinese[0]
			}

			update_key1 := r.URL.Query()["update_key1"]
			if len(update_key1) > 0 {
				update["key1"] = update_key1[0]
			}
			update_key2 := r.URL.Query()["update_key2"]
			if len(update_key2) > 0 {
				update["key2"] = update_key2[0]
			}
			update_key3 := r.URL.Query()["update_key3"]
			if len(update_key3) > 0 {
				update["key3"] = update_key3[0]
			}
			update_city := r.URL.Query()["update_city"]
			if len(update_city) > 0 {
				update["city"] = update_city[0]
			}
			update_position := r.URL.Query()["update_position"]
			if len(update_position) > 0 {
				update["position"] = update_position[0]
			}
			update_lane := r.URL.Query()["update_lane"]
			if len(update_lane) > 0 {
				update["lane"] = update_lane[0]
			}
			update_death := r.URL.Query()["update_death"]
			if len(update_death) > 0 {
				update["death"] = update_death[0]
			}
			update_injured := r.URL.Query()["update_injured"]
			if len(update_injured) > 0 {
				update["injured"] = update_injured[0]
			}
			update_death_exceed := r.URL.Query()["update_death_exceed"]
			if len(update_death_exceed) > 0 {
				update["death_exceed"] = update_death_exceed[0]
			}
			update_weather := r.URL.Query()["update_weather"]
			if len(update_weather) > 0 {
				update["weather"] = update_weather[0]
			}
			update_light := r.URL.Query()["update_light"]
			if len(update_light) > 0 {
				update["light"] = update_light[0]
			}
			update_time_year := r.URL.Query()["update_time_year"]
			if len(update_time_year) > 0 {
				update["time_year"], _ = strconv.Atoi(update_time_year[0])
			}
			update_time_month := r.URL.Query()["update_time_month"]
			if len(update_time_month) > 0 {
				update["time_month"] = update_time_month[0]
			}
			update_accident_chinese := r.URL.Query()["update_accident_chinese"]
			if len(update_accident_chinese) > 0 {
				update["accident_chinese"] = update_accident_chinese[0]
			}
			update_anecdote_chinese := r.URL.Query()["update_anecdote_chinese"]
			if len(update_anecdote_chinese) > 0 {
				update["anecdote_chinese"] = update_anecdote_chinese[0]
			}

			if err := userRepo.MssqlUpdateData(MsSqlDb, event, conditions, update); err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			utils.SendSuccess(w, "Success!")
		}
	}
}

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
//@Router /v1/mysql/delete/{key1}/{key2}/{key3} [delete]
func (c Controller) Delete(MySqlDb *driver.MySqlDb, MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
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

		switch strings.ToLower(params["sql"]) {
		case "mysql":
			//處理query參數
			key1 := r.URL.Query()["key1"]
			if len(key1) > 0 {
				conditions["key1"] = key1[0]
			}
			key2 := r.URL.Query()["key2"]
			if len(key2) > 0 {
				conditions["key2"] = key2[0]
			}
			key3 := r.URL.Query()["key3"]
			if len(key3) > 0 {
				conditions["key3"] = key3[0]
			}
			number := r.URL.Query()["number"]
			if len(number) > 0 {
				conditions["number"], _ = strconv.Atoi(number[0])
			}
			gender := r.URL.Query()["gender"]
			if len(gender) > 0 {
				conditions["gender"], _ = strconv.Atoi(gender[0])
			}
			birth := r.URL.Query()["birth"]
			if len(birth) > 0 {
				conditions["birth"], _ = strconv.Atoi(birth[0])
			}
			injury_degree := r.URL.Query()["injury_degree"]
			if len(injury_degree) > 0 {
				conditions["injury_degree"] = injury_degree[0]
			}
			injury_position := r.URL.Query()["injury_position"]
			if len(injury_position) > 0 {
				conditions["injury_position"], _ = strconv.Atoi(injury_position[0])
			}
			protection := r.URL.Query()["protection"]
			if len(protection) > 0 {
				conditions["protection"], _ = strconv.Atoi(protection[0])
			}
			phone := r.URL.Query()["phone"]
			if len(phone) > 0 {
				conditions["phone"], _ = strconv.Atoi(phone[0])
			}
			person := r.URL.Query()["person"]
			if len(person) > 0 {
				conditions["person"] = person[0]
			}
			car := r.URL.Query()["car"]
			if len(car) > 0 {
				conditions["car"] = car[0]
			}
			action_status := r.URL.Query()["action_status"]
			if len(action_status) > 0 {
				conditions["action_status"], _ = strconv.Atoi(action_status[0])
			}
			qualification := r.URL.Query()["qualification"]
			if len(qualification) > 0 {
				conditions["qualification"], _ = strconv.Atoi(qualification[0])
			}
			license := r.URL.Query()["license"]
			if len(license) > 0 {
				conditions["license"], _ = strconv.Atoi(license[0])
			}
			drinking := r.URL.Query()["drinking"]
			if len(drinking) > 0 {
				conditions["drinking"], _ = strconv.Atoi(drinking[0])
			}
			hit := r.URL.Query()["hit"]
			if len(hit) > 0 {
				conditions["hit"], _ = strconv.Atoi(hit[0])
			}
			if err := userRepo.MysqlDeleteData(MySqlDb, people, conditions); err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			utils.SendSuccess(w, "Success!")
		case "mssql":
			key1 := r.URL.Query()["key1"]
			if len(key1) > 0 {
				conditions["key1"] = key1[0]
			}
			key2 := r.URL.Query()["key2"]
			if len(key2) > 0 {
				conditions["key2"] = key2[0]
			}
			key3 := r.URL.Query()["key3"]
			if len(key3) > 0 {
				conditions["key3"] = key3[0]
			}
			city := r.URL.Query()["city"]
			if len(city) > 0 {
				conditions["city"] = city[0]
			}
			position := r.URL.Query()["position"]
			if len(position) > 0 {
				conditions["position"] = position[0]
			}
			lane := r.URL.Query()["lane"]
			if len(lane) > 0 {
				conditions["lane"] = lane[0]
			}
			death := r.URL.Query()["death"]
			if len(death) > 0 {
				conditions["death"] = death[0]
			}
			injured := r.URL.Query()["injured"]
			if len(injured) > 0 {
				conditions["injured"] = injured[0]
			}
			death_exceed := r.URL.Query()["death_exceed"]
			if len(death_exceed) > 0 {
				conditions["death_exceed"] = death_exceed[0]
			}
			weather := r.URL.Query()["weather"]
			if len(weather) > 0 {
				conditions["weather"] = weather[0]
			}
			light := r.URL.Query()["light"]
			if len(light) > 0 {
				conditions["light"] = light[0]
			}
			time_year := r.URL.Query()["time_year"]
			if len(time_year) > 0 {
				conditions["time_year"], _ = strconv.Atoi(time_year[0])
			}
			time_month := r.URL.Query()["time_month"]
			if len(time_month) > 0 {
				conditions["time_month"] = time_month[0]
			}
			accident_chinese := r.URL.Query()["accident_chinese"]
			if len(accident_chinese) > 0 {
				conditions["accident_chinese"] = accident_chinese[0]
			}
			anecdote_chinese := r.URL.Query()["anecdote_chinese"]
			if len(anecdote_chinese) > 0 {
				conditions["anecdote_chinese"] = anecdote_chinese[0]
			}
			if err := userRepo.MssqlDeleteData(MsSqlDb, event, conditions); err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			utils.SendSuccess(w, "Success!")
		}
	}
}
