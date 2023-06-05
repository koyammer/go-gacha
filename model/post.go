package model

type Post struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID uint   `json:"user_id"`
	User   User   `json:"user"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
