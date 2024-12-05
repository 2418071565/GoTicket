package dto

import "time"

/*
User 用户信息

- Name：用户名字

- Sex：性别

- Password：密码

- PhoneNumber：电话号

- CreateDate：创建日期

- IDNumber：身份证号

- Orders：用户所有的订单

- Tickers：用户所有的车票
*/
type User struct {
	ID          uint32    `json:"id"`
	Name        string    `json:"name"`
	Sex         string    `json:"sex"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
	Create_date time.Time `json:"create_date"`
	Id_number   string    `json:"id_number"`
	Is_admin    bool      `json:"is_admin"`
}

type Admin struct {
	ID uint32 `json:"id"`
}
