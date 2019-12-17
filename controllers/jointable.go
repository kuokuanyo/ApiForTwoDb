package controllers

import (
	"ApiForTwoDb/driver"
	models "ApiForTwoDb/model"
	"ApiForTwoDb/repository"
	"ApiForTwoDb/utils"
	"encoding/json"
	"net/http"
)

func (c Controller) JoinGetAll(MySqlDb *driver.MySqlDb, MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var peoples []models.People
		var events []models.Event
		var error models.Error
		var jointables []models.JoinTable

		userRepo := repository.UserRepository{}
		peoples, err := userRepo.MysqlQueryAllData(MySqlDb, peoples)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		events, err = userRepo.MssqlQueryAllData(MsSqlDb, events)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		for _, people := range peoples {
			for _, event := range events {
				var jointable models.JoinTable
				//關聯表欄位
				if people.Key1 == event.Key1 && people.Key2 == event.Key2 && people.Key3 == event.Key3 {
					jointable.Key1 = people.Key1
					jointable.Key2 = people.Key2
					jointable.Key3 = people.Key3
					jointable.Number = people.Number
					jointable.Gender = people.Gender
					jointable.Birth = people.Birth
					jointable.Injury_degree = people.Injury_degree
					jointable.Injury_position = people.Injury_position
					jointable.Protection = people.Protection
					jointable.Phone = people.Phone
					jointable.Person = people.Person
					jointable.Car = people.Car
					jointable.Action_status = people.Action_status
					jointable.Qualification = people.Qualification
					jointable.License = people.License
					jointable.Drinking = people.Drinking
					jointable.Hit = people.Hit
					jointable.City = event.City
					jointable.Position = event.Position
					jointable.Lane = event.Lane
					jointable.Death = event.Death
					jointable.Injured = event.Injured
					jointable.Death_exceed = event.Death_exceed
					jointable.Weather = event.Weather
					jointable.Light = event.Light
					jointable.Time_year = event.Time_year
					jointable.Time_month = event.Time_month
					jointable.Accident_chinese = event.Accident_chinese
					jointable.Anecdote_chinese = event.Anecdote_chinese
					jointables = append(jointables, jointable)
				}
			}
		}
		utils.SendSuccess(w, jointables)
	}
}

func (c Controller) JoinGetSome(MySqlDb *driver.MySqlDb, MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var peoples []models.People
		var events []models.Event
		var error models.Error
		var jointable models.JoinTable
		var jointables []models.JoinTable

		//decode
		json.NewDecoder(r.Body).Decode(&jointable)

		userRepo := repository.UserRepository{}
		peoples, err := userRepo.MysqlQueryAllData(MySqlDb, peoples)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		events, err = userRepo.MssqlQueryAllData(MsSqlDb, events)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		for _, people := range peoples {
			for _, event := range events {
				var jointable models.JoinTable
				//關聯表欄位
				if people.Key1 == event.Key1 && people.Key2 == event.Key2 && people.Key3 == event.Key3 {
					jointable.Key1 = people.Key1
					jointable.Key2 = people.Key2
					jointable.Key3 = people.Key3
					jointable.Number = people.Number
					jointable.Gender = people.Gender
					jointable.Birth = people.Birth
					jointable.Injury_degree = people.Injury_degree
					jointable.Injury_position = people.Injury_position
					jointable.Protection = people.Protection
					jointable.Phone = people.Phone
					jointable.Person = people.Person
					jointable.Car = people.Car
					jointable.Action_status = people.Action_status
					jointable.Qualification = people.Qualification
					jointable.License = people.License
					jointable.Drinking = people.Drinking
					jointable.Hit = people.Hit
					jointable.City = event.City
					jointable.Position = event.Position
					jointable.Lane = event.Lane
					jointable.Death = event.Death
					jointable.Injured = event.Injured
					jointable.Death_exceed = event.Death_exceed
					jointable.Weather = event.Weather
					jointable.Light = event.Light
					jointable.Time_year = event.Time_year
					jointable.Time_month = event.Time_month
					jointable.Accident_chinese = event.Accident_chinese
					jointable.Anecdote_chinese = event.Anecdote_chinese
					jointables = append(jointables, jointable)
				}
			}
		}
		utils.SendSuccess(w, jointables)
	}
}
