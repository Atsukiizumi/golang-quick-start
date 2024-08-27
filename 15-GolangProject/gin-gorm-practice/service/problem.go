package service

import (
	"GolangProject/gin-gorm-practice/define"
	"GolangProject/gin-gorm-practice/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList
// @Schemes	problem
// @Description 获取问题列表
// @Produce json
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json{"code":"200","msg":"","data":""}
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {
	size, err_size := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if err_size != nil {
		log.Println("GetProblemList Size strconv Error: ", err_size)
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv Error: ", err)
		return
	}
	page = (page - 1) * size
	var count int64
	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")

	tx := models.GetProblemList(keyword, categoryIdentity)

	list := make([]*models.ProblemBasic, 0)
	err = tx.Count(&count).Omit("content").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("Get Problem List Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"data": nil,
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

// GetProblemDetail
// @Schemes	problem
// @Description 问题详情
// @Produce json
// @Tags 公共方法
// @Summary 问题详情
// @Param identity query string false "problem_identity"
// @Success 200 {string} json{"code":"200","msg":"","data":""}
// @Router /problem-detail [get]
func GetProblemDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题唯一标识不能为空",
			"data": nil,
		})
		return
	}
	problemBasic := new(models.ProblemBasic)
	err := models.DB.Where("identities = ?", identity).
		Preload("ProblemCategories").
		Preload("ProblemCategories.CategoryBasic").
		First(&problemBasic).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "问题不存在",
				"data": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": nil,
			"msg":  "Get Problem Detail Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": problemBasic,
		"msg":  "success",
	})
	return

}
