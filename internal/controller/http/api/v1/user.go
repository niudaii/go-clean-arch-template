package v1

import (
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/jwt"
	"go-clean-template/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	UserUsecase entity.UserUsecase
}

func (a UserApi) GetInfo(c *gin.Context) {
	userID := jwt.GetUserID(c)
	user, err := a.UserUsecase.FindByID(userID)
	if err != nil {
		response.ErrorWithMessage("查询失败", err, c)
	} else {
		role := entity.Role{
			RoleName: user.Authority.AuthorityName,
			Value:    user.Authority.AuthorityName,
		}
		data := entity.UserInfo{
			UserID:   user.ID,
			Username: user.Username,
			Avatar:   "https://q1.qlogo.cn/g?b=qq&nk=190848757&s=640",
			HomePath: "/task/index",
			Roles: []entity.Role{
				role,
			},
		}
		response.Ok(data, "查询成功", c)
	}
}

func (a UserApi) Logout(c *gin.Context) {
	response.OkWithMessage("注销成功", c)
}
