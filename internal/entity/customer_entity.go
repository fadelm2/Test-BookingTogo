package entity

import "time"

type Customer struct {
	ID            int       `gorm:"column:cst_id;primaryKey;autoIncrement" json:"id"`
	Name          string    `gorm:"column:cst_name;;not null" json:"name"`
	Address       string    `gorm:"column:cst_address;" json:"address"`
	Phone         string    `gorm:"column:cst_phone;" json:"phone"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	NationalityId int       `gorm:"column:nationality_id" json:"nationality_id"`

	// belongs-to Nationality
	Nationality Nationality `gorm:"foreignKey:nationality_id;references:nationality_id" json:"nationality"`

	// Relasi has-many ke family_list
	Family []FamilyList `gorm:"foreignKey:cst_id;references:cst_id" json:"family"`
}

func (Customer) TableName() string {
	return "customer"
}
