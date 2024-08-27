package service

import (
	"GolangProject/gin-gorm-practice/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserDetail
// @Schemes	user
// @Description 用户详情
// @Produce json
// @Tags 公共方法
// @Summary 用户详情
// @Param identity query string false "user_identity"
// @Success 200 {string} json{"code":"200","msg":"","data":""}
// @Router /user-detail [get]
func GetUserDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户唯一标识不能为空",
		})
		return
	}

	data := new(models.UserBasic)
	err := models.DB.Omit("password").Where("identities =?", identity).First(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get User Detail By Identity " + identity + " Error: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
		"msg":  "success",
	})
}
