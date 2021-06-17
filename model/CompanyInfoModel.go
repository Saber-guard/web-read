package model

type CompanyInfoModel struct {
	BaseModel
	Code string `gorm:"comment:上市代码;size:10;default:'';not null"`
	Name string `gorm:"comment:上市名称;size:64;default:'';not null"`
}

func (CompanyInfoModel) TableName() string {
	return "company_info"
}
