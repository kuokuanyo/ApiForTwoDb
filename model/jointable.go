package models

type JoinTable struct {
	Key1             string `gorm:"column:key1" json:"key1"`
	Key2             string `gorm:"column:key2" json:"key2"`
	Key3             string `gorm:"column:key3" json:"key3"`
	Number           int    `gorm:"column:number" json:"number"`
	Gender           int    `gorm:"column:gender" json:"gender"`
	Birth            int    `gorm:"column:birth" json:"birth"`
	Injury_degree    string `gorm:"column:injury_degree" json:"injury_degree"`
	Injury_position  int    `gorm:"column:injury_position" json:"injury_position"`
	Protection       int    `gorm:"column:protection" json:"protection"`
	Phone            int    `gorm:"column:phone" json:"phone"`
	Person           string `gorm:"column:person" json:"person"`
	Car              string `gorm:"column:car" json:"car"`
	Action_status    int    `gorm:"column:action_status" json:"action_status"`
	Qualification    int    `gorm:"column:qualification" json:"qualification"`
	License          int    `gorm:"column:license" json:"license"`
	Drinking         int    `gorm:"column:drinking" json:"drinking"`
	Hit              int    `gorm:"column:hit" json:"hit"`
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
