package model

type Author struct {
	Aid   uint   `json:"aid" gorm:"primaryKey"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}
