package driver

import (
	models "ApiForTwoDb/model"
)

//update data
//mysql
func (db *MySqlDb) MysqlUpdateData(people models.People, condition map[string]interface{}, update map[string]interface{}) error {
	var err error
	err = db.Model(&people).Where(condition).Update(update).Error
	if err != nil {
		return err
	}
	return nil
}

//mysql
func (db *MsSqlDb) MssqlUpdateData(event models.Event, condition map[string]interface{}, update map[string]interface{}) error {
	err := db.Model(&event).Where(condition).Update(update).Error
	if err != nil {
		return err
	}
	return nil
}
