package service

import (
	"GolangProject/gin-gorm-practice/helper"
	"GolangProject/gin-gorm-practice/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
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
	salt := "golang"
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
	err := models.DB.Where("name =? AND password =? ", username, password).First(&data).Error
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
// @Param mail formData string true "user_mail"
// @Success 200 {string} json{"code":"200","msg":""}
// @Router /send-code [post]
func SendCode(c *gin.Context) {
	mail := c.PostForm("mail")
	if mail == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}

	// 生成验证码
	code := helper.GetRandCode()
	models.RDB.Set(c, mail, code, time.Second*300)
	err := helper.SendCodeOMail(mail, code)
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
		"data": map[string]interface{}{
			"mail_code": code,
		},
	})
	return
}

// Register
// @Schemes	user
// @Description 用户注册
// @Produce json
// @Tags 公共方法
// @Summary 用户注册
// @Param name formData string true "user_name"
// @Param userCode formData string true "user_code"
// @Param password formData string true "user_password"
// @Param phone formData string false "user_phone"
// @Param mail formData string true "user_mail"
// @Success 200 {string} json{"code":"200","msg":""}
// @Router /register [post]
func Register(c *gin.Context) {
	name := c.PostForm("name")
	userCode := c.PostForm("code")
	pwd := c.PostForm("password")
	phone := c.PostForm("phone")
	mail := c.PostForm("mail")
	if name == "" || userCode == "" || pwd == "" || mail == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}

	println(phone)
	// 验证码是否正确
	sysCode, err := models.RDB.Get(c, mail).Result()
	if err != nil {
		log.Printf("Get Code Error:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确，请重新获取验证码",
		})
		return
	}

	if sysCode != userCode {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确",
		})
		return
	}

	// 判断邮箱是否已存在
	var count int64
	err = models.DB.Where("mail = ?", mail).Model(new(models.UserBasic)).Count(&count).Error
	if err != nil {
		log.Printf("Get User Error:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "查询用户信息失败",
		})
		return
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该邮箱已注册",
		})
		return
	}

	userIdentity := helper.GetUUID()
	// 数据的插入
	data := models.UserBasic{
		Identity: userIdentity,
		Name:     name,
		Password: helper.GetMD5(pwd),
		Phone:    phone,
		Mail:     mail,
	}

	err = models.DB.Create(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create User Error:" + err.Error(),
		})
		return
	}

	//生成token
	token, err := helper.GenerateToken(userIdentity, name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GenerateToken Error:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}
