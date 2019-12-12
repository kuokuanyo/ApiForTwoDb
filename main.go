package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/subosito/gotenv"
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
	Key1            string `gorm:"column:key1" json:"key1"`
	Key2            string `gorm:"column:key2" json:"key2"`
	Key3            string `gorm:"column:key3" json:"key3"`
	Number          int    `gorm:"column:number" json:"number"`
	Gender          int    `gorm:"column:gender" json:"gender"`
	Birth           int    `gorm:"column:birth" json:"birth"`
	Injury_degree   string `gorm:"column:injury_degree" json:"injury_degree"`
	Injury_position int    `gorm:"column:injury_position" json:"injury_position"`
	Protection      int    `gorm:"column:protection" json:"protection"`
	Phone           int    `gorm:"column:phone" json:"phone"`
	Person          string `gorm:"column:person" json:"person"`
	Car             string `gorm:"column:car" json:"car"`
	Action_status   int    `gorm:"column:action_status" json:"action_status"`
	Qualification   int    `gorm:"column:qualification" json:"qualification"`
	License         int    `gorm:"column:license" json:"license"`
	Drinking        int    `gorm:"column:drinking" json:"drinking"`
	Hit             int    `gorm:"column:hit" json:"hit"`
}

//mssql informations資料
type Event struct {
	Key1             string `gorm:"column:key1" json:"key1"`
	Key2             string `gorm:"column:key2" json:"key2"`
	Key3             string `gorm:"column:key3" json:"key3"`
	City             string `gorm:"column:city" json:"city"`
	Position         string `gorm:"column:position" json:"position"`
	Lane             string `gorm:"column:lane" json:"lane"`
	Death            string `gorm:"column:death" json:"death"`
	Injured          string `gorm:"column:injured" json:"injured"`
	Death_exceed     string `gorm:"column:death_exceed" json:"death_exceed"`
	Weather          string `gorm:"column:weather" json:"weather"`
	Light            string `gorm:"column:light" json:"light"`
	Time_year        int    `gorm:"column:time_year" json:"time_year"`
	Time_month       string `gorm:"column:time_month" json:"time_month"`
	Accident_chinese string `gorm:"column:accident_chinese" json:"accident_chinese"`
	Anecdote_chinese string `gorm:"column:anecdote_chinese" json:"anecdote_chinese"`
}

//test
type T struct {
	Id   int
	Name string
	Math int
	Eng  int
	Pe   int
}

var MySqlDb *gorm.DB
var MsSqlDb *gorm.DB

func main() {

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
	MySqlDb, err := gorm.Open("mysql", MysqlDataSourceName)
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
		MysqlCreateTable(MySqlDb, People{})
	}

	//假設沒有資料表
	//func (s *DB) HasTable(value interface{}) bool
	if !MsSqlDb.HasTable("events") {
		MssqlCreateTable(MySqlDb, Event{})
	}

	//create router
	//func NewRouter() *Router
	router := mux.NewRouter()
	//func (r *Router) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route
	//func (r *Router) Methods(methods ...string) *Route

	//localhost
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	//connect server
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

//create table(if no table)
//mysql
func MysqlCreateTable(db *gorm.DB, people People) error {
	var err error
	//func (s *DB) CreateTable(models ...interface{}) *DB
	err = db.CreateTable(&people).Error
	if err != nil {
		return err
	}
	return nil
}

//mssql
func MssqlCreateTable(db *gorm.DB, event Event) error {
	var err error
	//func (s *DB) CreateTable(models ...interface{}) *DB
	err = db.CreateTable(&event).Error
	if err != nil {
		return err
	}
	return nil
}

//insert value
//mysql
func MysqlInsertValue(db *gorm.DB, people People) error {
	var err error
	err = db.Create(&people).Error
	if err != nil {
		return err
	}
	return nil
}

//mssql
func MssqlInsertValue(db *gorm.DB, event Event) error {
	var err error
	err = db.Create(&event).Error
	if err != nil {
		return err
	}
	return nil
}

//find all datas
//mysql
func MysqlQueryAllData(db *gorm.DB, peoples []People) ([]People, error) {
	var err error
	err = db.Find(&peoples).Error
	if err != nil {
		return peoples, err
	}
	return peoples, nil
}

//mssql
func MssqlQueryAllData(db *gorm.DB, events []Event) ([]Event, error) {
	var err error
	err = db.Find(&events).Error
	if err != nil {
		return events, err
	}
	return events, nil
}

//find datas by some condition
//mysql
func MysqlQuerySomeData(db *gorm.DB, peoples []People, condition map[string]interface{}) ([]People, error) {
	var err error
	err = db.Where(condition).Find(&peoples).Error
	if err != nil {
		return peoples, err
	}
	return peoples, nil
}

//mssql
func MssqlQuerySomeData(db *gorm.DB, events []Event, condition map[string]interface{}, args ...interface{}) ([]Event, error) {
	var err error
	err = db.Where(condition).Find(&events).Error
	if err != nil {
		return events, err
	}
	return events, nil
}

//fund one data
//mysql
func MysqlQueryOneData(db *gorm.DB, people People, condition string, args ...interface{}) (People, error) {
	var err error
	err = db.Where(condition, args...).First(&people).Error
	if err != nil {
		return people, err
	}
	return people, nil
}

//mssql
func MssqlQueryOneData(db *gorm.DB, event Event, order string, condition string, args ...interface{}) (Event, error) {
	var err error
	err = db.Order(order).Where(condition, args...).First(&event).Error
	fmt.Println(event)
	if err != nil {
		return event, err
	}
	return event, nil
}

//update data
//mysql
func MysqlUpdateData(db *gorm.DB, people People, condition map[string]interface{}, update map[string]interface{}) error {
	var err error
	err = db.Model(&people).Where(condition).Update(update).Error
	if err != nil {
		return err
	}
	return nil
}

//mysql
func MssqlUpdateData(db *gorm.DB, event Event, condition map[string]interface{}, update map[string]interface{}) error {
	var err error
	err = db.Model(&event).Where(condition).Update(update).Error
	if err != nil {
		return err
	}
	return nil
}

//delete data
//mysql
func MysqlDeleteData(db *gorm.DB, people People, condition map[string]interface{}) error {
	var err error
	err = db.Where(condition).Delete(&people).Error
	if err != nil {
		return err
	}
	return nil
}

//mssql
func MssqlDeleteData(db *gorm.DB, event Event, condition map[string]interface{}) error {
	var err error
	err = db.Where(condition).Delete(&event).Error
	if err != nil {
		return err
	}
	return nil
}
