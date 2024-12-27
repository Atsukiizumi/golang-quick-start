package service

import (
	"GolangProject/gin-gorm-practice/define"
	"GolangProject/gin-gorm-practice/helper"
	"GolangProject/gin-gorm-practice/models"
	"encoding/json"
	"errors"
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
	size, errSize := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if errSize != nil {
		log.Println("GetProblemList Size strconv Error: ", errSize)
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
	err := models.DB.Where("identity = ?", identity).
		Preload("ProblemCategories").
		Preload("ProblemCategories.CategoryBasic").
		First(&problemBasic).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
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

// ProblemCreate
// @Schemes	problem_basic
// @Description 创建问题
// @Produce json
// @Tags 管理员私有方法
// @Summary 创建问题
// @Param authorization header string true "token"
// @Param title formData string true "title"
// @Param content formData string true "content"
// @Param max_runtime formData int false "max_runtime"
// @Param max_mem formData int false "max_mem"
// @Param category_ids formData []string false "category_ids" collectionFormat(multi)
// @Param test_cases formData []string true "test_cases" collectionFormat(multi)
// @Success 200 {string} json{"code":"200","msg":""}
// @Router /admin/problem-create [post]
func ProblemCreate(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	maxRuntime, _ := strconv.Atoi(c.PostForm("max_runtime"))
	maxMem, _ := strconv.Atoi(c.PostForm("max_mem"))
	categoryIds := c.PostFormArray("category_ids")
	testCases := c.PostFormArray("test_cases")
	if title == "" || content == "" || len(testCases) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}
	identity := helper.GetUUID()
	problems := models.ProblemBasic{
		Identity:    identity,
		Title:       title,
		Content:     content,
		Max_Men:     maxMem,
		Max_runtime: maxRuntime,
	}
	//处理分类
	categoryBasics := make([]*models.ProblemCategory, 0)
	for _, categoryId := range categoryIds {
		atoi, _ := strconv.Atoi(categoryId)
		categoryBasics = append(categoryBasics, &models.ProblemCategory{
			ProblemId:  problems.ID,
			CategoryId: uint(atoi),
		})
	}
	problems.ProblemCategories = categoryBasics

	//处理测试用例
	testCasesBasic := make([]*models.TestCase, 0)
	for _, testCase := range testCases {
		// 举个例子 "input": "1 2\n" "output": "3\n"
		caseMap := map[string]string{}
		err := json.Unmarshal([]byte(testCase), &caseMap)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "测试用例格式错误",
			})
			return
		}
		if _, ok := caseMap["input"]; !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "测试用例输入格式错误",
			})
			return
		}
		if _, ok := caseMap["output"]; !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "测试用例输出格式错误",
			})
			return
		}

		testCasesBasic = append(testCasesBasic, &models.TestCase{
			Identity:        helper.GetUUID(),
			ProblemIdentity: identity,
			Input:           caseMap["input"],
			Output:          caseMap["output"],
		})
	}
	problems.TestCases = testCasesBasic

	//创建问题
	//程序并不会收到mysql返回的主键值？？？
	err := models.DB.Create(&problems).Error
	if err != nil {
		log.Printf("Problem Create Error: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "创建问题失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

// ProblemUpdate
// @Schemes	problem_update
// @Description 修改问题
// @Produce json
// @Tags 管理员私有方法
// @Summary 修改问题
// @Param authorization header string true "token"
// @Param identity formData string true "identity"
// @Param title formData string true "title"
// @Param content formData string true "content"
// @Param max_runtime formData int false "max_runtime"
// @Param max_mem formData int false "max_mem"
// @Param category_ids formData []string false "category_ids" collectionFormat(multi)
// @Param test_cases formData []string true "test_cases" collectionFormat(multi)
// @Success 200 {string} json{"code":"200","msg":""}
// @Router /admin/problem-update [put]
func ProblemUpdate(c *gin.Context) {
	identity := c.PostForm("identity")
	title := c.PostForm("title")
	content := c.PostForm("content")
	maxRuntime, _ := strconv.Atoi(c.PostForm("max_runtime"))
	maxMem, _ := strconv.Atoi(c.PostForm("max_mem"))
	categoryIds := c.PostFormArray("category_ids")
	testCases := c.PostFormArray("test_cases")
	if identity == "" || title == "" || content == "" || len(testCases) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}

	if err := models.DB.Transaction(func(tx *gorm.DB) error {
		// 问题基础信息保存 problem_basic
		problemBasic := &models.ProblemBasic{
			Identity:    identity,
			Title:       title,
			Content:     content,
			Max_Men:     maxMem,
			Max_runtime: maxRuntime,
		}
		err := tx.Where("identity = ?", identity).Updates(problemBasic).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "更新问题失败",
			})
			return err
		}
		// 查询问题详情
		err = tx.Where("identity = ?", identity).Find(&problemBasic).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "查询问题失败",
			})
			return err
		}
		// 关联问题分类的更新
		// 1.删除已存在的关联关系
		err = tx.Where("problem_id = ?", problemBasic.ID).Delete(&models.ProblemCategory{}).Error
		if err != nil {
			return err
		}
		// 2.新增新的关联关系
		pcs := make([]*models.ProblemCategory, 0)
		for _, id := range categoryIds {
			intId, _ := strconv.Atoi(id)
			pcs = append(pcs, &models.ProblemCategory{
				ProblemId:  problemBasic.ID,
				CategoryId: uint(intId),
			})
		}
		err = tx.Create(&pcs).Error
		if err != nil {
			return err
		}
		// 关联测试用例的更新
		// 1.删除已存在的关联关系
		err = tx.Where("problem_identity = ?", identity).Delete(&models.TestCase{}).Error
		if err != nil {
			return err
		}

		// 2.新增新的关联关系
		tcs := make([]*models.TestCase, 0)
		for _, testCase := range testCases {
			caseMap := map[string]string{}
			err := json.Unmarshal([]byte(testCase), &caseMap)
			if err != nil {
				return err
			}
			if _, ok := caseMap["input"]; !ok {
				return err
			}
			if _, ok := caseMap["output"]; !ok {
				return err
			}
			tcs = append(tcs, &models.TestCase{
				Identity:        helper.GetUUID(),
				ProblemIdentity: identity,
				Input:           caseMap["input"],
				Output:          caseMap["output"],
			})
		}

		err = tx.Create(&tcs).Error
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Printf("Problem Update err:%v \n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "更新问题失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
