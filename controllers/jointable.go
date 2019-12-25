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

//@Summary get all data from join table
//@Tags JoinTable
//@Description 取得所有資料
//@Accept json
//@Produce json
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
//@Success 200 {object} JoinData "Successfully"
//@Failure 500 {object} models.Error "Internal Server Error"
//@Router /v1/join/getall [get]
func (c Controller) JoinGetAll(MySqlDb *driver.MySqlDb, MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
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
				if value[0] == ColName {
					MysqlColName = append(MysqlColName, value[0])
				}
			}
		}

		for _, ColName := range MssqlAllCol {
			value := r.URL.Query()[ColName]
			if len(value) > 0 {
				if value[0] == ColName {
					MssqlColName = append(MssqlColName, value[0])
				}
			}
		}

		//read all datas from peoples
		//mysql
		peoples, err := userRepo.MysqlQueryAllDataBySomeCol(MySqlDb, peoples, MysqlColName)
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
		events, err = userRepo.MssqlQueryAllDataBySomeCol(MsSqlDb, events, MssqlColName)
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
					} else {
						continue
					}
					if len(MssqlColName) > 3 {
						for n := 3; n < len(MssqlColName); n++ {
							MssqlColName[n] = strings.Title(MssqlColName[n])
							data[MssqlColName[n]] = event[MssqlColName[n]]
						}
					} else {
						continue
					}
					JoinData = append(JoinData, data)
				}
			}
		}
		utils.SendSuccess(w, JoinData)
	}
}

//@Summary get all data from join table
//@Tags JoinTable
//@Description 取得所有資料
//@Accept json
//@Produce json
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
//@Success 200 {object} JoinData "Successfully"
//@Failure 500 {object} models.Error "Internal Server Error"
//@Router /v1/join/getall [get]
func (c Controller) JoinGetSome(MySqlDb *driver.MySqlDb, MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
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

		//挑選的欄位
		//key為合併的欄位
		MysqlColName := []string{"key1", "key2", "key3"}
		MssqlColName := []string{"key1", "key2", "key3"}

		//處理條件的欄位以及要挑選的欄位
		//mysql
		for _, ColName := range MysqlAllCol {
			value := r.URL.Query()[ColName]
			if len(value) == 0 {
				continue
			}
			for i := 0; i < len(value); i++ {
				if i < 3 {
					if value[i] == ColName {
						continue
					} else {
						MysqlCondition[ColName] = value[i]
					}
				}
				if value[i] == ColName {
					MysqlColName = append(MysqlColName, value[i])
				} else {
					MysqlCondition[ColName] = value[i]
				}
			}
		}
		//mssql
		for _, ColName := range MssqlAllCol {
			value := r.URL.Query()[ColName]
			if len(value) == 0 {
				continue
			}
			for i := 0; i < len(value); i++ {
				if i < 3 {
					if value[i] == ColName {
						continue
					} else {
						MssqlCondition[ColName] = value[i]
					}
				}
				if value[i] == ColName {
					MssqlColName = append(MysqlColName, value[i])
				} else {
					MssqlCondition[ColName] = value[i]
				}
			}
		}
		fmt.Println(MysqlColName)
		fmt.Println(MysqlCondition)
		//read all datas from peoples
		//mysql
		peoples, err := userRepo.MysqlQuerySomeDataBySomeCol(MySqlDb, peoples, MysqlColName, MysqlCondition)
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
		events, err = userRepo.MssqlQuerySomeDataBySomeCol(MsSqlDb, events, MssqlColName, MssqlCondition)
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
					} else {
						continue
					}
					if len(MssqlColName) > 3 {
						for n := 3; n < len(MssqlColName); n++ {
							MssqlColName[n] = strings.Title(MssqlColName[n])
							data[MssqlColName[n]] = event[MssqlColName[n]]
						}
					} else {
						continue
					}
					JoinData = append(JoinData, data)
				}
			}
		}
		utils.SendSuccess(w, JoinData)
	}
}
