//@title Restful API(connect two databases)
//@version 1.0.0
//@description Define an API
//@Schemes http
//@host localhost:8080
//@BasePath /v1
package main

import (
	"log"
	"net/http"
	"os"

	"ApiForTwoDb/controllers"
	"ApiForTwoDb/driver"
	"ApiForTwoDb/utils"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

//MySQLDb is mysql engine
var MySQLDb *driver.MySQLDb

//MsSQLDb is mssql engine
var MsSQLDb *driver.MsSQLDb

func init() {
	//read .env file
	gotenv.Load()

	//設定資料庫資訊
	mysql := driver.MySQLUser{
		Host:     os.Getenv("mysql_host"), //主機
		MaxIdle:  10,                      //閒置的連接數
		MaxOpen:  10,                      //最大連接數
		User:     os.Getenv("mysql_user"), //用戶名
		Password: os.Getenv("mysql_pass"), //密碼
		Database: os.Getenv("mysql_name"), //資料庫名稱
		Port:     os.Getenv("mysql_port"), //端口
	}
	mssql := driver.MsSQLUser{
		Host:     os.Getenv("mssql_host"), //主機
		MaxIdle:  10,                      //閒置的連接數
		MaxOpen:  10,                      //最大連接數
		User:     os.Getenv("mssql_user"), //用戶名
		Password: os.Getenv("mssql_pass"), //密碼
		Database: os.Getenv("mssql_name"), //資料庫名稱
		Port:     os.Getenv("mssql_port"), //端口
	}

	//初始化連線
	MySQLDb = mysql.Init()
	MsSQLDb = mssql.Init()
}

func main() {
	//最後必須關閉
	defer MySQLDb.Close()
	defer MsSQLDb.Close()

	//create router
	//func NewRouter() *Router
	router := mux.NewRouter()
	controller := controllers.Controller{}

	//func (r *Router) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route
	//func (r *Router) Methods(methods ...string) *Route
	//註冊及登入
	router.HandleFunc("/v1/signup/{sql}", controller.Signup(MySQLDb, MsSQLDb)).Methods("POST") //註冊
	router.HandleFunc("/v1/login/{sql}", controller.Login(MySQLDb, MsSQLDb)).Methods("POST")   //登入

	//CRUD
	router.HandleFunc("/v1/addvalue/{sql}", controller.AddValue(MySQLDb, MsSQLDb)).Methods("POST") //插入值
	router.HandleFunc("/v1/getall/{sql}", controller.GetAll(MySQLDb, MsSQLDb)).Methods("GET")      //取得所有值
	router.HandleFunc("/v1/getsome/{sql}", controller.GetSome(MySQLDb, MsSQLDb)).Methods("GET")    //取得部分值
	router.HandleFunc("/v1/update/{sql}", controller.Update(MySQLDb, MsSQLDb)).Methods("PUT")      //更新值
	router.HandleFunc("/v1/delete/{sql}", controller.Delete(MySQLDb, MsSQLDb)).Methods("DELETE")   //刪除值

	//join table
	router.HandleFunc("/v1/join/getall", controller.JoinGetAll(MySQLDb, MsSQLDb)).Methods("GET")   //mssql取得所有值
	router.HandleFunc("/v1/join/getsome", controller.JoinGetSome(MySQLDb, MsSQLDb)).Methods("GET") //mssql取得部分值

	//安全性驗證
	//func (r *Router) Use(mc MiddlewareChain)
	//attach JWT auth middleware
	router.Use(utils.JwtAuthentication)

	//伺服器連線
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
