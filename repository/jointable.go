package repository

import (
	"ApiForTwoDb/driver"
	models "ApiForTwoDb/model"
)

//更新資料表
//關聯表
func (u UserRepository) UpdateJoinData(MySqlDb *driver.MySqlDb, MsSqlDb *driver.MsSqlDb, peoples []models.People, events []models.Event) error {
	//var jointables []models.JoinTable
	var err error

	//創建新table
	if !MySqlDb.HasTable(&models.JoinTable{}) {
		MySqlDb.CreateTable(&models.JoinTable{})
	} else {
		MySqlDb.DropTable(&models.JoinTable{})
		MySqlDb.CreateTable(&models.JoinTable{})
	}

	//read all datas from peoples
	peoples, err = MySqlDb.MysqlQueryAllData(peoples)
	if err != nil {
		return err
	}
	//read all datas from events
	events, err = MsSqlDb.MssqlQueryAllData(events)
	if err != nil {
		return err
	}

	for _, people := range peoples {
		for _, event := range events {
			//關聯表欄位
			if people.Key1 == event.Key1 && people.Key2 == event.Key2 && people.Key3 == event.Key3 {
				var jointable models.JoinTable
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
				//插入數值
				err := MySqlDb.Create(&jointable).Error
				if err != nil {
					return err
				}
				//jointables = append(jointables, jointable)
			}
		}
	}
	return nil
}

func (u UserRepository) QueryAllJoinData(MySqlDb *driver.MySqlDb, jointables []models.JoinTable) ([]models.JoinTable, error) {
	jointables, err := MySqlDb.QueryAllJoinData(jointables)
	return jointables, err
}

func (u UserRepository) QuerySomeJoinData(MySqlDb *driver.MySqlDb, jointables []models.JoinTable, jointable models.JoinTable) ([]models.JoinTable, error) {
	jointables, err := MySqlDb.QuerySomeJoinData(jointables,
		map[string]interface{}{"key1": jointable.Key1,
			"key2": jointable.Key2,
			"key3": jointable.Key3})
	return jointables, err
}
