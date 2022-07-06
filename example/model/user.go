package model

type User struct {
	BaseModel
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	CompanyID int64    `json:"company_id"`
	Company   *Company `json:"company" foreign:"company_id"`
}

func (u User) TableName() string {
	return "users"
}
