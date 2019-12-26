package repository

import (
	"ApiForTwoDb/driver"
	models "ApiForTwoDb/model"
)

//MysqlInsertValue mysql insert value
func (u UserRepository) MysqlInsertValue(MySQLDb *driver.MySQLDb, people models.People) error {
	err := MySQLDb.Create(&people).Error
	return err
}

//MssqlInsertValue mssql insert value
func (u UserRepository) MssqlInsertValue(MsSQLDb *driver.MsSQLDb, event models.Event) error {
	err := MsSQLDb.Create(&event).Error
	return err
}

//MysqlQueryAllData get all data
func (u UserRepository) MysqlQueryAllData(MySQLDb *driver.MySQLDb, peoples []models.People) ([]models.People, error) {
	err := MySQLDb.Find(&peoples).Error
	return peoples, err
}

//MssqlQueryAllData mssql get all data
func (u UserRepository) MssqlQueryAllData(MsSQLDb *driver.MsSQLDb, events []models.Event) ([]models.Event, error) {
	err := MsSQLDb.Find(&events).Error
	return events, err
}

//MysqlQuerySomeData mysql get some data
func (u UserRepository) MysqlQuerySomeData(MySQLDb *driver.MySQLDb, peoples []models.People, condition map[string]interface{}) ([]models.People, error) {
	err := MySQLDb.Where(condition).Find(&peoples).Error
	return peoples, err
}

//MssqlQuerySomeData mssql get some data
func (u UserRepository) MssqlQuerySomeData(MsSQLDb *driver.MsSQLDb, events []models.Event, condition map[string]interface{}) ([]models.Event, error) {
	err := MsSQLDb.Where(condition).Find(&events).Error
	return events, err
}

//MysqlUpdateData mysql update data
func (u UserRepository) MysqlUpdateData(MySQLDb *driver.MySQLDb, people models.People, condition map[string]interface{}, update map[string]interface{}) error {
	err := MySQLDb.Model(&people).Where(condition).Update(update).Error
	return err
}

//MssqlUpdateData mssql update data
func (u UserRepository) MssqlUpdateData(MsSQLDb *driver.MsSQLDb, event models.Event, condition map[string]interface{}, update map[string]interface{}) error {
	err := MsSQLDb.Model(&event).Where(condition).Update(update).Error
	return err
}

//MysqlDeleteData mysql delete data
func (u UserRepository) MysqlDeleteData(MySQLDb *driver.MySQLDb, people models.People, condition map[string]interface{}) error {
	err := MySQLDb.Where(condition).Delete(&people).Error
	return err
}

//MssqlDeleteData mssql delete data
func (u UserRepository) MssqlDeleteData(MsSQLDb *driver.MsSQLDb, event models.Event, condition map[string]interface{}) error {
	err := MsSQLDb.Where(condition).Delete(&event).Error
	return err
}
