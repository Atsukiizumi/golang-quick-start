package service

import (
	"GolangProject/gin-gorm-practice/define"
	"GolangProject/gin-gorm-practice/helper"
	"GolangProject/gin-gorm-practice/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetCategoryList
// @Schemes	category
// @Description 获取分类列表
// @Produce json
// @Tags 管理员私有方法
// @Summary 分类列表
// @Param authorization header string true "token"
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Success 200 {string} json{"code":"200","msg":"","data":""}
// @Router /admin/category-list [get]
func GetCategoryList(c *gin.Context) {
	size, errSize := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if errSize != nil {
		log.Println("GetCategoryList Size strconv Error: ", errSize)
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetCategoryList Page strconv Error: ", err)
		return
	}
	page = (page - 1) * size
	var count int64
	keyword := c.Query("keyword")

	categoryList := make([]*models.CategoryBasic, 0)
	err = models.DB.Model(&models.CategoryBasic{}).Where("name LIKE ?", "%"+keyword+"%").Count(&count).Limit(size).Offset(page).Find(&categoryList).Error
	if err != nil {
		log.Println("Get GetCategoryList Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": nil,
			"msg":  "获取分类列表失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"count": count,
			"list":  categoryList,
		},
		"msg": "success",
	})
}

// CategoryCreate
// @Schemes	category
// @Description 创建分类
// @Produce json
// @Tags 管理员私有方法
// @Summary 创建分类
// @Param authorization header string true "token"
// @Param name formData string true "name"
// @Param parentId formData int false "parentId"
// @Success 200 {string} json{"code":"200","msg":""}
// @Router /admin/category-create [post]
func CategoryCreate(c *gin.Context) {
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parentId"))
	identity := helper.GetUUID()

	category := &models.CategoryBasic{
		Identity: identity,
		Name:     name,
		ParentId: parentId,
	}

	err := models.DB.Create(&category).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "创建分类失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

// CategoryDelete
// @Schemes	category
// @Description 删除分类
// @Produce json
// @Tags 管理员私有方法
// @Summary 删除分类
// @Param authorization header string true "token"
// @Param identity query string true "identity"
// @Success 200 {string} json{"code":"200","msg":""}
// @Router /admin/category-delete [delete]
func CategoryDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
		return
	}

	var count int64
	// 先找出所有的category关联的
	err := models.DB.Model(new(models.ProblemCategory)).
		Where("category_id = (SELECT id FROM category_basic where identity = ? and deleted_at = null LIMIT 1 )", identity).
		Count(&count).Error
	if err != nil {
		log.Println("Get ProblemCategory err: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "获取分类关联的问题失败",
		})
		return
	}

	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该分类下有题目，无法删除",
		})
		return
	}

	err = models.DB.Where("identity = ?", identity).Delete(new(models.CategoryBasic)).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除分类失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})

}

// CategoryUpdate
// @Schemes	category
// @Description 修改分类
// @Produce json
// @Tags 管理员私有方法
// @Summary 修改分类
// @Param authorization header string true "token"
// @Param identity formData string true "identity"
// @Param name formData string true "name"
// @Param parentId formData int false "parentId"
// @Success 200 {string} json{"code":"200","msg":""}
// @Router /admin/category-update [put]
func CategoryUpdate(c *gin.Context) {
	name := c.PostForm("name")
	identity := c.PostForm("identity")
	parentId, _ := strconv.Atoi(c.PostForm("parentId"))

	if name == "" || identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
		return
	}

	category := &models.CategoryBasic{
		Identity: identity,
		Name:     name,
		ParentId: parentId,
	}

	err := models.DB.Model(new(models.CategoryBasic)).Where("identity = ?", identity).Updates(&category).Error
	if err != nil {
		log.Println("CategoryUpdate err: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "修改分类失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
