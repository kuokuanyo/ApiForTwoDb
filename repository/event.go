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

func (u UserRepository) MssqlQuerySomeData(MsSqlDb *driver.MsSqlDb, events []models.Event, event models.Event) ([]models.Event, error) {
	events, err := MsSqlDb.MssqlQuerySomeData(events,
		map[string]interface{}{"key1": event.Key1,
			"key2": event.Key2,
			"key3": event.Key3})
	return events, err
}

func (u UserRepository) MssqlUpdateData(MsSqlDb *driver.MsSqlDb, event models.Event) error {
	err := MsSqlDb.MssqlUpdateData(event,
		map[string]interface{}{"key1": event.Key1},
		map[string]interface{}{"birth": event.Death})
	return err
}

func (u UserRepository) MssqlDeleteData(MsSqlDb *driver.MsSqlDb, event models.Event) error {
	err := MsSqlDb.MssqlDeleteData(event, map[string]interface{}{"key1": event.Key1})
	return err
}
