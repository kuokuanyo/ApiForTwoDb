//@title Restful API(connect two databases)
//@version 1.0.0
//@description Define an API
//@Schemes http
//@host localhost:8080
//@BasePath /v1
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ApiForTwoDb/controllers"
	"ApiForTwoDb/driver"
	models "ApiForTwoDb/model"

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
	router.HandleFunc("/v1/getall/{sql}", controller.GetAll(MySqlDb, MsSqlDb)).Methods("GET")      //取得所有值
	router.HandleFunc("/v1/getsome/{sql}", controller.GetSome(MySqlDb, MsSqlDb)).Methods("GET")    //取得部分值
	router.HandleFunc("/v1/update/{sql}", controller.Update(MySqlDb, MsSqlDb)).Methods("PUT")      //更新值
	router.HandleFunc("/v1/delete/{sql}", controller.Delete(MySqlDb, MsSqlDb)).Methods("DELETE")   //刪除值

	//join table
	router.HandleFunc("/v1/join/getall", test).Methods("GET") //mssql取得所有值
	//router.HandleFunc("/v1/join/getsome/{key1}/{key2}/{key3}", controller.JoinGetSome(MySqlDb, MsSqlDb)).Methods("GET") //mssql取得部分值

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

func test(w http.ResponseWriter, r *http.Request) {
	var (
		//error   models.Error
		peoples []models.People
		events  []models.Event
		//jointables []models.JoinTable
		//userRepo repository.UserRepository
	)

	//挑選的欄位
	//key為合併的欄位
	//MysqlColName := []string{"key1", "key2", "key3"}
	MssqlColName := []string{"key1, key2, key3"}

	//更新關聯資料庫
	//read all datas from peoples
	//MySqlDb.Table("peoples").Select(MssqlColName).Rows()

	rows, err := MySqlDb.Model(&models.People{}).Select(MssqlColName).Rows()
	for rows.Next() {
		var people models.People
		rows.Scan(&people.Key1, &people.Key2, &people.Key3)
		peoples = append(peoples, people)
	}

	//read all datas from events
	events, err = MsSqlDb.MssqlQueryAllDataBySomeCol(events, MssqlColName)
	if err != nil {
		return
	}
	fmt.Println(peoples)
	/*
		//取得所有資料
		jointables, err := userRepo.QueryAllJoinData(MySqlDb, jointables)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		utils.SendSuccess(w, jointables)
	*/
}

/*
	//處理query參數
	number := r.URL.Query()["number"]
	if len(number) > 0 {
		if number[0] == "number" {
			MysqlColName = append(MysqlColName, number[0])
		}
	}
	gender := r.URL.Query()["gender"]
	if len(gender) > 0 {
		if gender[0] == "gender" {
			MysqlColName = append(MysqlColName, gender[0])
		}
	}
	birth := r.URL.Query()["birth"]
	if len(birth) > 0 {
		if birth[0] == "birth" {
			MysqlColName = append(MysqlColName, birth[0])
		}
	}
	injury_degree := r.URL.Query()["injury_degree"]
	if len(injury_degree) > 0 {
		if injury_degree[0] == "injury_degree" {
			MysqlColName = append(MysqlColName, injury_degree[0])
		}
	}
	injury_position := r.URL.Query()["injury_position"]
	if len(injury_position) > 0 {
		if injury_position[0] == "injury_position" {
			MysqlColName = append(MysqlColName, injury_position[0])
		}
	}
	protection := r.URL.Query()["protection"]
	if len(protection) > 0 {
		if protection[0] == "protection" {
			MysqlColName = append(MysqlColName, protection[0])
		}
	}
	phone := r.URL.Query()["phone"]
	if len(phone) > 0 {
		if phone[0] == "phone" {
			MysqlColName = append(MysqlColName, phone[0])
		}
	}
	person := r.URL.Query()["person"]
	if len(person) > 0 {
		if person[0] == "person" {
			MysqlColName = append(MysqlColName, person[0])
		}
	}
	car := r.URL.Query()["car"]
	if len(car) > 0 {
		if car[0] == "car" {
			MysqlColName = append(MysqlColName, car[0])
		}
	}
	action_status := r.URL.Query()["action_status"]
	if len(action_status) > 0 {
		if action_status[0] == "action_status" {
			MysqlColName = append(MysqlColName, action_status[0])
		}
	}
	qualification := r.URL.Query()["qualification"]
	if len(qualification) > 0 {
		if qualification[0] == "qualification" {
			MysqlColName = append(MysqlColName, qualification[0])
		}
	}
	license := r.URL.Query()["license"]
	if len(license) > 0 {
		if license[0] == "license" {
			MysqlColName = append(MysqlColName, license[0])
		}
	}
	drinking := r.URL.Query()["drinking"]
	if len(drinking) > 0 {
		if drinking[0] == "drinking" {
			MysqlColName = append(MysqlColName, drinking[0])
		}
	}
	hit := r.URL.Query()["hit"]
	if len(hit) > 0 {
		if hit[0] == "hit" {
			MysqlColName = append(MysqlColName, hit[0])
		}
	}
	city := r.URL.Query()["city"]
	if len(city) > 0 {
		if city[0] == "city" {
			MssqlColName = append(MssqlColName, city[0])
		}
	}
	position := r.URL.Query()["position"]
	if len(position) > 0 {
		if position[0] == "position" {
			MssqlColName = append(MssqlColName, position[0])
		}
	}
	lane := r.URL.Query()["lane"]
	if len(lane) > 0 {
		if lane[0] == "lane" {
			MssqlColName = append(MssqlColName, lane[0])
		}
	}
	death := r.URL.Query()["death"]
	if len(death) > 0 {
		if death[0] == "death" {
			MssqlColName = append(MssqlColName, death[0])
		}
	}
	injured := r.URL.Query()["injured"]
	if len(injured) > 0 {
		if injured[0] == "injured" {
			MssqlColName = append(MssqlColName, injured[0])
		}
	}
	death_exceed := r.URL.Query()["death_exceed"]
	if len(death_exceed) > 0 {
		if death_exceed[0] == "death_exceed" {
			MssqlColName = append(MssqlColName, death_exceed[0])
		}
	}
	weather := r.URL.Query()["weather"]
	if len(weather) > 0 {
		if weather[0] == "weather" {
			MssqlColName = append(MssqlColName, weather[0])
		}
	}
	light := r.URL.Query()["light"]
	if len(light) > 0 {
		if light[0] == "light" {
			MssqlColName = append(MssqlColName, light[0])
		}
	}
	time_year := r.URL.Query()["time_year"]
	if len(time_year) > 0 {
		if time_year[0] == "time_year" {
			MssqlColName = append(MssqlColName, time_year[0])
		}
	}
	time_month := r.URL.Query()["time_month"]
	if len(time_month) > 0 {
		if time_month[0] == "time_month" {
			MssqlColName = append(MssqlColName, time_month[0])
		}
	}
	accident_chinese := r.URL.Query()["accident_chinese"]
	if len(accident_chinese) > 0 {
		if accident_chinese[0] == "accident_chinese" {
			MssqlColName = append(MssqlColName, accident_chinese[0])
		}
	}
	anecdote_chinese := r.URL.Query()["anecdote_chinese"]
	if len(anecdote_chinese) > 0 {
		if anecdote_chinese[0] == "anecdote_chinese" {
			MssqlColName = append(MssqlColName, anecdote_chinese[0])
		}
	}
*/
