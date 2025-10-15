package models

type Group struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	ParentGroupID *uint
	ParentGroup   *Group  `gorm:"foreignKey:ParentGroupID"`
	SubGroups     []Group `gorm:"foreignKey:ParentGroupID"`
	Users         []User  `gorm:"foreignKey:GroupID"`
}
