package models

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

type mssqlgetsome struct {
	Key1 string
	Key2 string
	Key3 string
}

type mssqlupdate struct {
	Key1  string
	Death string
}

type mssqldelete struct {
	Key1 string
}
