package dto

import (
	"github.com/gin-gonic/gin"
	"jiac_gateway/public"
)

type AdminLoginInput struct {
	// tag:介绍  example 设默认值，comment 定义错误输出，validate 必填，form 由request转为struct——json相反
	//validate的is_valid_username  定义在middleware/translate.go  实现对UserName的校验  如果不等于adminss则报错
	UserName string `form:"name" json:"name" comment:"管理员用户名"  validate:"required,is_valid_username" example:"admin"`
	Password string `form:"pwd" json:"pwd" comment:"密码"  validate:"required" example:"123456"`
}

//定义结构体校验方法
func (param *AdminLoginInput) BindValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}


type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""` //token
}