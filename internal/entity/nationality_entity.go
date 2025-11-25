package entity

type Nationality struct {
	NationalityId int    `gorm:"column:nationality_id;primaryKey;autoIncrement" json:"id"`
	Name          string `gorm:"column:nationality_name" json:"name"`
	Code          string `gorm:"column:nationality_code" json:"code"`
}

func (Nationality) TableName() string {
	return "nationality"
}
