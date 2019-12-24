package controllers

/*
//@Summary 取得關聯表所有資料
//@Tags JoinTable
//@Description 取得所有資料
//@Accept json
//@Produce json
//@Success 200 {object} models.JoinTable "join data"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /v1/join/getall [get]
func (c Controller) JoinGetAll(MySqlDb *driver.MySqlDb, MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			peoples    []models.People
			events     []models.Event
			jointables []models.JoinTable
			error      models.Error
			userRepo   repository.UserRepository
		)

		//更新關聯資料庫
		if err := userRepo.UpdateJoinData(MySqlDb, MsSqlDb, peoples, events); err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		//取得所有資料
		jointables, err := userRepo.QueryAllJoinData(MySqlDb, jointables)
		if err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		utils.SendSuccess(w, jointables)
	}
}
*/
/*
//@Summary 取得關聯表部分資料
//@Tags JoinTable
//@Description 取得部分資料
//@Accept json
//@Produce json
//@Param key1 path int true "Key1"
//@Param key2 path int true "Key2"
//@Param key3 path int true "Key3"
//@Success 200 {object} models.JoinTable "join data"
//@Failure 404 {object} models.Error "The user does not exist!"
//@Failure 500 {object} models.Error "Serve(database) error!"
//@Router /v1/join/getsome/{key1}/{key2}/{key3} [get]
func (c Controller) JoinGetSome(MySqlDb *driver.MySqlDb, MsSqlDb *driver.MsSqlDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			error      models.Error
			jointable  models.JoinTable
			jointables []models.JoinTable
			userRepo   repository.UserRepository
		)

		//return map
		//func Vars(r *http.Request) map[string]string
		params := mux.Vars(r)
		jointable.Key1 = params["key1"]
		jointable.Key2 = params["key2"]
		jointable.Key3 = params["key3"]

		//更新關聯資料庫

		if err := userRepo.UpdateJoinData(MySqlDb, MsSqlDb, []models.People{}, []models.Event{}); err != nil {
			error.Message = "Server(database) error!"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		//尋找部分資料
		jointables, err := userRepo.QuerySomeJoinData(MySqlDb, jointables, jointable)
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
		utils.SendSuccess(w, jointables)
	}
}
*/
