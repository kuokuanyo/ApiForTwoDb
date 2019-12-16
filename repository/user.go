package repository

import (
	"ApiForTwoDb/driver"

	models "ApiForTwoDb/model"
)

type UserRepository struct{}

//mysql
func (u UserRepository) MysqlCheckSignup(MySqlDb *driver.MySqlDb, users []models.User) ([]models.User, error) {
	users, err := MySqlDb.MysqlQueryAllUser(users)
	return users, err
}

func (u UserRepository) MysqlInsertUser(MySqlDb *driver.MySqlDb, user models.User) error {
	err := MySqlDb.MysqlInsertUser(user)
	return err
}

func (u UserRepository) MysqlCheckLogin(MySqlDb *driver.MySqlDb, user models.User) (models.User, error) {
	user, err := MySqlDb.MysqlReadUser(user, "email =?", user.Email)
	return user, err
}

//mssql
func (u UserRepository) MssqlCheckSignup(MsSqlDb *driver.MsSqlDb, users []models.User) ([]models.User, error) {
	users, err := MsSqlDb.MssqlQueryAllUser(users)
	return users, err
}

func (u UserRepository) MssqlInsertUser(MsSqlDb *driver.MsSqlDb, user models.User) error {
	err := MsSqlDb.MssqlInsertUser(user)
	return err
}

func (u UserRepository) MssqlCheckLogin(MsSqlDb *driver.MsSqlDb, user models.User) (models.User, error) {
	user, err := MsSqlDb.MssqlReadUser(user, "email =?", user.Email)
	return user, err
}
