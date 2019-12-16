package driver

import (
	models "ApiForTwoDb/model"
	"fmt"
)

//find all datas
//mysql
func (db *MySqlDb) MysqlQueryAllData(peoples []models.People) ([]models.People, error) {
	err := db.Find(&peoples).Error
	if err != nil {
		return []models.People{}, err
	}
	return peoples, nil
}

//mssql
func (db *MsSqlDb) MssqlQueryAllData(events []models.Event) ([]models.Event, error) {
	err := db.Find(&events).Error
	if err != nil {
		return []models.Event{}, err
	}
	return events, nil
}

//find datas by some condition
//mysql
func (db *MySqlDb) MysqlQuerySomeData(peoples []models.People, condition map[string]interface{}) ([]models.People, error) {
	err := db.Where(condition).Find(&peoples).Error
	if err != nil {
		return []models.People{}, err
	}
	return peoples, nil
}

//mssql
func (db *MsSqlDb) MssqlQuerySomeData(events []models.Event, condition map[string]interface{}, args ...interface{}) ([]models.Event, error) {
	err := db.Where(condition).Find(&events).Error
	if err != nil {
		return []models.Event{}, err
	}
	return events, nil
}

//find one data
//mysql
func (db *MySqlDb) MysqlQueryOneData(people models.People, condition string, args ...interface{}) (models.People, error) {
	err := db.Where(condition, args...).First(&people).Error
	if err != nil {
		return models.People{}, err
	}
	return people, nil
}

//mssql
func (db *MsSqlDb) MssqlQueryOneData(event models.Event, order string, condition string, args ...interface{}) (models.Event, error) {
	err := db.Order(order).Where(condition, args...).First(&event).Error
	fmt.Println(event)
	if err != nil {
		return models.Event{}, err
	}
	return event, nil
}

//get all user data
//mysql
func (db *MySqlDb) MysqlQueryAllUser(users []models.User) ([]models.User, error) {
	err := db.Find(&users).Error
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

//mssql
func (db *MsSqlDb) MssqlQueryAllUser(users []models.User) ([]models.User, error) {
	err := db.Find(&users).Error
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

//read user
//mysql
func (db *MySqlDb) MysqlReadUser(user models.User, condition string, args ...interface{}) (models.User, error) {
	err := db.Where(condition, args...).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

//mysql
func (db *MsSqlDb) MssqlReadUser(user models.User, condition string, args ...interface{}) (models.User, error) {
	err := db.Where(condition, args...).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
