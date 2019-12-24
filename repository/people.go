package repository

import (
	"ApiForTwoDb/driver"
	models "ApiForTwoDb/model"
)

//mysql
func (u UserRepository) MysqlInsertValue(MySqlDb *driver.MySqlDb, people models.People) error {
	err := MySqlDb.MysqlInsertValue(people)
	return err
}

func (u UserRepository) MysqlQueryAllData(MySqlDb *driver.MySqlDb, peoples []models.People) ([]models.People, error) {
	users, err := MySqlDb.MysqlQueryAllData(peoples)
	return users, err
}

func (u UserRepository) MysqlQuerySomeData(MySqlDb *driver.MySqlDb, peoples []models.People, conditions map[string]interface{}) ([]models.People, error) {
	peoples, err := MySqlDb.MysqlQuerySomeData(peoples, conditions)
	return peoples, err
}

func (u UserRepository) MysqlUpdateData(MySqlDb *driver.MySqlDb, people models.People, conditions map[string]interface{}, update map[string]interface{}) error {
	err := MySqlDb.MysqlUpdateData(people, conditions, update)
	return err
}

func (u UserRepository) MysqlDeleteData(MySqlDb *driver.MySqlDb, people models.People, conditions map[string]interface{}) error {
	err := MySqlDb.MysqlDeleteData(people, conditions)
	return err
}
