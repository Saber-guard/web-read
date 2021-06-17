package model

type DayDataModel struct {
	BaseModel
	Code          string  `gorm:"comment:上市代码;size:10;default:'';not null"` //f12
	TodayStart    float32 `gorm:"comment:今开;default:0;not null"`            //f17
	YestdayEnd    float32 `gorm:"comment:昨收;default:0;not null"`            //f18
	Highest       float32 `gorm:"comment:最高价;default:0;not null"`           //f15
	Minimum       float32 `gorm:"comment:最低价;default:0;not null"`           //f16
	DealNum       int32   `gorm:"comment:成交量;default:0;not null"`           //f5
	DealMoney     float64 `gorm:"comment:成交额;default:0;not null"`           //f6
	IncreaseRange float32 `gorm:"comment:涨跌幅;default:0;not null"`           //f3
	IncreaseMoney float32 `gorm:"comment:涨跌额;default:0;not null"`           //f4
	TurnoverRate  float32 `gorm:"comment:换手率;default:0;not null"`           //f8
	ProfitRatio   float32 `gorm:"comment:市盈率(动态);default:0;not null"`       //f9
	ValueRatio    float32 `gorm:"comment:市净率;default:0;not null"`           //f23

	Ext string `gorm:"comment:扩展信息;type:json;not null"`
}

func (DayDataModel) TableName() string {
	return "day_data"
}
