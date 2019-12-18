package driver

import (
	models "ApiForTwoDb/model"
)

//delete data
//mysql
func (db *MySqlDb) MysqlDeleteData(people models.People, condition map[string]interface{}) error {
	if err := db.Where(condition).Delete(&people).Error; err != nil {
		return err
	}
	return nil
}

//mssql
func (db *MsSqlDb) MssqlDeleteData(event models.Event, condition map[string]interface{}) error {
	if err := db.Where(condition).Delete(&event).Error; err != nil {
		return err
	}
	return nil
}
