package model

type DayDataModel struct {
	BaseModel
	Date          string `gorm:"type:date;comment:日期;not null;index:idx_code_date,priority:2"`
	Code          string `gorm:"comment:上市代码;size:10;default:'';not null;index:idx_code_date,priority:1"` //f12
	TodayStart    string `gorm:"comment:今开;default:0;not null"`                                           //f17
	TodayEnd      string `gorm:"comment:今收;default:0;not null"`                                           //f2
	YestdayEnd    string `gorm:"comment:昨收;default:0;not null"`                                           //f18
	Highest       string `gorm:"comment:最高价;default:0;not null"`                                          //f15
	Minimum       string `gorm:"comment:最低价;default:0;not null"`                                          //f16
	DealNum       string `gorm:"comment:成交量;default:0;not null"`                                          //f5
	DealMoney     string `gorm:"comment:成交额;default:0;not null"`                                          //f6
	IncreaseRange string `gorm:"comment:涨跌幅;default:0;not null"`                                          //f3
	IncreaseMoney string `gorm:"comment:涨跌额;default:0;not null"`                                          //f4
	TurnoverRate  string `gorm:"comment:换手率;default:0;not null"`                                          //f8
	ProfitRatio   string `gorm:"comment:市盈率(动态);default:0;not null"`                                      //f9
	ValueRatio    string `gorm:"comment:市净率;default:0;not null"`                                          //f23

	Ext string `gorm:"comment:扩展信息;type:json;not null"`
}

func (DayDataModel) TableName() string {
	return "day_data"
}
