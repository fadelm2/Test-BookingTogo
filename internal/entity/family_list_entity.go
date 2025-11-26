package entity

type FamilyList struct {
	ID         int    `gorm:"column:fl_id;primaryKey;autoIncrement" json:"id"`
	CustomerID int    `gorm:"column:cst_id" json:"customer_id"`
	Relation   string `gorm:"column:fl_relation" json:"relation"`
	Name       string `gorm:"column:fl_name" json:"name"`
	Dob        string `gorm:"column:fl_dob" json:"dob"`
}

func (FamilyList) TableName() string {
	return "family_list"
}
