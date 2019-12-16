package driver

import (
	models "ApiForTwoDb/model"
)

//insert value
//mysql
func (db *MySqlDb) MysqlInsertValue(people models.People) error {
	err := db.Create(&people).Error
	if err != nil {
		return err
	}
	return nil
}

//mssql
func (db *MsSqlDb) MssqlInsertValue(event models.Event) error {
	err := db.Create(&event).Error
	if err != nil {
		return err
	}
	return nil
}

//insert the new user
//mysql
func (db *MySqlDb) MysqlInsertUser(user models.User) error {
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

//mssql
func (db *MsSqlDb) MssqlInsertUser(user models.User) error {
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
