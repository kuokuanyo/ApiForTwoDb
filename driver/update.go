package driver

import (
	models "ApiForTwoDb/model"
)

//update data
//mysql
func (db *MySqlDb) MysqlUpdateData(people models.People, condition map[string]interface{}, update map[string]interface{}) error {
	if err := db.Model(&people).Where(condition).Update(update).Error; err != nil {
		return err
	}
	return nil
}

//mysql
func (db *MsSqlDb) MssqlUpdateData(event models.Event, condition map[string]interface{}, update map[string]interface{}) error {
	if err := db.Model(&event).Where(condition).Update(update).Error; err != nil {
		return err
	}
	return nil
}
