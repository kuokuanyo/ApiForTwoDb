package driver

import (
	models "ApiForTwoDb/model"
)

//delete data
//mysql
func (db *MySqlDb) MysqlDeleteData(people models.People, condition map[string]interface{}) error {
	err := db.Where(condition).Delete(&people).Error
	if err != nil {
		return err
	}

	return nil
}

//mssql
func (db *MsSqlDb) MssqlDeleteData(event models.Event, condition map[string]interface{}) error {
	err := db.Where(condition).Delete(&event).Error
	if err != nil {
		return err
	}
	return nil
}
