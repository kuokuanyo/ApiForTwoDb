package repository

import (
	"ApiForTwoDb/driver"
	models "ApiForTwoDb/model"
)

//mssql
func (u UserRepository) MssqlInsertValue(MsSqlDb *driver.MsSqlDb, event models.Event) error {
	err := MsSqlDb.MssqlInsertValue(event)
	return err
}

func (u UserRepository) MssqlQueryAllData(MsSqlDb *driver.MsSqlDb, events []models.Event) ([]models.Event, error) {
	users, err := MsSqlDb.MssqlQueryAllData(events)
	return users, err
}

func (u UserRepository) MssqlQuerySomeData(MsSqlDb *driver.MsSqlDb, events []models.Event, conditions map[string]interface{}) ([]models.Event, error) {
	events, err := MsSqlDb.MssqlQuerySomeData(events, conditions)
	return events, err
}

func (u UserRepository) MssqlUpdateData(MsSqlDb *driver.MsSqlDb, event models.Event, conditions map[string]interface{}, update map[string]interface{}) error {
	err := MsSqlDb.MssqlUpdateData(event, conditions, update)
	return err
}

func (u UserRepository) MssqlDeleteData(MsSqlDb *driver.MsSqlDb, event models.Event, conditions map[string]interface{}) error {
	err := MsSqlDb.MssqlDeleteData(event, conditions)
	return err
}
