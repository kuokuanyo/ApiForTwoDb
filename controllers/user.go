package controllers

import (
	"ApiForTwoDb/driver"
	"ApiForTwoDb/repository"
	"ApiForTwoDb/utils"

	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	models "ApiForTwoDb/model"
)

//Controller struct
type Controller struct{}

//Signup create a new account
//@Summary create a new account
//@Tags User
//@Description 註冊
//@Accept json
//@Produce json
//@Param sql path string true "資料庫引擎"
//@Param information body models.user true "個人資料"
//@Success 200 {object} models.user "Successfully"
//@Failure 400 {object} models.Error "Bad Request"
//@Failure 403 {object} models.Error "Forbidden"
//@Failure 500 {object} models.Error "Internal Server Error"
//@Router /signup/{sql} [post]
func (c Controller) Signup(MySQLDb *driver.MySQLDb, MsSQLDb *driver.MsSQLDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			error    models.Error
			user     models.User
			users    []models.User
			userRepo repository.UserRepository
		)

		//decode(須指標)
		json.NewDecoder(r.Body).Decode(&user)

		//email、password不能是空的
		if user.Email == "" {
			error.Message = "E-mail is not empty!"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		if user.Password == "" {
			error.Message = "Password is not empty!"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		//email must have '@'
		if !strings.Contains(user.Email, "@") {
			error.Message = "Email address is error!"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		//password length needs to be greater than six
		if len(user.Password) < 6 {
			error.Message = "Password length is not enough(6 char)!"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		//印出url參數
		params := mux.Vars(r)

		switch strings.ToLower(params["sql"]) {
		case "mysql":
			//檢查信箱是否已經被使用過
			users, err := userRepo.MysqlCheckSignup(MySQLDb, users)
			if err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}

			//check
			for _, userdb := range users {
				if user.Email == userdb.Email {
					error.Message = "E-mail already taken!"
					utils.SendError(w, http.StatusForbidden, error)
					return
				}
			}

			//密碼加密
			//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
			if err != nil {
				error.Message = "Server error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			//convert type
			//assign hash to password
			user.Password = string(hash)

			//insert the new user
			if err = userRepo.MysqlInsertUser(MySQLDb, user); err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		case "mssql":
			//檢查信箱是否已經被使用過
			users, err := userRepo.MssqlCheckSignup(MsSQLDb, users)
			if err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}

			//check
			for _, userdb := range users {
				if user.Email == userdb.Email {
					error.Message = "E-mail already taken!"
					utils.SendError(w, http.StatusForbidden, error)
					return
				}
			}

			//密碼加密
			//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
			if err != nil {
				error.Message = "Server error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
			//convert type
			//assign hash to password
			user.Password = string(hash)

			//insert the new user
			if err = userRepo.MssqlInsertUser(MsSQLDb, user); err != nil {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}
		//加入資料庫後，密碼為空白
		user.Password = ""
		utils.SendSuccess(w, user)
		return
	}
}

//Login user login
//@Summary login
//@Tags User
//@Description 登入
//@Accept json
//@Produce json
//@Param sql path string true "資料庫引擎"
//@Param information body models.user true "個人資料"
//@Success 200 {object} models.JWT "Successfully"
//@Failure 400 {object} models.Error "Bad Request"
//@Failure 401 {object} models.Error "Unauthorized"
//@Failure 500 {object} models.Error "Internal Server Error"
//@Router /login/{sql} [post]
func (c Controller) Login(MySQLDb *driver.MySQLDb, MsSQLDb *driver.MsSQLDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			hashedpassword string
			jwt            models.JWT
			error          models.Error
			user           models.User
			userRepo       repository.UserRepository
		)

		//decode(須指標)
		json.NewDecoder(r.Body).Decode(&user)

		//email、password不能是空的
		if user.Email == "" {
			error.Message = "E-mail is not empty!"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		if user.Password == "" {
			error.Message = "Password is not empty!"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		//email must have '@'
		if !strings.Contains(user.Email, "@") {
			error.Message = "Email address is error!"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		//password length needs to be greater than six
		if len(user.Password) < 6 {
			error.Message = "Password length is not enough(6 char)!"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		//login's password
		password := user.Password

		//印出url參數
		params := mux.Vars(r)

		switch strings.ToLower(params["sql"]) {
		case "mysql":
			user, err := userRepo.MysqlCheckLogin(MySQLDb, user)
			if err != nil {
				//找不到資料
				if err.Error() == "record not found" {
					error.Message = "The user does not exist!"
					utils.SendError(w, http.StatusBadRequest, error)
				} else {
					error.Message = "Server(database) error!"
					utils.SendError(w, http.StatusInternalServerError, error)
				}
			}
			//database password
			hashedpassword = user.Password

		case "mssql":
			user, err := userRepo.MssqlCheckLogin(MsSQLDb, user)
			if err != nil {
				//找不到資料
				if err.Error() == "record not found" {
					error.Message = "The user does not exist!"
					utils.SendError(w, http.StatusBadRequest, error)
				} else {
					error.Message = "Server(database) error!"
					utils.SendError(w, http.StatusInternalServerError, error)
				}
			}
			//database password
			hashedpassword = user.Password
		}

		//check password
		if password != hashedpassword {
			//compare the password
			//func CompareHashAndPassword(hashedPassword, password []byte) error
			if err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password)); err != nil {
				error.Message = "Invaild Password!"
				utils.SendError(w, http.StatusUnauthorized, error)
				return
			}
		}

		//create token
		token, err := utils.GenerateToken(user)
		if err != nil {
			error.Message = "Generate Token error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.WriteHeader(http.StatusOK)
		jwt.Token = token
		utils.SendSuccess(w, jwt)
		return
	}
}
