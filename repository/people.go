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

func (u UserRepository) MysqlQuerySomeData(MySqlDb *driver.MySqlDb, peoples []models.People, people models.People) ([]models.People, error) {
	peoples, err := MySqlDb.MysqlQuerySomeData(peoples, map[string]interface{}{"key1": people.Key1})
	return peoples, err
}

func (u UserRepository) MysqlUpdateData(MySqlDb *driver.MySqlDb, people models.People) error {
	err := MySqlDb.MysqlUpdateData(people,
		map[string]interface{}{"key1": people.Key1},
		map[string]interface{}{"birth": people.Birth})
	return err
}

func (u UserRepository) MysqlDeleteData(MySqlDb *driver.MySqlDb, people models.People) error {
	err := MySqlDb.MysqlDeleteData(people, map[string]interface{}{"key1": people.Key1})
	return err
}
