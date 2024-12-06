package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

/*
Station 站点信息

- Name： 站名

- Routes： 经过该站的路线

- Postion： 所在城市
*/
type Station struct {
	ID      uint16   `json:"station_id"`
	Name    string   `json:"station_name"`
	Postion string   `json:"station_postion"`
	Routes  []uint16 `json:"route_ids"`
}

/*
Route 路线信息
*/
type Route struct {
	ID       uint16          `json:"route_id"`
	Price_pk decimal.Decimal `json:"price_pk"`
}

type AvailableRoute struct {
	Route_id              uint16          `json:"route_id"`
	Code                  string          `json:"code"`
	Station_ids           interface{}     `json:"station_ids"`
	Station_distances     interface{}     `json:"station_distances"`
	Station_expected_time []time.Duration `json:"Station_expected_time"`
	Start_station_index   uint16          `json:"start_station_index"`
	End_station_index     uint16          `json:"end_station_index"`
	Price_pk              decimal.Decimal `json:"price_pk"`
	Available_seats       uint16          `json:"available_seats"`
	Avg_speed             float64         `json:"avg_speed"`
	Seats                 interface{}     `json:"seats"`
}

type Preferences struct {
	// 在 Departure_time_before 时间前出发
	// 在 Departure_time_after 时间后出发
	// 必须满足 Departure_time_before > Departure_time_after
	Departure_time_before time.Time `json:"departure_time_before"`
	Departure_time_after  time.Time `json:"departure_time_after"`

	// 在 Arrival_time_before 时间前到达
	// 在 Arrival_time_after 时间后到达
	// 必须满足 Departure_time_before > Departure_time_after
	Arrival_time_before time.Time `json:"arrival_time_before"`
	Arrival_time_after  time.Time `json:"arrival_time_after"`
}

type RouteRequest struct {
	Start_station string `json:"start_station"`
	End_station   string `json:"end_station"`
	//
	User_preferences Preferences `json:"preferences"`
}
