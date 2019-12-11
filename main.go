package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"src/github.com/subosito/gotenv"
)

//mysql資料
type MySqlUser struct {
	Host string //主機
	//最大連接數
	MaxIdle  int
	MaxOpen  int
	User     string //用戶名
	Password string //密碼
	Database string //資料庫名稱
	Port     string //端口
}

//mssql資料
type MsSqlUser struct {
	Host string //主機
	//最大連接數
	MaxIdle  int
	MaxOpen  int
	User     string //用戶名
	Password string //密碼
	Database string //資料庫名稱
	Port     string //端口
}

//mysql peoples資料
type People struct {
	Key1            string
	Key2            string
	Key3            string
	Number          int
	Gender          int
	Birth           int
	Injury_degree   string
	Injury_position int
	Protection      int
	Phone           int
	Person          string
	Car             string
	Action_status   int
	Qualification   int
	License         int
	Drinking        int
	Hit             int
}

//mssql informations資料
type Event struct {
	Key1             string
	Key2             string
	Key3             string
	Hours            string
	City             string
	Position         string
	Lane             string
	Death            string
	Injured          string
	Death_exceed     string
	Weather          string
	Light            string
	Time_year        int
	Time_month       string
	Accident_chinese string
	Anecdote_chinese string
}

var MySqlDb *gorm.DB
var MsSqlDb *gorm.DB

func main() {

	var err error

	//read .env file
	gotenv.Load()

	//設定資料庫資訊
	mysql := MySqlUser{
		Host:     os.Getenv("mysql_host"), //主機
		MaxIdle:  10,                      //閒置的連接數
		MaxOpen:  10,                      //最大連接數
		User:     os.Getenv("mysql_user"), //用戶名
		Password: os.Getenv("mysql_pass"), //密碼
		Database: os.Getenv("mysql_name"), //資料庫名稱
		Port:     os.Getenv("mysql_port"), //端口
	}
	mssql := MsSqlUser{
		Host:     os.Getenv("mssql_host"), //主機
		MaxIdle:  10,                      //閒置的連接數
		MaxOpen:  10,                      //最大連接數
		User:     os.Getenv("mssql_user"), //用戶名
		Password: os.Getenv("mssql_pass"), //密碼
		Database: os.Getenv("mssql_name"), //資料庫名稱
		Port:     os.Getenv("mssql_port"), //端口
	}

	//完整的資料格式: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	MysqlDataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		mysql.User,
		mysql.Password,
		mysql.Host,
		mysql.Port,
		mysql.Database)

	//開啟資料庫連線
	MySqlDb, err = gorm.Open("mysql", MysqlDataSourceName)
	defer MySqlDb.Close()
	if err != nil {
		log.Println(err)
	}
	MySqlDb.DB().SetMaxIdleConns(mysql.MaxIdle)
	MySqlDb.DB().SetMaxOpenConns(mysql.MaxOpen)

	MssqlDataSourceName := fmt.Sprintf("serve=%s;user id=%s;password=%s;port=%s;database=%s",
		mssql.Host,
		mssql.User,
		mssql.Password,
		mssql.Port,
		mssql.Database)

	MsSqlDb, err = gorm.Open("mssql", MssqlDataSourceName)
	defer MsSqlDb.Close()
	if err != nil {
		log.Println(err)
	}
	MsSqlDb.DB().SetMaxIdleConns(mssql.MaxIdle)
	MsSqlDb.DB().SetMaxOpenConns(mssql.MaxOpen)

	//假設沒有資料表
	//func (s *DB) HasTable(value interface{}) bool
	if !MySqlDb.HasTable("peoples") {
		//如需創建新表(依據設定的struct)，使用AutoMigrate(指標)創建
		//只建立結構表
		//func (s *DB) AutoMigrate(values ...interface{}) *DB
		//MySqlDb.AutoMigrate(&People{})
		//func (s *DB) CreateTable(models ...interface{}) *DB
		//或使用MySqlDb.CreateTable(&People{})
		MySqlDb.CreateTable(&People{})
	}

	//假設沒有資料表
	//func (s *DB) HasTable(value interface{}) bool
	if !MsSqlDb.HasTable("datas") {
		//如需創建新表(依據設定的struct)，使用AutoMigrate(指標)創建
		//只建立結構表
		//func (s *DB) AutoMigrate(values ...interface{}) *DB
		//MsSqlDb.AutoMigrate(&Information{})
		//func (s *DB) CreateTable(models ...interface{}) *DB
		MsSqlDb.CreateTable(&Event{})
	}
	/*
		var peoples []People
		MySqlDb.Find(&peoples)
		fmt.Println(peoples)
	*/

	var events []Event
	MsSqlDb.Where("death=1").Find(&events)
	fmt.Println(events)
}
