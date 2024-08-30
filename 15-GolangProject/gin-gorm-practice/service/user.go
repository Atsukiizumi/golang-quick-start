package service

import (
	"GolangProject/gin-gorm-practice/helper"
	"GolangProject/gin-gorm-practice/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// Login
// @Schemes	user
// @Description 用户登录
// @Produce json
// @Tags 公共方法
// @Summary 用户登录
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json{"code":"200","msg":""}
// @Router /login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	salt := ""
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码为空",
		})
		return
	}

	// md5转换密文
	password = helper.GetSaltMD5(password, salt)

	data := new(models.UserBasic)
	err := models.DB.First(&data, "name =? AND password =? ", username, password).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GetUserBasic Error:" + err.Error(),
		})
		return
	}

	token, tokenError := helper.GenerateToken(data.Identity, data.Name)
	if tokenError != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GenerateToken Error:" + tokenError.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

// SendCode
// @Schemes	user
// @Description 发送验证码
// @Produce json
// @Tags 公共方法
// @Summary 发送验证码
// @Param mail formData string false "user_smail"
// @Success 200 {string} json{"code":"200","msg":""}
// @Router /send-code [post]
func SendCode(c *gin.Context) {
	email := c.PostForm("mail")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}

	code := "111122"
	err := helper.SendCode(email, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "SendCode Error:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
	return
}
