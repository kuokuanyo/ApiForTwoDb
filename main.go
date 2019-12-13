package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/subosito/gotenv"
	"golang.org/x/crypto/bcrypt"
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

//user information
type User struct {
	ID       int
	Email    string
	Password string
}

//error
type Error struct {
	Message string
}

//驗證
type JWT struct {
	Token string
}

var MySqlDb *gorm.DB
var MsSqlDb *gorm.DB
var users []User
var user User
var people People
var peoples []People
var event Event
var events []Event

var err error

func init() {

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
}

func main() {
	//最後必須關閉
	defer MySqlDb.Close()
	defer MsSqlDb.Close()

	//create router
	//func NewRouter() *Router
	router := mux.NewRouter()

	//func (r *Router) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route
	//func (r *Router) Methods(methods ...string) *Route
	//mysql
	router.HandleFunc("/v1/mysql/signup", MysqlSignup).Methods("POST")     //mysql註冊
	router.HandleFunc("/v1/mysql/login", MysqlLogin).Methods("POST")       //mysql登入
	router.HandleFunc("/v1/mysql/addvalue", MysqlAddValue).Methods("POST") //mysql插入值
	router.HandleFunc("/v1/mysql/getall", MysqlGetAll).Methods("GET")      //mysql取得所有值
	router.HandleFunc("/v1/mysql/getsome", MysqlGetSome).Methods("GET")    //mysql取得部分值
	router.HandleFunc("/v1/mysql/update", MysqlUpdate).Methods("PUT")      //mysql更新值
	router.HandleFunc("/v1/mysql/delete", MysqlDelete).Methods("DELETE")   //mysql刪除值
	//mssql
	router.HandleFunc("/v1/mssql/signup", MssqlSignup).Methods("POST")     //mssql註冊
	router.HandleFunc("/v1/mssql/login", MssqlLogin).Methods("POST")       //mssql登入
	router.HandleFunc("/v1/mssql/addvalue", MssqlAddValue).Methods("POST") //mysql插入值
	router.HandleFunc("/v1/mssql/getall", MssqlGetAll).Methods("GET")      //mssql取得所有值
	router.HandleFunc("/v1/mssql/getsome", MssqlGetSome).Methods("GET")    //mssql取得部分值
	router.HandleFunc("/v1/mssql/update", MssqlUpdate).Methods("PUT")      //mssql更新值
	router.HandleFunc("/v1/mssql/delete", MssqlDelete).Methods("DELETE")   //mysql刪除值

	//func (r *Router) Use(mc MiddlewareChain)
	//attach JWT auth middleware
	router.Use(MysqlJwtAuthentication)
	router.Use(MssqlJwtAuthentication)

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

//delete value
//mysql
func MysqlDelete(w http.ResponseWriter, r *http.Request) {
	var error Error

	//decode
	json.NewDecoder(r.Body).Decode(&people)

	err = MysqlDeleteData(MySqlDb, people, map[string]interface{}{"key1": people.Key1})
	if err != nil {
		error.Message = "Server(database) error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}

	SendSuccess(w, "Success!")
}

//mssql
func MssqlDelete(w http.ResponseWriter, r *http.Request) {
	var error Error

	//decode
	json.NewDecoder(r.Body).Decode(&event)

	err = MssqlDeleteData(MsSqlDb, event, map[string]interface{}{"key1": event.Key1})
	if err != nil {
		error.Message = "Server(database) error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}

	SendSuccess(w, "Success!")
}

//update value
//mysql
func MysqlUpdate(w http.ResponseWriter, r *http.Request) {
	var error Error

	//decode
	json.NewDecoder(r.Body).Decode(&people)

	err = MysqlUpdateData(MySqlDb, people,
		map[string]interface{}{"key1": people.Key1},
		map[string]interface{}{"birth": people.Birth})
	if err != nil {
		error.Message = "Server(database) error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}

	SendSuccess(w, "Success!")
}

//mssql
func MssqlUpdate(w http.ResponseWriter, r *http.Request) {
	var error Error

	//decode
	json.NewDecoder(r.Body).Decode(&event)

	err = MssqlUpdateData(MsSqlDb, event,
		map[string]interface{}{"key1": event.Key1},
		map[string]interface{}{"death": event.Death})
	if err != nil {
		error.Message = "Server(database) error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}

	SendSuccess(w, "Success!")
}

//get some data
//mysql
func MysqlGetSome(w http.ResponseWriter, r *http.Request) {
	var error Error

	//decode
	json.NewDecoder(r.Body).Decode(&people)

	peoples, err = MysqlQuerySomeData(MySqlDb, peoples, map[string]interface{}{"key1": people.Key1})
	if err != nil {
		//找不到資料
		if err.Error() == "record not found" {
			error.Message = "The user does not exist!"
			SendError(w, http.StatusBadRequest, error)
			return
		} else {
			error.Message = "Server(database) error!"
			SendError(w, http.StatusInternalServerError, error)
			return
		}
	}
	SendSuccess(w, peoples)
}

//mssql
func MssqlGetSome(w http.ResponseWriter, r *http.Request) {
	var error Error

	//decode
	json.NewDecoder(r.Body).Decode(&event)

	events, err = MssqlQuerySomeData(MsSqlDb, events, map[string]interface{}{"key1": event.Key1})
	if err != nil {
		//找不到資料
		if err.Error() == "record not found" {
			error.Message = "The user does not exist!"
			SendError(w, http.StatusBadRequest, error)
			return
		} else {
			error.Message = "Server(database) error!"
			SendError(w, http.StatusInternalServerError, error)
			return
		}
	}
	SendSuccess(w, events)
}

//get all data
//mysql
func MysqlGetAll(w http.ResponseWriter, r *http.Request) {
	var error Error
	peoples, err := MysqlQueryAllData(MySqlDb, peoples)
	if err != nil {
		error.Message = "Server(database) error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}
	SendSuccess(w, peoples)
}

//mssql
func MssqlGetAll(w http.ResponseWriter, r *http.Request) {
	var error Error
	events, err := MssqlQueryAllData(MsSqlDb, events)
	if err != nil {
		error.Message = "Server(database) error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}
	SendSuccess(w, events)
}

//add value
//mysql
func MysqlAddValue(w http.ResponseWriter, r *http.Request) {
	//decode
	json.NewDecoder(r.Body).Decode(&people)

	var error Error
	err := MysqlInsertValue(MySqlDb, people)
	if err != nil {
		error.Message = "Server(database) error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}
	SendSuccess(w, people)
}

//mssql
func MssqlAddValue(w http.ResponseWriter, r *http.Request) {
	//decode
	json.NewDecoder(r.Body).Decode(&event)

	var error Error
	err := MssqlInsertValue(MsSqlDb, event)
	if err != nil {
		error.Message = "Server(database) error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}
	SendSuccess(w, event)
}

//jwt驗證
//mysql
func MysqlJwtAuthentication(next http.Handler) http.Handler {
	//匿名函式
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errorObject Error
		//從header取得token
		authHeader := r.Header.Get("Authorization")
		//不須驗證的路徑
		paths := []string{"/v1/mysql/signup", "/v1/mysql/login",
			"/v1/mssql/signup", "/v1/mssql/login",
			"/v1/mssql/addvalue", "/v1/mssql/getall",
			"/v1/mssql/getsome", "/v1/mssql/update",
			"/v1/mssql/delete"}
		//current request path
		requestPath := r.URL.Path

		//不須驗證的路徑，直接執行
		for _, path := range paths {
			if path == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		//if authHeader is empty
		if authHeader == "" {
			errorObject.Message = "Missing auth token!"
			SendError(w, http.StatusForbidden, errorObject)
			return
		}

		//split
		splitted := strings.Split(authHeader, " ")

		//if length is not 2
		if len(splitted) != 2 {
			errorObject.Message = "Invaild token!"
			SendError(w, http.StatusUnauthorized, errorObject)
			return
		}

		//取得第二個位置的值
		authHeader = splitted[1]

		//jwt驗證並解析
		//func Parse(tokenString string, keyFunc Keyfunc) (*Token, error)
		//type Keyfunc func(*Token) (interface{}, error)
		/*
			type Token struct {
			Raw       string                 // The raw token.  Populated when you Parse a token
			Method    SigningMethod          // The signing method used or to be used
			Header    map[string]interface{} // The first segment of the token
			Claims    Claims                 // The second segment of the token
			Signature string                 // The third segment of the token.  Populated when you Parse a token
			Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
			}
		*/
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error!")
			}
			return []byte(os.Getenv("mysql_token_password")), nil
		})
		if err != nil {
			errorObject.Message = err.Error()
			SendError(w, http.StatusUnauthorized, errorObject)
			return
		}

		//if token is vaild, return true
		if token.Valid {
			//通驗驗證
			next.ServeHTTP(w, r)
			return
		} else {
			errorObject.Message = err.Error()
			SendError(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}

//mssql
func MssqlJwtAuthentication(next http.Handler) http.Handler {
	//匿名函式
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errorObject Error
		//從header取得token
		authHeader := r.Header.Get("Authorization")
		//不須驗證的路徑
		paths := []string{"/v1/mysql/signup", "/v1/mysql/login",
			"/v1/mssql/signup", "/v1/mssql/login",
			"/v1/mysql/addvalue", "/v1/mysql/getall",
			"/v1/mysql/getsome", "/v1/mysql/update",
			"/v1/mysql/delete"}
		//current request path
		requestPath := r.URL.Path

		//不須驗證的路徑，直接執行
		for _, path := range paths {
			if path == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		//if authHeader is empty
		if authHeader == "" {
			errorObject.Message = "Missing auth token!"
			SendError(w, http.StatusForbidden, errorObject)
			return
		}

		//split
		splitted := strings.Split(authHeader, " ")

		//if length is not 2
		if len(splitted) != 2 {
			errorObject.Message = "Invaild token!"
			SendError(w, http.StatusUnauthorized, errorObject)
			return
		}

		//取得第二個位置的值
		authHeader = splitted[1]

		//jwt驗證並解析
		//func Parse(tokenString string, keyFunc Keyfunc) (*Token, error)
		//type Keyfunc func(*Token) (interface{}, error)
		/*
			type Token struct {
			Raw       string                 // The raw token.  Populated when you Parse a token
			Method    SigningMethod          // The signing method used or to be used
			Header    map[string]interface{} // The first segment of the token
			Claims    Claims                 // The second segment of the token
			Signature string                 // The third segment of the token.  Populated when you Parse a token
			Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
			}
		*/
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error!")
			}
			return []byte(os.Getenv("mssql_token_password")), nil
		})
		if err != nil {
			errorObject.Message = err.Error()
			SendError(w, http.StatusUnauthorized, errorObject)
			return
		}

		//if token is vaild, return true
		if token.Valid {
			//通驗驗證
			next.ServeHTTP(w, r)
			return
		} else {
			errorObject.Message = err.Error()
			SendError(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}

//login
//mysql
func MysqlLogin(w http.ResponseWriter, r *http.Request) {
	var jwt JWT
	var error Error

	//decode(pointer)
	json.NewDecoder(r.Body).Decode(&user)

	//email、password不能為空
	if user.Email == "" {
		error.Message = "E-mail is not empty!"
		SendError(w, http.StatusBadRequest, error)
		return
	}
	if user.Password == "" {
		error.Message = "Password is not empty!"
		SendError(w, http.StatusBadRequest, error)
		return
	}

	//login's password
	password := user.Password

	user, err = MysqlReadUser(MySqlDb, user, "email =?", user.Email)
	if err != nil {
		//找不到資料
		if err.Error() == "record not found" {
			error.Message = "The user does not exist!"
			SendError(w, http.StatusBadRequest, error)
			return
		} else {
			error.Message = "Server(database) error!"
			SendError(w, http.StatusInternalServerError, error)
			return
		}
	}
	//database password
	hashedpassword := user.Password

	//check password
	if password != hashedpassword {
		//compare the password
		//func CompareHashAndPassword(hashedPassword, password []byte) error
		err = bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
		if err != nil {
			error.Message = "Invaild Password!"
			SendError(w, http.StatusUnauthorized, error)
			return
		}
	}

	//create token
	token, err := MysqlGenerateToken(user)
	if err != nil {
		error.Message = "Generate Token error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}

	w.WriteHeader(http.StatusOK)
	jwt.Token = token
	SendSuccess(w, jwt)
}

//mssql
func MssqlLogin(w http.ResponseWriter, r *http.Request) {
	var jwt JWT
	var error Error

	//decode(pointer)
	json.NewDecoder(r.Body).Decode(&user)

	//email、password不能為空
	if user.Email == "" {
		error.Message = "E-mail is not empty!"
		SendError(w, http.StatusBadRequest, error)
		return
	}
	if user.Password == "" {
		error.Message = "Password is not empty!"
		SendError(w, http.StatusBadRequest, error)
		return
	}

	//login's password
	password := user.Password

	user, err = MssqlReadUser(MsSqlDb, user, "email =?", user.Email)
	if err != nil {
		//找不到資料
		if err.Error() == "record not found" {
			error.Message = "The user does not exist!"
			SendError(w, http.StatusBadRequest, error)
			return
		} else {
			error.Message = "Server(database) error!"
			SendError(w, http.StatusInternalServerError, error)
			return
		}
	}
	//database password
	hashedpassword := user.Password

	//check password
	if password != hashedpassword {
		//compare the password
		//func CompareHashAndPassword(hashedPassword, password []byte) error
		err = bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
		if err != nil {
			error.Message = "Invaild Password!"
			SendError(w, http.StatusUnauthorized, error)
			return
		}
	}

	//create token
	token, err := MssqlGenerateToken(user)
	if err != nil {
		error.Message = "Generate Token error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}

	w.WriteHeader(http.StatusOK)
	jwt.Token = token
	SendSuccess(w, jwt)
}

//json-web-token
//mysql
func MysqlGenerateToken(user User) (string, error) {
	s := os.Getenv("mysql_token_password")

	//a jwt
	//header.payload.s
	//func NewWithClaims(method SigningMethod, claims Claims) *Token
	claims := jwt.MapClaims{
		"email":    user.Email,
		"password": user.Password,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), //增加過期時間
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//生成簽名字串(s)
	//func (t *Token) SignedString(key interface{}) (string, error)
	tokenString, err := token.SignedString([]byte(s))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//mssql
func MssqlGenerateToken(user User) (string, error) {
	s := os.Getenv("mssql_token_password")

	//a jwt
	//header.payload.s
	//func NewWithClaims(method SigningMethod, claims Claims) *Token
	claims := jwt.MapClaims{
		"email":    user.Email,
		"password": user.Password,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), //增加過期時間
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//生成簽名字串(s)
	//func (t *Token) SignedString(key interface{}) (string, error)
	tokenString, err := token.SignedString([]byte(s))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//read user
//mysql
func MysqlReadUser(db *gorm.DB, user User, condition string, args ...interface{}) (User, error) {
	err := db.Where(condition, args...).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

//mysql
func MssqlReadUser(db *gorm.DB, user User, condition string, args ...interface{}) (User, error) {
	err := db.Where(condition, args...).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

//sign up
//mysql
func MysqlSignup(w http.ResponseWriter, r *http.Request) {
	var error Error

	//decode(須指標)
	//寫入user
	json.NewDecoder(r.Body).Decode(&user)

	//email、password不能是空的
	if user.Email == "" {
		error.Message = "E-mail is not empty!"
		SendError(w, http.StatusBadRequest, error)
		return
	}
	if user.Password == "" {
		error.Message = "Password is not empty!"
		SendError(w, http.StatusBadRequest, error)
		return
	}
	//email must have '@'
	if !strings.Contains(user.Email, "@") {
		error.Message = "Email address is error!"
		SendError(w, http.StatusBadRequest, error)
		return
	}
	//password length needs to be greater than six
	if len(user.Password) < 6 {
		error.Message = "Password length is not enough(6 char)!"
		SendError(w, http.StatusBadRequest, error)
		return
	}

	//檢查信箱是否已經被使用過
	users, err = MysqlQueryAllUser(MySqlDb, users)
	if err != nil {
		error.Message = "Server(database) error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}

	//check
	for _, user_db := range users {
		if user.Email == user_db.Email {
			error.Message = "E-mail already taken!"
			SendError(w, http.StatusForbidden, error)
			return
		}
	}

	//密碼加密
	//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		error.Message = "Server error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}
	//convert type
	//assign hash to password
	user.Password = string(hash)

	//insert the new user
	err = MysqlInsertUser(MySqlDb, user)
	if err != nil {
		error.Message = "Server(database) error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}

	//加入資料庫後，密碼為空白
	user.Password = ""
	SendSuccess(w, user)
	return
}

//mssql
func MssqlSignup(w http.ResponseWriter, r *http.Request) {
	var error Error

	//decode(須指標)
	//寫入user
	json.NewDecoder(r.Body).Decode(&user)

	//email、password不能是空的
	if user.Email == "" {
		error.Message = "E-mail is not empty!"
		SendError(w, http.StatusBadRequest, error)
		return
	}
	if user.Password == "" {
		error.Message = "Password is not empty!"
		SendError(w, http.StatusBadRequest, error)
		return
	}
	//email must have '@'
	if !strings.Contains(user.Email, "@") {
		error.Message = "Email address is error!"
		SendError(w, http.StatusBadRequest, error)
		return
	}
	//password length needs to be greater than six
	if len(user.Password) < 6 {
		error.Message = "Password length is not enough(6 char)!"
		SendError(w, http.StatusBadRequest, error)
		return
	}

	//檢查信箱是否已經被使用過
	users, err = MysqlQueryAllUser(MsSqlDb, users)
	if err != nil {
		error.Message = "Server(database) error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}

	//check
	for _, user_db := range users {
		if user.Email == user_db.Email {
			error.Message = "E-mail already taken!"
			SendError(w, http.StatusForbidden, error)
			return
		}
	}

	//密碼加密
	//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		error.Message = "Server error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}
	//convert type
	//assign hash to password
	user.Password = string(hash)

	//insert the new user
	err = MysqlInsertUser(MsSqlDb, user)
	if err != nil {
		error.Message = "Server(database) error!"
		SendError(w, http.StatusInternalServerError, error)
		return
	}

	//加入資料庫後，密碼為空白
	user.Password = ""
	SendSuccess(w, user)
	return
}

//insert the new user
//mysql
func MysqlInsertUser(db *gorm.DB, user User) error {
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

//mssql
func MssqlInsertUser(db *gorm.DB, user User) error {
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

//get all user data
//mysql
func MysqlQueryAllUser(db *gorm.DB, users []User) ([]User, error) {
	err := db.Find(&users).Error
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

//mssql
func MssqlQueryAllUser(db *gorm.DB, users []User) ([]User, error) {
	err := db.Find(&users).Error
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

//response error
func SendError(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	//encode
	json.NewEncoder(w).Encode(error)
}

//response success
func SendSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	//encode
	json.NewEncoder(w).Encode(data)
}

//create table(if no table)
//mysql
func MysqlCreateTable(db *gorm.DB, people People) error {
	//func (s *DB) CreateTable(models ...interface{}) *DB
	err := db.CreateTable(&people).Error
	if err != nil {
		return err
	}
	return nil
}

//mssql
func MssqlCreateTable(db *gorm.DB, event Event) error {
	//func (s *DB) CreateTable(models ...interface{}) *DB
	err := db.CreateTable(&event).Error
	if err != nil {
		return err
	}
	return nil
}

//insert value
//mysql
func MysqlInsertValue(db *gorm.DB, people People) error {
	err := db.Create(&people).Error
	if err != nil {
		return err
	}
	return nil
}

//mssql
func MssqlInsertValue(db *gorm.DB, event Event) error {
	err := db.Create(&event).Error
	if err != nil {
		return err
	}
	return nil
}

//find all datas
//mysql
func MysqlQueryAllData(db *gorm.DB, peoples []People) ([]People, error) {
	err := db.Find(&peoples).Error
	if err != nil {
		return []People{}, err
	}
	return peoples, nil
}

//mssql
func MssqlQueryAllData(db *gorm.DB, events []Event) ([]Event, error) {
	err := db.Find(&events).Error
	if err != nil {
		return []Event{}, err
	}
	return events, nil
}

//find datas by some condition
//mysql
func MysqlQuerySomeData(db *gorm.DB, peoples []People, condition map[string]interface{}) ([]People, error) {
	err := db.Where(condition).Find(&peoples).Error
	if err != nil {
		return []People{}, err
	}
	return peoples, nil
}

//mssql
func MssqlQuerySomeData(db *gorm.DB, events []Event, condition map[string]interface{}, args ...interface{}) ([]Event, error) {
	err := db.Where(condition).Find(&events).Error
	if err != nil {
		return []Event{}, err
	}
	return events, nil
}

//find one data
//mysql
func MysqlQueryOneData(db *gorm.DB, people People, condition string, args ...interface{}) (People, error) {
	err := db.Where(condition, args...).First(&people).Error
	if err != nil {
		return People{}, err
	}
	return people, nil
}

//mssql
func MssqlQueryOneData(db *gorm.DB, event Event, order string, condition string, args ...interface{}) (Event, error) {
	err := db.Order(order).Where(condition, args...).First(&event).Error
	fmt.Println(event)
	if err != nil {
		return Event{}, err
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
	err := db.Model(&event).Where(condition).Update(update).Error
	if err != nil {
		return err
	}
	return nil
}

//delete data
//mysql
func MysqlDeleteData(db *gorm.DB, people People, condition map[string]interface{}) error {
	err := db.Where(condition).Delete(&people).Error
	if err != nil {
		return err
	}
	return nil
}

//mssql
func MssqlDeleteData(db *gorm.DB, event Event, condition map[string]interface{}) error {
	err := db.Where(condition).Delete(&event).Error
	if err != nil {
		return err
	}
	return nil
}
