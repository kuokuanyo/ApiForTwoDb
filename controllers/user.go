package controllers

import (
	"ApiForTwoDb/driver"
	"ApiForTwoDb/repository"
	"ApiForTwoDb/utils"

	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	models "ApiForTwoDb/model"
)

type Controller struct{}

//@Summary create a new account
//@Tags Mysql_User
//@Description 註冊
//@Accept json
//@Produce json
//@Param information body models.user true "個人資料"
//@Success 200 {object} models.user "Successfully sign up!"
//@Failure 400 {object} models.Error "email or password error!"
//@Failure 403 {object} models.Error "E-mail already taken!"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /mysql/signup [post]
func (c Controller) MysqlSignup(MySqlDb *driver.MySqlDb) http.HandlerFunc {
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

		//檢查信箱是否已經被使用過
		users, err := userRepo.MysqlCheckSignup(MySqlDb, users)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		//check
		for _, user_db := range users {
			if user.Email == user_db.Email {
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
		if err = userRepo.MysqlInsertUser(MySqlDb, user); err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		//加入資料庫後，密碼為空白
		user.Password = ""
		utils.SendSuccess(w, user)
		return
	}
}

//@Summary create a new account
//@Tags Mssql_User
//@Description 註冊
//@Accept json
//@Produce json
//@Param information body models.user true "個人資料"
//@Success 200 {object} models.user "Successfully sign up!"
//@Failure 400 {object} models.Error "email or password error!"
//@Failure 403 {object} models.Error "E-mail already taken!"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /mssql/signup [post]
func (c Controller) MssqlSignup(MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
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

		//檢查信箱是否已經被使用過
		users, err := userRepo.MssqlCheckSignup(MsSqlDb, users)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		//check
		for _, user_db := range users {
			if user.Email == user_db.Email {
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
		if err = userRepo.MssqlInsertUser(MsSqlDb, user); err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		//加入資料庫後，密碼為空白
		user.Password = ""
		utils.SendSuccess(w, user)
		return
	}
}

//@Summary login
//@Tags Mysql_User
//@Description 登入
//@Accept json
//@Produce json
//@Param information body models.user true "個人資料"
//@Success 200 {object} models.JWT "get json-token-web"
//@Failure 400 {object} models.Error "email or password error!"
//@Failure 401 {object} models.Error "Invaild Password"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /mysql/login [post]
func (c Controller) MysqlLogin(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			jwt      models.JWT
			error    models.Error
			user     models.User
			userRepo repository.UserRepository
		)

		//decode(pointer)
		json.NewDecoder(r.Body).Decode(&user)

		//email、password不能為空
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

		//login's password
		password := user.Password

		user, err := userRepo.MysqlCheckLogin(MySqlDb, user)
		if err != nil {
			//找不到資料
			if err.Error() == "record not found" {
				error.Message = "The user does not exist!"
				utils.SendError(w, http.StatusBadRequest, error)
				return
			} else {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}
		//database password
		hashedpassword := user.Password

		//check password
		if password != hashedpassword {
			//compare the password
			//func CompareHashAndPassword(hashedPassword, password []byte) error
			if err = bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password)); err != nil {
				error.Message = "Invaild Password!"
				utils.SendError(w, http.StatusUnauthorized, error)
				return
			}
		}

		//create token
		token, err := utils.MysqlGenerateToken(user)
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

//@Summary login
//@Tags Mssql_User
//@Description 登入
//@Accept json
//@Produce json
//@Param information body models.user true "個人資料"
//@Success 200 {object} models.JWT "get json-token-web"
//@Failure 400 {object} models.Error "email or password error!"
//@Failure 401 {object} models.Error "Invaild Password"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /mssql/login [post]
func (c Controller) MssqlLogin(MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			jwt      models.JWT
			error    models.Error
			user     models.User
			userRepo repository.UserRepository
		)

		//decode(pointer)
		json.NewDecoder(r.Body).Decode(&user)

		//email、password不能為空
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

		//login's password
		password := user.Password

		user, err := userRepo.MssqlCheckLogin(MsSqlDb, user)
		if err != nil {
			//找不到資料
			if err.Error() == "record not found" {
				error.Message = "The user does not exist!"
				utils.SendError(w, http.StatusBadRequest, error)
				return
			} else {
				error.Message = "Server(database) error!"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}
		//database password
		hashedpassword := user.Password

		//check password
		if password != hashedpassword {
			//compare the password
			//func CompareHashAndPassword(hashedPassword, password []byte) error
			if err = bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password)); err != nil {
				error.Message = "Invaild Password!"
				utils.SendError(w, http.StatusUnauthorized, error)
				return
			}
		}

		//create token
		token, err := utils.MssqlGenerateToken(user)
		if err != nil {
			error.Message = "Generate Token error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.WriteHeader(http.StatusOK)
		jwt.Token = token
		utils.SendSuccess(w, jwt)
	}
}
