package models

type Payment struct {
	Id           int     `json:"-" db:"id"`
	UserId       int     `json:"user_id" db:"user_id" binding:"required"`
	UserEmail    string  `json:"user_email" db:"user_email" binding:"required"`
	Sum          float32 `json:"sum" db:"summ" binding:"required"`
	Currency     string  `json:"currency" db:"currency" binding:"required"`
	CreationTime string  `json:"creation_time" db:"creation_time"`  //date
	UpdateTime   string  `json:"update_time" db:"update_time"` //date
	Status       string  `json:"status" db:"status"`
}
