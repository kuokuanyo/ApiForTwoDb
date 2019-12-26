package controllers

import (
	"ApiForTwoDb/driver"
	models "ApiForTwoDb/model"
	"ApiForTwoDb/repository"
	"ApiForTwoDb/utils"
	"fmt"
	"net/http"
	"strings"
)

//JoinGetAll join table get all
//@Summary get all data from join table
//@Tags JoinTable
//@Description 取得所有資料
//@Accept json
//@Produce json
//@Param number query string false "Number"
//@Param gender query string false "Gender"
//@Param birth query string false "Birth"
//@Param injury_degree query string false "Injury_degree"
//@Param injury_position query string false "Injury_position"
//@Param protection query string false "Protection"
//@Param phone query string false "Phone"
//@Param person query string false "Person"
//@Param car query string false "Car"
//@Param action_status query string false "Action_status"
//@Param qualification query string false "Qualification"
//@Param license query string false "License"
//@Param drinking query string false "Drinking"
//@Param hit query string false "Hit"
//@Param city query string false "City"
//@Param position query string false "Position"
//@Param lane query string false "Lane"
//@Param death query string false "Death"
//@Param injured query string false "Injured"
//@Param death_exceed query string false "Death_exceed"
//@Param weather query string false "Weather"
//@Param light query string false "Light"
//@Param time_year query string false "Time_year"
//@Param time_month query string false "Time_month"
//@Param accident_chinese query string false "Accident_chinese"
//@Param anecdote_chinese query string false "Anecdote_chinese"
//@Success 200 {object}  models.JoinTable "Successfully"
//@Failure 500 {object} models.Error "Internal Server Error"
//@Router /v1/join/getall [get]
func (c Controller) JoinGetAll(MySQLDb *driver.MySQLDb, MsSQLDb *driver.MsSQLDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			error      models.Error
			peoples    []models.People
			events     []models.Event
			userRepo   repository.UserRepository
			PeopleData []map[string]interface{}
			EventData  []map[string]interface{}
			JoinData   []map[string]interface{}
		)

		//所有欄位
		MysqlAllCol := []string{"number", "gender", "birth", "injury_degree", "injury_position",
			"protection", "phone", "person", "car", "action_status",
			"qualification", "license", "drinking", "hit"}
		MssqlAllCol := []string{"city", "position", "lane", "death", "injured", "death_exceed",
			"weather", "light", "time_year", "time_month",
			"accident_chinese", "anecdote_chinese"}

		//挑選的欄位
		//key為合併的欄位
		MysqlColName := []string{"key1", "key2", "key3"}
		MssqlColName := []string{"key1", "key2", "key3"}

		for _, ColName := range MysqlAllCol {
			value := r.URL.Query()[ColName]
			if len(value) > 0 {
				for i := 0; i < len(value); i++ {
					if value[i] == ColName {
						MysqlColName = append(MysqlColName, value[i])
					}
				}
			}
		}
		for _, ColName := range MssqlAllCol {
			value := r.URL.Query()[ColName]
			if len(value) > 0 {
				for i := 0; i < len(value); i++ {
					if value[i] == ColName {
						MssqlColName = append(MssqlColName, value[i])
					}
				}
			}
		}

		//read all datas from peoples
		//mysql
		peoples, err := userRepo.MysqlQueryAllDataBySomeCol(MySQLDb, peoples, MysqlColName)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		//將沒選擇的欄位刪除(轉換為map)
		for _, people := range peoples {
			var data = make(map[string]interface{})
			//convert map
			MapPeople := utils.StructToMap(people)
			for _, ColName := range MysqlColName {
				//字首大寫
				ColName = strings.Title(ColName)
				data[ColName] = MapPeople[ColName]
			}
			PeopleData = append(PeopleData, data)
		}

		//mssql
		events, err = userRepo.MssqlQueryAllDataBySomeCol(MsSQLDb, events, MssqlColName)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		for _, event := range events {
			var data = make(map[string]interface{})
			//convert map
			MapEvent := utils.StructToMap(event)
			for _, ColName := range MssqlColName {
				ColName = strings.Title(ColName)
				data[ColName] = MapEvent[ColName]
			}
			EventData = append(EventData, data)
		}

		//合併
		for _, people := range PeopleData {
			for _, event := range EventData {
				//以key合併
				if event["Key1"] == people["Key1"] && event["Key2"] == people["Key2"] && event["Key3"] == people["Key3"] {
					var data = make(map[string]interface{})
					data["Key1"] = event["Key1"]
					data["Key2"] = event["Key2"]
					data["Key3"] = event["Key3"]
					if len(MysqlColName) > 3 {
						for i := 3; i < len(MysqlColName); i++ {
							MysqlColName[i] = strings.Title(MysqlColName[i])
							data[MysqlColName[i]] = people[MysqlColName[i]]
						}
					}
					if len(MssqlColName) > 3 {
						for n := 3; n < len(MssqlColName); n++ {
							MssqlColName[n] = strings.Title(MssqlColName[n])
							data[MssqlColName[n]] = event[MssqlColName[n]]
						}
					}
					JoinData = append(JoinData, data)
				}
			}
		}
		utils.SendSuccess(w, JoinData)
	}
}

//JoinGetSome join table get some
//@Summary get some data from join table by condition
//@Tags JoinTable
//@Description 取得部分資料
//@Accept json
//@Produce json
//@Param number query string false "Number"
//@Param gender query string false "Gender"
//@Param birth query string false "Birth"
//@Param injury_degree query string false "Injury_degree"
//@Param injury_position query int false "Injury_position"
//@Param protection query string false "Protection"
//@Param phone query string false "Phone"
//@Param person query string false "Person"
//@Param car query string false "Car"
//@Param action_status query string false "Action_status"
//@Param qualification query string false "Qualification"
//@Param license query string false "License"
//@Param drinking query string false "Drinking"
//@Param hit query string false "Hit"
//@Param city query string false "City"
//@Param position query string false "Position"
//@Param lane query string false "Lane"
//@Param death query string false "Death"
//@Param injured query string false "Injured"
//@Param death_exceed query string false "Death_exceed"
//@Param weather query string false "Weather"
//@Param light query string false "Light"
//@Param time_year query string false "Time_year"
//@Param time_month query string false "Time_month"
//@Param accident_chinese query string false "Accident_chinese"
//@Param anecdote_chinese query string false "Anecdote_chinese"
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
//@Success 200 {object}  models.JoinTable "Successfully"
//@Failure 500 {object} models.Error "Internal Server Error"
//@Router /v1/join/getsome [get]
func (c Controller) JoinGetSome(MySQLDb *driver.MySQLDb, MsSQLDb *driver.MsSQLDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			error          models.Error
			peoples        []models.People
			events         []models.Event
			userRepo       repository.UserRepository
			PeopleData     []map[string]interface{}
			EventData      []map[string]interface{}
			JoinData       []map[string]interface{}
			MysqlCondition = make(map[string]interface{})
			MssqlCondition = make(map[string]interface{})
		)

		//所有欄位
		MysqlAllCol := []string{"key1", "key2", "key3",
			"number", "gender", "birth",
			"injury_degree", "injury_position",
			"protection", "phone", "person", "car", "action_status",
			"qualification", "license", "drinking", "hit"}
		MssqlAllCol := []string{"key1", "key2", "key3",
			"city", "position", "lane", "death", "injured", "death_exceed",
			"weather", "light", "time_year", "time_month",
			"accident_chinese", "anecdote_chinese"}

		MysqlAllCondition := []string{"where_key1", "where_key2", "where_key3",
			"where_number", "where_gender", "where_birth",
			"where_injury_degree", "where_injury_position",
			"where_protection", "where_phone", "where_person", "where_car", "where_action_status",
			"where_qualification", "where_license", "where_drinking", "where_hit"}
		MssqlAllCondition := []string{"where_key1", "where_key2", "where_key3",
			"where_city", "where_position", "where_lane", "where_death", "where_injured", "where_death_exceed",
			"where_weather", "where_light", "where_time_year", "where_time_month",
			"where_accident_chinese", "where_anecdote_chinese"}

		//挑選的欄位
		//key為合併的欄位
		MysqlColName := []string{"key1", "key2", "key3"}
		MssqlColName := []string{"key1", "key2", "key3"}

		//處理條件的欄位以及要挑選的欄位
		//mysql
		for _, ColName := range MysqlAllCol {
			value := r.URL.Query()[ColName]
			if len(value) > 0 {
				for i := 0; i < len(value); i++ {
					if value[i] == ColName {
						MysqlColName = append(MysqlColName, value[i])
					}
				}
			}
		}
		for _, ColName := range MysqlAllCondition {
			value := r.URL.Query()[ColName]
			if len(value) > 0 {
				for i := 0; i < len(value); i++ {
					MysqlCondition[ColName] = value[i]
				}
			}
		}

		//mssql
		for _, ColName := range MssqlAllCol {
			value := r.URL.Query()[ColName]
			if len(value) > 0 {
				for i := 0; i < len(value); i++ {
					if value[i] == ColName {
						MssqlColName = append(MssqlColName, value[i])
					}
				}
			}
		}
		for _, ColName := range MssqlAllCondition {
			value := r.URL.Query()[ColName]
			if len(value) > 0 {
				for i := 0; i < len(value); i++ {
					MssqlCondition[ColName] = value[i]
				}
			}
		}

		//read all datas from peoples
		//mysql
		peoples, err := userRepo.MysqlQuerySomeDataBySomeCol(MySQLDb, peoples, MysqlColName, MysqlCondition)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		//將沒選擇的欄位刪除(轉換為map)
		for _, people := range peoples {
			var data = make(map[string]interface{})
			//convert map
			MapPeople := utils.StructToMap(people)
			for _, ColName := range MysqlColName {
				//字首大寫
				ColName = strings.Title(ColName)
				data[ColName] = MapPeople[ColName]
			}
			PeopleData = append(PeopleData, data)
		}

		//mssql
		events, err = userRepo.MssqlQuerySomeDataBySomeCol(MsSQLDb, events, MssqlColName, MssqlCondition)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		for _, event := range events {
			var data = make(map[string]interface{})
			//convert map
			MapEvent := utils.StructToMap(event)
			for _, ColName := range MssqlColName {
				ColName = strings.Title(ColName)
				data[ColName] = MapEvent[ColName]
			}
			EventData = append(EventData, data)
		}

		//合併
		for _, people := range PeopleData {
			for _, event := range EventData {
				//以key合併
				if event["Key1"] == people["Key1"] && event["Key2"] == people["Key2"] && event["Key3"] == people["Key3"] {
					var data = make(map[string]interface{})
					data["Key1"] = event["Key1"]
					data["Key2"] = event["Key2"]
					data["Key3"] = event["Key3"]
					if len(MysqlColName) > 3 {
						for i := 3; i < len(MysqlColName); i++ {
							MysqlColName[i] = strings.Title(MysqlColName[i])
							data[MysqlColName[i]] = people[MysqlColName[i]]
						}
					}
					if len(MssqlColName) > 3 {
						for n := 3; n < len(MssqlColName); n++ {
							MssqlColName[n] = strings.Title(MssqlColName[n])
							data[MssqlColName[n]] = event[MssqlColName[n]]
						}
					}
					JoinData = append(JoinData, data)

				}
			}
		}
		fmt.Println(len(JoinData))
	}
}
