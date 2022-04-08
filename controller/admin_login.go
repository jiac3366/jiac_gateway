package controller

import (
	"github.com/gin-gonic/gin"
	"jiac_gateway/dto"
	"jiac_gateway/middleware"
)

type AdminLoginController struct {
}

// AdminLoginRegister handlefunc注册到Controller,
// Controller绑定在RouterGroup，RouterGroup注册到Route
// Route -- RouteGroup -- Controller -->Func
func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
}


// ListPage godoc
// @Summary 管理员登陆
// @Description 管理员登陆
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (demo *AdminLoginController) AdminLogin(c *gin.Context) {
	param := &dto.AdminLoginInput{}
	if err := param.BindValidParams(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	// 以下数据会以data形式输出
	out := &dto.AdminLoginOutput{Token: param.UserName}
	middleware.ResponseSuccess(c, out)
}

