package dto

/*
Train 列车信息

- TrainType：列车类型（高铁、动车和绿皮）

- MaxCapacity：列车最大载客量

- Seats：列车座位表，[车厢][行][列]

- AvgSpeed：列车平均速度，用于预测到站时间

- TrainNumbers：该列车所有的车次
*/
type Train struct {
	ID            uint16        `gorm:"type:smallint unsigned;auto_increment"`
	Train_type    string        `gorm:"type:enum('G','D','K');not null"`
	Max_capacity  uint16        `gorm:"type:smallint unsigned;not null"`
	Seats         string        `gorm:"type:json;not null"`
	Avg_speed     float32       `gorm:"not null"`
	Train_numbers []TrainNumber `gorm:"foreignKey:TrainID"`
}
