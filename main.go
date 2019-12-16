package main

import (
	"log"
	"net/http"
	"os"

	"ApiForTwoDb/controllers"
	"ApiForTwoDb/driver"
	"ApiForTwoDb/utils"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/subosito/gotenv"
)

var MySqlDb *driver.MySqlDb
var MsSqlDb *driver.MsSqlDb

var err error

func init() {

	//read .env file
	gotenv.Load()

	//設定資料庫資訊
	mysql := driver.MySqlUser{
		Host:     os.Getenv("mysql_host"), //主機
		MaxIdle:  10,                      //閒置的連接數
		MaxOpen:  10,                      //最大連接數
		User:     os.Getenv("mysql_user"), //用戶名
		Password: os.Getenv("mysql_pass"), //密碼
		Database: os.Getenv("mysql_name"), //資料庫名稱
		Port:     os.Getenv("mysql_port"), //端口
	}
	mssql := driver.MsSqlUser{
		Host:     os.Getenv("mssql_host"), //主機
		MaxIdle:  10,                      //閒置的連接數
		MaxOpen:  10,                      //最大連接數
		User:     os.Getenv("mssql_user"), //用戶名
		Password: os.Getenv("mssql_pass"), //密碼
		Database: os.Getenv("mssql_name"), //資料庫名稱
		Port:     os.Getenv("mssql_port"), //端口
	}

	//初始化連線
	MySqlDb = mysql.Init()
	MsSqlDb = mssql.Init()
}

func main() {
	//最後必須關閉
	defer MySqlDb.Close()
	defer MsSqlDb.Close()

	//create router
	//func NewRouter() *Router
	router := mux.NewRouter()
	controller := controllers.Controller{}

	//func (r *Router) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route
	//func (r *Router) Methods(methods ...string) *Route
	//mysql
	router.HandleFunc("/v1/mysql/signup", controller.MysqlSignup(MySqlDb)).Methods("POST")     //mysql註冊
	router.HandleFunc("/v1/mysql/login", controller.MysqlLogin(MySqlDb)).Methods("POST")       //mysql登入
	router.HandleFunc("/v1/mysql/addvalue", controller.MysqlAddValue(MySqlDb)).Methods("POST") //mysql插入值
	router.HandleFunc("/v1/mysql/getall", controller.MysqlGetAll(MySqlDb)).Methods("GET")      //mysql取得所有值
	router.HandleFunc("/v1/mysql/getsome", controller.MysqlGetSome(MySqlDb)).Methods("GET")    //mysql取得部分值
	router.HandleFunc("/v1/mysql/update", controller.MysqlUpdate(MySqlDb)).Methods("PUT")      //mysql更新值
	router.HandleFunc("/v1/mysql/delete", controller.MysqlDelete(MySqlDb)).Methods("DELETE")   //mysql刪除值
	//mssql
	router.HandleFunc("/v1/mssql/signup", controller.MssqlSignup(MsSqlDb)).Methods("POST")     //mssql註冊
	router.HandleFunc("/v1/mssql/login", controller.MssqlLogin(MsSqlDb)).Methods("POST")       //mssql登入
	router.HandleFunc("/v1/mssql/addvalue", controller.MssqlAddValue(MsSqlDb)).Methods("POST") //mysql插入值
	router.HandleFunc("/v1/mssql/getall", controller.MssqlGetAll(MsSqlDb)).Methods("GET")      //mssql取得所有值
	router.HandleFunc("/v1/mssql/getsome", controller.MssqlGetSome(MsSqlDb)).Methods("GET")    //mssql取得部分值
	router.HandleFunc("/v1/mssql/update", controller.MssqlUpdate(MsSqlDb)).Methods("PUT")      //mssql更新值
	router.HandleFunc("/v1/mssql/delete", controller.MssqlDelete(MsSqlDb)).Methods("DELETE")   //mysql刪除值

	//func (r *Router) Use(mc MiddlewareChain)
	//attach JWT auth middleware
	router.Use(utils.MysqlJwtAuthentication)
	router.Use(utils.MssqlJwtAuthentication)

	//localhost
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	//connect server
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
