package repository

import (
	"ApiForTwoDb/driver"
	models "ApiForTwoDb/model"
)

//find some col datas
//mysql
func (u UserRepository) MysqlQueryAllDataBySomeCol(db *driver.MySqlDb, peoples []models.People, col_name []string) ([]models.People, error) {
	if err := db.Select(col_name).Find(&peoples).Error; err != nil {
		return nil, err
	}
	return peoples, nil
}

//mssql
func (u UserRepository) MssqlQueryAllDataBySomeCol(db *driver.MsSqlDb, events []models.Event, col_name []string) ([]models.Event, error) {
	if err := db.Select(col_name).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

//find some datas
//mysql
func (u UserRepository) MysqlQuerySomeDataBySomeCol(db *driver.MySqlDb, peoples []models.People, col_name []string, condition map[string]interface{}) ([]models.People, error) {
	if err := db.Where(condition).Select(col_name).Find(&peoples).Error; err != nil {
		return nil, err
	}
	return peoples, nil
}

//find some datas
//mysql
func (u UserRepository) MssqlQuerySomeDataBySomeCol(db *driver.MsSqlDb, events []models.Event, col_name []string, condition map[string]interface{}) ([]models.Event, error) {
	if err := db.Where(condition).Select(col_name).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
