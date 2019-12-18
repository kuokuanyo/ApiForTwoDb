package driver

import (
	models "ApiForTwoDb/model"
)

//insert value
//mysql
func (db *MySqlDb) MysqlInsertValue(people models.People) error {
	if err := db.Create(&people).Error; err != nil {
		return err
	}
	return nil
}

//mssql
func (db *MsSqlDb) MssqlInsertValue(event models.Event) error {
	if err := db.Create(&event).Error; err != nil {
		return err
	}
	return nil
}

//insert the new user
//mysql
func (db *MySqlDb) MysqlInsertUser(user models.User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

//mssql
func (db *MsSqlDb) MssqlInsertUser(user models.User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
