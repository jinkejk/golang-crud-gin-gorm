package model

type Tags struct {
	Id   *int    `gorm:"column:id;type:int;primaryKey"`
	Name *string `gorm:"column:name;type:varchar(255)"`
}

func (p Tags) TableName() string {
	return "tags"
}
