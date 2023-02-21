package api

import (
	"net/http"
	"strconv"
	middlewarejwt "test_gin/common/middleware/jwt"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/google/wire"

	"test_gin/app/domain"
	"test_gin/app/models"
	service "test_gin/app/services"
	"test_gin/app/util"
	"test_gin/common/codes"
	"test_gin/common/logger"

	"github.com/gin-gonic/gin"
)

var ProviderSet = wire.NewSet(NewUser)

type User struct {
	Log     logger.ILogger       `inject:""`
	Jwt     middlewarejwt.JWT    `inject:""`
	Service service.IUserService `inject:""`
}

func NewUser(log *logger.Logger, jwt *middlewarejwt.JWT, user service.IUserService) *User {
	return &User{
		Log:     log,
		Jwt:     *jwt,
		Service: user,
	}
}

// @Summary 获取用户信息
// @Tags user
// @Accept application/json
// @Produce json
// @Router /user/info [get]
// @Success 200 {object} domain.Users
func (a *User) GetUserInfo(c *gin.Context) {
	roles := jwt.ExtractClaims(c)
	userName := roles["userName"].(string)
	avatar := a.Service.GetUserAvatar(userName)
	code := codes.SUCCESS
	data := domain.Users{Avatar: *avatar, Name: userName}
	util.RespData(c, http.StatusOK, code, &data)
}

//Logout 退出登录
func (a *User) Logout(c *gin.Context) {
	util.RespOk(c, http.StatusOK, codes.SUCCESS)
}

//GetUsers 获取用户信息
func (a *User) GetUsers(c *gin.Context) {
	var maps string
	code := codes.SUCCESS
	name := c.Query("name")
	if name != "" {
		maps = "username LIKE '%" + name + "%'"
	}
	page, pagesize := util.GetPage(c)
	data := a.Service.GetUsers(page, pagesize, maps)
	util.RespData(c, http.StatusOK, code, data)
}

//AddUser 新建用户
func (a *User) AddUser(c *gin.Context) {
	user := models.User{}
	code := codes.InvalidParams
	err := c.Bind(&user)
	if err != nil {
		a.Log.Error(err)
		util.RespFail(c, http.StatusOK, codes.ERROR, "参数出错!")
		return
	} else {
		roles := jwt.ExtractClaims(c)
		createdBy := roles["userName"].(string)
		user.CreatedBy = createdBy
		user.State = 1
		if !a.Service.ExistUserByName(user.Username) {
			if a.Service.AddUser(&user) {
				code = codes.SUCCESS
			} else {
				code = codes.ERROR
			}
		} else {
			code = codes.ErrExistUser
		}
	}

	util.RespOk(c, http.StatusOK, code)
}

//UpdateUser 修改用户
func (a *User) UpdateUser(c *gin.Context) {
	user := models.User{}
	code := codes.InvalidParams
	err := c.Bind(&user)
	if err != nil {
		a.Log.Error(err)
		util.RespFail(c, http.StatusOK, codes.ERROR, "参数出错!")
		return
	} else if user.ID == 0 {
		code = codes.ERROR
	} else {
		roles := jwt.ExtractClaims(c)
		modifiedBy := roles["userName"].(string)
		user.ModifiedBy = modifiedBy
		if a.Service.UpdateUser(&user) {
			code = codes.SUCCESS
		} else {
			code = codes.ERROR
		}
	}
	util.RespOk(c, http.StatusOK, code)
}

//DeleteUser 删除用户
func (a *User) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := codes.SUCCESS
	if !a.Service.DeleteUser(id) {
		code = codes.ERROR
		util.RespFail(c, http.StatusOK, code, "不允许删除admin账号!")
	} else {
		util.RespOk(c, http.StatusOK, code)
	}
}
