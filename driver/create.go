package driver

import models "ApiForTwoDb/model"

//create table(if no table)
//mysql
func (db *MySqlDb) MysqlCreateTable(people models.People) error {
	//func (s *DB) CreateTable(models ...interface{}) *DB
	if err := db.CreateTable(&people).Error; err != nil {
		return err
	}
	return nil
}

//mssql
func (db *MsSqlDb) MssqlCreateTable(event models.Event) error {
	//func (s *DB) CreateTable(models ...interface{}) *DB
	if err := db.CreateTable(&event).Error; err != nil {
		return err
	}
	return nil
}
