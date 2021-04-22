package models

// json 格式
// password，設為忽略，表示 response 就不會回傳此欄位
type User struct {
  Id uint `json:"id"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  Email string `json:"email" gorm:"unique"`
  Password []byte `json:"-"`
}