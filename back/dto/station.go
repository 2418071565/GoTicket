package dto

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
