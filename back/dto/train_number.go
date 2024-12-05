package dto

/*
TrainNumber 车次信息

- Status：车次状态

- AvailableSeats：剩余的座位数

- StartTime：列车开始时间
*/
type TrainNumber struct {
	ID     uint32
	Status string
}

type TrainNumberInfo struct {
	Route_id        uint16   `json:"route_id"`
	Train_number    string   `json:"train_number"`
	Start_station   string   `json:"start_station"`
	End_station     string   `json:"end_station"`
	Departure_time  string   `json:"departure_time"`
	Arrival_time    string   `json:"arrival_time"`
	Available_seats uint16   `json:"available_seats"`
	Price           float32  `json:"price"`
	All_station     []string `json:"all_station"`
}
