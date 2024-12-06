package dto

/*
Ticket 车票信息

- StartStation：起始站点名

- EndStation：终点站点名

- DepartureTime： 出发时间

- ArrivalTime： 到达时间

- Carriage：座位所在车厢

- Row：座位所在行

- Seat：座位所在列

- Status：车票状态

- Order：该票的订单

- UserID: 车票所属用户
*/
type Tickets struct {
	ID             uint32 `json:"ticket_id"`
	Start_station  string `json:"start_station"`
	End_station    string `json:"end_station"`
	Departure_time string `json:"departure_time"`
	Arrival_time   string `json:"arrival_time"`
	Carriage       uint8  `json:"carriage"`
	Row            uint8  `json:"row"`
	Seat           uint8  `json:"seat"`
	Status         string `json:"status"`
}
