package model

type Company struct {
	BaseModel
	Name     string  `json:"name"`
	PicName  string  `json:"pic_name"`
	PicPhone *string `json:"pic_phone"`
}

func (c Company) TableName() string {
	return "companies"
}
