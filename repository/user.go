package repository

import (
	"ApiForTwoDb/driver"

	models "ApiForTwoDb/model"
)

//UserRepository struct
type UserRepository struct{}

//MysqlCheckSignup mysql get all users
func (u UserRepository) MysqlCheckSignup(MySQLDb *driver.MySQLDb, users []models.User) ([]models.User, error) {
	err := MySQLDb.Find(&users).Error
	return users, err
}

//MysqlInsertUser mysql inert a ner user
func (u UserRepository) MysqlInsertUser(MySQLDb *driver.MySQLDb, user models.User) error {
	err := MySQLDb.Create(&user).Error
	return err
}

//MysqlCheckLogin mysql check log in
func (u UserRepository) MysqlCheckLogin(MySQLDb *driver.MySQLDb, user models.User) (models.User, error) {
	err := MySQLDb.Where("email=?", user.Email).First(&user).Error
	return user, err
}

//MssqlCheckSignup mssql get all users
func (u UserRepository) MssqlCheckSignup(MsSQLDb *driver.MsSQLDb, users []models.User) ([]models.User, error) {
	err := MsSQLDb.Find(&users).Error
	return users, err
}

//MssqlInsertUser mssql inert a ner user
func (u UserRepository) MssqlInsertUser(MsSQLDb *driver.MsSQLDb, user models.User) error {
	err := MsSQLDb.Create(&user).Error
	return err
}

//MssqlCheckLogin mssql check log in
func (u UserRepository) MssqlCheckLogin(MsSQLDb *driver.MsSQLDb, user models.User) (models.User, error) {
	err := MsSQLDb.Where("email=?", user.Email).First(&user).Error
	return user, err
}
