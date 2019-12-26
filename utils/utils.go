package utils

import (
	models "ApiForTwoDb/model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//SendError response error
func SendError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	//encode
	json.NewEncoder(w).Encode(error)
}

//SendSuccess response success
func SendSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	//encode
	json.NewEncoder(w).Encode(data)
}

//JwtAuthentication jwt驗證
func JwtAuthentication(next http.Handler) http.Handler {
	//匿名函式
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errorObject models.Error
		//從header取得token
		authHeader := r.Header.Get("Authorization")
		//不須驗證的路徑
		paths := []string{"/v1/signup/mysql", "/v1/login/mysql",
			"/v1/signup/mssql", "/v1/login/mssql",
		}

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
				return nil, fmt.Errorf("there was an error")
			}
			return []byte(os.Getenv("token_password")), nil
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
		} else {
			errorObject.Message = err.Error()
			SendError(w, http.StatusUnauthorized, errorObject)
		}
	})
}

//GenerateToken generate json-web-token
func GenerateToken(user models.User) (string, error) {
	s := os.Getenv("token_password")

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

//StructToMap : struct to map
func StructToMap(obj interface{}) map[string]interface{} {
	//func TypeOf(i interface{}) Type
	//回傳類型
	t := reflect.TypeOf(obj)
	//func ValueOf(i interface{}) Value
	//回傳數值
	v := reflect.ValueOf(obj)

	//map
	var data = make(map[string]interface{})

	//t.NumField()返回struct數量
	for i := 0; i < t.NumField(); i++ {
		//Field返回第i個位置
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
