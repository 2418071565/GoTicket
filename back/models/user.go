package models

import (
	"time"

	"github.com/2418071565/GoTicket/dto"
	"github.com/2418071565/GoTicket/pkg/config"
	"github.com/2418071565/GoTicket/pkg/logger"
	"github.com/2418071565/GoTicket/storage/db"
)

type User struct{}

type Admin struct{}

func init() {
	var cnt int64
	if err := db.DB.Table("users").Where("phone = ?", "goticket").Count(&cnt).Error; err != nil {
		logger.Errorf("init admin `goticket` failed with err: %s", err)
	}
	if cnt == 0 {
		default_admin := &dto.User{
			Name:        "goticket",
			Sex:         "Male",
			Password:    config.DEFAULT_ADMIN_PASSWORD,
			Phone:       config.DEFAULT_ADMIN,
			Create_date: time.Now(),
			Id_number:   "",
		}
		db.DB.Create(default_admin)
		db.DB.Table("admins").Create(map[string]interface{}{"id": default_admin.ID})
	}
}

// 通过手机号判断用户是否存在
func (User) IsUserExists(phone string) (bool, error) {
	var cnt int64
	if err := db.DB.Table("users").Select("id").Where("phone = ?", phone).Count(&cnt).Error; err != nil {
		return false, err
	}
	if cnt == 0 {
		return false, nil
	}
	return true, nil
}

func (User) GetUserByPhone(phone string) (*dto.User, error) {
	user := &dto.User{}
	if err := db.DB.Table("users").
		Select("users.id, users.name, users.sex, users.password, users.phone, users.create_date, users.id_number, admins.id IS NOT NULL AS is_admin").
		Joins("LEFT JOIN admins ON users.id = admins.id").
		Where("users.phone = ?", phone).
		Scan(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (User) GetUserById(id uint32) (*dto.User, error) {
	user := &dto.User{}
	if err := db.DB.Table("users").
		Select("users.id, users.name, users.sex, users.password, users.phone, users.create_date, users.id_number, admins.id IS NOT NULL AS is_admin").
		Joins("LEFT JOIN admins ON users.id = admins.id").
		Where("users.id = ?", id).
		Scan(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (User) AddUser(user *dto.User) (uint32, error) {
	new_user := dto.User{
		ID:          user.ID,
		Name:        user.Name,
		Sex:         user.Sex,
		Password:    user.Password,
		Phone:       user.Phone,
		Create_date: user.Create_date,
		Id_number:   user.Id_number,
	}
	err := db.DB.Create(&new_user).Error
	return new_user.ID, err
}

func (Admin) AddAdmin(user *dto.User) error {
	admin := &dto.Admin{ID: user.ID}
	err := db.DB.Table("admins").Create(admin).Error
	return err
}
