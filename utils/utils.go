package utils

import (
	"ApiForTwoDb/driver"
	models "ApiForTwoDb/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//response error
func SendError(w http.ResponseWriter, status int, error models.Error) {
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

//jwt驗證
//mysql
func MysqlJwtAuthentication(next http.Handler) http.Handler {
	//匿名函式
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errorObject models.Error
		//從header取得token
		authHeader := r.Header.Get("Authorization")
		//不須驗證的路徑
		paths := []string{"/v1/mysql/signup", "/v1/mysql/login",
			"/v1/mssql/signup", "/v1/mssql/login",
			"/v1/mssql/addvalue", "/v1/mssql/getall",
			"/v1/mssql/getsome", "/v1/mssql/update",
			"/v1/mssql/delete", `/v1/join/getall`,
			"/v1/join/getsome"}
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
		var errorObject models.Error
		//從header取得token
		authHeader := r.Header.Get("Authorization")
		//不須驗證的路徑
		paths := []string{"/v1/mysql/signup", "/v1/mysql/login",
			"/v1/mssql/signup", "/v1/mssql/login",
			"/v1/mysql/addvalue", "/v1/mysql/getall",
			"/v1/mysql/getsome", "/v1/mysql/update",
			"/v1/mysql/delete", `/v1/join/getall`,
			"/v1/join/getsome"}
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

//json-web-token
//mysql
func MysqlGenerateToken(user models.User) (string, error) {
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
func MssqlGenerateToken(user models.User) (string, error) {
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

//建立關聯表
func CreateJoinTable(MySqlDb *driver.MySqlDb, MsSqlDb *driver.MsSqlDb, peoples []models.People, events []models.Event) {
	if !MySqlDb.HasTable(&models.JoinTable{}) {
		MySqlDb.CreateTable(&models.JoinTable{})
	}

	peoples, err := MySqlDb.MysqlQueryAllData(peoples)
	if err != nil {
		log.Fatal(err)
	}
	events, err = MsSqlDb.MssqlQueryAllData(events)
	if err != nil {
		log.Fatal(err)
	}

	for _, people := range peoples {
		for _, event := range events {
			var jointable models.JoinTable
			//關聯表欄位
			if people.Key1 == event.Key1 && people.Key2 == event.Key2 && people.Key3 == event.Key3 {
				jointable.Key1 = people.Key1
				jointable.Key2 = people.Key2
				jointable.Key3 = people.Key3
				jointable.Number = people.Number
				jointable.Gender = people.Gender
				jointable.Birth = people.Birth
				jointable.Injury_degree = people.Injury_degree
				jointable.Injury_position = people.Injury_position
				jointable.Protection = people.Protection
				jointable.Phone = people.Phone
				jointable.Person = people.Person
				jointable.Car = people.Car
				jointable.Action_status = people.Action_status
				jointable.Qualification = people.Qualification
				jointable.License = people.License
				jointable.Drinking = people.Drinking
				jointable.Hit = people.Hit
				jointable.City = event.City
				jointable.Position = event.Position
				jointable.Lane = event.Lane
				jointable.Death = event.Death
				jointable.Injured = event.Injured
				jointable.Death_exceed = event.Death_exceed
				jointable.Weather = event.Weather
				jointable.Light = event.Light
				jointable.Time_year = event.Time_year
				jointable.Time_month = event.Time_month
				jointable.Accident_chinese = event.Accident_chinese
				jointable.Anecdote_chinese = event.Anecdote_chinese
				//插入數值
				err := MySqlDb.Create(&jointable).Error
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
