package services

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/2418071565/GoTicket/dto"
	"github.com/2418071565/GoTicket/models"
	"github.com/2418071565/GoTicket/pkg/jwt"
	"github.com/2418071565/GoTicket/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Authentication(ctx *gin.Context, phone string, password string) (*dto.User, *dto.JsonErrorStruct) {
	// 判断是否存在该用户，如果存在判断密码是否正确
	user, err := models.User{}.GetUserByPhone(phone)
	if err != nil {
		return user, &dto.JsonErrorStruct{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("login failed with err: %s", err),
		}
	}
	if user.Password != password {
		return user, &dto.JsonErrorStruct{
			Code:    http.StatusOK,
			Message: "phone or password is incorrect",
		}
	}

	// 生成 token，返回给用户
	if err := GenerateToken(ctx, user.ID, user.Is_admin); err != nil {
		return user, &dto.JsonErrorStruct{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}
	return user, nil
}

// 从请求体中获取用户信息
func GetUserInfoFromRequest(ctx *gin.Context) (*dto.User, error) {
	user := &dto.User{}
	if err := ctx.ShouldBindJSON(user); err != nil {
		return nil, err
	}
	if len(user.Name) > 20 {
		logger.Infof("name cannot exceed 20 characters")
		return nil, fmt.Errorf("name cannot exceed 20 characters")
	}
	if user.Sex != "Male" && user.Sex != "Female" {
		logger.Infof("invalid sex: %s", user.Sex)
		return nil, fmt.Errorf("invalid sex: %s", user.Sex)
	}
	if len(user.Password) > 20 {
		logger.Infof("password cannot exceed 20 characters")
		return nil, fmt.Errorf("password cannot exceed 20 characters")
	}
	if ok, err := regexp.MatchString(`^\d+$`, user.Phone); !ok || err != nil {
		logger.Infof("invalid phone number: %s", user.Phone)
		return nil, fmt.Errorf("invalid phone number: %s", user.Phone)
	}
	if ok, err := regexp.MatchString(`^\d{15}$|^\d{17}(\d|X|x)$`, user.Id_number); !ok || err != nil {
		logger.Infof("invalid id_number: %s", user.Id_number)
		return nil, fmt.Errorf("invalid id_number: %s", user.Id_number)
	}

	user.Create_date = time.Now()
	return user, nil
}

func GenerateToken(ctx *gin.Context, id uint32, is_admin bool) error {
	token, err := jwt.GenerateToken(id, is_admin)
	if err != nil {
		logger.Errorf("generate token failed with err: %s", err)
		return fmt.Errorf("generate token failed with err: %s", err)
	}
	ctx.Header("Authorization", token)
	return nil
}
