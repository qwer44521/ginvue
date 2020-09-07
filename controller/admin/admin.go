package admin

import (
	"fmt"
	"ginvue/serializer"
	service2 "ginvue/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//添加管理员
func Addadmin(c *gin.Context) {
	var service service2.Administrators
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  0,
			Msg:   "参数缺失",
			Error: err.Error(),
		})

	} else {
		//添加操作
		res := service.Addadmin()
		c.JSON(http.StatusOK, res)
	}
}

//@title:管理员登录
//
//管理员登录
func AdminLogin(c *gin.Context) {
	var service service2.Adminlogin
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  0,
			Msg:   "参数缺失",
			Error: err.Error(),
		})
	} else {
		res := service.Adminlogin()
		c.JSON(http.StatusOK, res)
	}

}

//返回数据信息
func AdminInfo(c *gin.Context) {
	var service service2.AdminInfo
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  0,
			Msg:   "未携带token",
			Error: err.Error(),
		})

	} else {
		res := service.AdminInfo()
		fmt.Printf("%#v", res)
		c.JSON(http.StatusOK, res)
	}

}
