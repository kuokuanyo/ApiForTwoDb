package models

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

type mysqlgetsome struct {
	Key1 string
	Key2 string
	Key3 string
}

type mysqlupdate struct {
	Key1  string
	Birth int
}

type mysqldelete struct {
	Key1 string
}
