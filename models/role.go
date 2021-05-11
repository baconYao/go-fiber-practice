package models

type Role struct {
  Id uint `json:"id"`
  Name string `json:"name"`
  Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions"`   // 建立 many to many 的 table (table名為role_permissions)
}