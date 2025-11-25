package entity

import "time"

type Post struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"size:200" json:"title"`
	Content     string    `gorm:"type:text" json:"content"`
	Category    string    `gorm:"size:100" json:"category"`
	Status      string    `gorm:"size:100" json:"status"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date"`
}

func (u *Post) TableName() string {
	return "posts"
}
