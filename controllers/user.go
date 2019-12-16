package controllers

import (
	"ApiForTwoDb/driver"
	models "ApiForTwoDb/model"
	"ApiForTwoDb/repository"
	"ApiForTwoDb/utils"
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Controller struct{}

//sign up
//mysql
func (c Controller) MysqlSignup(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var error models.Error
		var user models.User
		var users []models.User

		//decode(須指標)
		//寫入user
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

		userRepo := repository.UserRepository{}
		//檢查信箱是否已經被使用過
		users, err = userRepo.MysqlCheckSignup(MySqlDb, users)
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
		err = userRepo.MysqlInsertUser(MySqlDb, user)
		if err != nil {
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

//sign up
//mssql
func (c Controller) MssqlSignup(MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var error models.Error
		var user models.User
		var users []models.User

		//decode(須指標)
		//寫入user
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

		userRepo := repository.UserRepository{}

		//檢查信箱是否已經被使用過
		users, err = userRepo.MssqlCheckSignup(MsSqlDb, users)
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
		err = userRepo.MssqlInsertUser(MsSqlDb, user)
		if err != nil {
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

//login
//mysql
func (c Controller) MysqlLogin(MySqlDb *driver.MySqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var jwt models.JWT
		var err error
		var error models.Error
		var user models.User

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

		userRepo := repository.UserRepository{}

		user, err = userRepo.MysqlCheckLogin(MySqlDb, user)
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
			err = bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
			if err != nil {
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

//mssql
func (c Controller) MssqlLogin(MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var jwt models.JWT
		var err error
		var error models.Error
		var user models.User

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

		userRepo := repository.UserRepository{}

		user, err = userRepo.MssqlCheckLogin(MsSqlDb, user)
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
			err = bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
			if err != nil {
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
