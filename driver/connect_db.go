package driver

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MySqlDb struct {
	*gorm.DB
}

type MsSqlDb struct {
	*gorm.DB
}

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

//mysql初始化連線
func (msu *MySqlUser) Init() *MySqlDb {

	//完整的資料格式: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	MysqlDataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		msu.User,
		msu.Password,
		msu.Host,
		msu.Port,
		msu.Database)

	//開啟資料庫連線
	Db, err := gorm.Open("mysql", MysqlDataSourceName)
	if err != nil {
		log.Println(err)
	}
	//設定最大連接
	Db.DB().SetMaxIdleConns(msu.MaxIdle)
	Db.DB().SetMaxOpenConns(msu.MaxOpen)

	return &MySqlDb{Db}
}

//mssql初始化連線
func (mssu *MsSqlUser) Init() *MsSqlDb {

	MssqlDataSourceName := fmt.Sprintf("serve=%s;user id=%s;password=%s;port=%s;database=%s",
		mssu.Host,
		mssu.User,
		mssu.Password,
		mssu.Port,
		mssu.Database)

	Db, err := gorm.Open("mssql", MssqlDataSourceName)
	if err != nil {
		log.Println(err)
	}
	Db.DB().SetMaxIdleConns(mssu.MaxIdle)
	Db.DB().SetMaxOpenConns(mssu.MaxOpen)

	return &MsSqlDb{Db}
}
