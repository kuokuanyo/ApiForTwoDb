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

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

//資料庫引擎
var MySqlDb *driver.MySqlDb
var MsSqlDb *driver.MsSqlDb

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
	//註冊及登入
	router.HandleFunc("/v1/signup/{sql}", controller.Signup(MySqlDb, MsSqlDb)).Methods("POST") //註冊
	router.HandleFunc("/v1/login/{sql}", controller.Login(MySqlDb, MsSqlDb)).Methods("POST")   //登入

	//CRUD
	router.HandleFunc("/v1/addvalue/{sql}", controller.AddValue(MySqlDb, MsSqlDb)).Methods("POST") //插入值

	//mysql
	router.HandleFunc("/v1/mysql/getsome/{key1}/{key2}/{key3}", controller.MysqlGetSome(MySqlDb)).Methods("GET")  //mysql取得部分值
	router.HandleFunc("/v1/mysql/update", controller.MysqlUpdate(MySqlDb)).Methods("PUT")                         //mysql更新值
	router.HandleFunc("/v1/mysql/delete/{key1}/{key2}/{key3}", controller.MysqlDelete(MySqlDb)).Methods("DELETE") //mysql刪除值

	//mssql
	router.HandleFunc("/v1/mssql/addvalue", controller.MssqlAddValue(MsSqlDb)).Methods("POST")                    //mysql插入值
	router.HandleFunc("/v1/mssql/getall", controller.MssqlGetAll(MsSqlDb)).Methods("GET")                         //mssql取得所有值
	router.HandleFunc("/v1/mssql/getsome/{key1}/{key2}/{key3}", controller.MssqlGetSome(MsSqlDb)).Methods("GET")  //mssql取得部分值
	router.HandleFunc("/v1/mssql/update", controller.MssqlUpdate(MsSqlDb)).Methods("PUT")                         //mssql更新值
	router.HandleFunc("/v1/mssql/delete/{key1}/{key2}/{key3}", controller.MssqlDelete(MsSqlDb)).Methods("DELETE") //mysql刪除值

	//join table
	router.HandleFunc("/v1/join/getall", controller.JoinGetAll(MySqlDb, MsSqlDb)).Methods("GET")                        //mssql取得所有值
	router.HandleFunc("/v1/join/getsome/{key1}/{key2}/{key3}", controller.JoinGetSome(MySqlDb, MsSqlDb)).Methods("GET") //mssql取得部分值

	//test

	//安全性驗證
	//func (r *Router) Use(mc MiddlewareChain)
	//attach JWT auth middleware
	//router.Use(utils.JwtAuthentication)

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

/*
func test(w http.ResponseWriter, r *http.Request) {
	var (
		error models.Error
		//people     models.People
		peoples    []models.People
		conditions = make(map[string]interface{})
	)

	//處理query參數
	key1 := r.URL.Query()["key1"]
	if len(key1) > 0 {
		conditions["key1"] = key1[0]
	}
	key2 := r.URL.Query()["key2"]
	if len(key2) > 0 {
		conditions["key2"] = key2[0]
	}
	key3 := r.URL.Query()["key3"]
	if len(key3) > 0 {
		conditions["key3"] = key3[0]
	}
	number := r.URL.Query()["number"]
	if len(number) > 0 {
		conditions["number"] = number[0]
	}
	gender := r.URL.Query()["gender"]
	if len(gender) > 0 {
		conditions["gender"] = gender[0]
	}
	birth := r.URL.Query()["birth"]
	if len(birth) > 0 {
		conditions["birth"] = birth[0]
	}
	injury_degree := r.URL.Query()["injury_degree"]
	if len(injury_degree) > 0 {
		conditions["injury_degree"] = injury_degree[0]
	}
	injury_position := r.URL.Query()["injury_position"]
	if len(injury_position) > 0 {
		conditions["injury_position"] = injury_position[0]
	}
	protection := r.URL.Query()["protection"]
	if len(protection) > 0 {
		conditions["protection"] = protection[0]
	}
	phone := r.URL.Query()["phone"]
	if len(phone) > 0 {
		conditions["phone"] = phone[0]
	}
	person := r.URL.Query()["person"]
	if len(person) > 0 {
		conditions["person"] = person[0]
	}
	car := r.URL.Query()["car"]
	if len(car) > 0 {
		conditions["car"] = car[0]
	}
	action_status := r.URL.Query()["action_status"]
	if len(action_status) > 0 {
		conditions["action_status"] = action_status[0]
	}
	qualification := r.URL.Query()["qualification"]
	if len(qualification) > 0 {
		conditions["qualification"] = qualification[0]
	}
	license := r.URL.Query()["license"]
	if len(license) > 0 {
		conditions["license"] = license[0]
	}
	drinking := r.URL.Query()["drinking"]
	if len(drinking) > 0 {
		conditions["drinking"] = drinking[0]
	}
	hit := r.URL.Query()["hit"]
	if len(hit) > 0 {
		conditions["hit"] = hit[0]
	}

	//peoples, err := MySqlDb.MysqlQuerySomeData(peoples, conditions)
	err := MySqlDb.Where(conditions).Find(&peoples).Error
	if len(peoples) == 0 {
		error.Message = "The user does not exist!"
		utils.SendError(w, http.StatusBadRequest, error)
		return
	}
	if err != nil {
		error.Message = "Server(database) error!"
		utils.SendError(w, http.StatusInternalServerError, error)
		return
	}
	utils.SendSuccess(w, peoples)
}
*/
