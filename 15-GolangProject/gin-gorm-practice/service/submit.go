package service

import (
	"GolangProject/gin-gorm-practice/define"
	"GolangProject/gin-gorm-practice/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetSubmitList
// @Schemes	submit
// @Description 提交列表
// @Produce json
// @Tags 公共方法
// @Summary 提交列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param problem_identity query string false "problem_identity"
// @Param user_identity query string false "user_identity"
// @Param status query int false "status"
// @Success 200 {string} json{"code":"200","msg":"","data":""}
// @Router /submit-list [get]
func GetSubmitList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv Error: ", err)
		return
	}

	page = (page - 1) * size

	var count int64
	list := make([]*models.SubmitBasic, 0)

	problemIdentity := c.Query("problem_identity")
	userIdentity := c.Query("user_identity")
	status, _ := strconv.Atoi(c.Query("status"))

	tx := models.GetSubmitList(problemIdentity, userIdentity, status)

	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("Get Problem List Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"count": count,
			"list":  list,
		},
		"msg": "success",
	})

}
