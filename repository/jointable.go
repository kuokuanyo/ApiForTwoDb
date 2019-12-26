package repository

import (
	"ApiForTwoDb/driver"
	models "ApiForTwoDb/model"
)

//MysqlQueryAllDataBySomeCol mysql find some col datas
func (u UserRepository) MysqlQueryAllDataBySomeCol(db *driver.MySQLDb, peoples []models.People, colName []string) ([]models.People, error) {
	if err := db.Select(colName).Find(&peoples).Error; err != nil {
		return nil, err
	}
	return peoples, nil
}

//MssqlQueryAllDataBySomeCol mssql find some col datas
func (u UserRepository) MssqlQueryAllDataBySomeCol(db *driver.MsSQLDb, events []models.Event, colName []string) ([]models.Event, error) {
	if err := db.Select(colName).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

//MysqlQuerySomeDataBySomeCol mysql find some datas
func (u UserRepository) MysqlQuerySomeDataBySomeCol(db *driver.MySQLDb, peoples []models.People, colName []string, condition map[string]interface{}) ([]models.People, error) {
	if err := db.Where(condition).Select(colName).Find(&peoples).Error; err != nil {
		return nil, err
	}
	return peoples, nil
}

//MssqlQuerySomeDataBySomeCol mssql find some datas
func (u UserRepository) MssqlQuerySomeDataBySomeCol(db *driver.MsSQLDb, events []models.Event, colName []string, condition map[string]interface{}) ([]models.Event, error) {
	if err := db.Where(condition).Select(colName).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
