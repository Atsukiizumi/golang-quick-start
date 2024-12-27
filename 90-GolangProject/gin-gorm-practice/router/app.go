package router

import (
	_ "GolangProject/gin-gorm-practice/docs"
	"GolangProject/gin-gorm-practice/middleware"
	"GolangProject/gin-gorm-practice/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	//Swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 配置路由规则

	// 公有方法
	// 问题
	r.GET("/problem-list", service.GetProblemList)
	r.GET("/problem-detail", service.GetProblemDetail)

	// 用户
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/login", service.Login)
	r.POST("/send-code", service.SendCode)
	r.POST("/register", service.Register)
	r.GET("/rank-list", service.GetRankList)

	// 提交记录
	r.GET("/submit-list", service.GetSubmitList)

	//管理员私有方法
	// 管理员分组
	authAdmin := r.Group("/admin", middleware.AuthAdminCheck())
	// 创建问题
	authAdmin.POST("/problem-create", service.ProblemCreate)
	// 修改问题
	authAdmin.PUT("/problem-update", service.ProblemUpdate)
	// 分类列表
	authAdmin.GET("/category-list", service.GetCategoryList)
	// 创建分类
	authAdmin.POST("/category-create", service.CategoryCreate)
	// 删除分类
	authAdmin.DELETE("/category-delete", service.CategoryDelete)
	// 修改分类
	authAdmin.PUT("/category-update", service.CategoryUpdate)

	//用户私有方法
	authUser := r.Group("/user", middleware.AuthUserCheck())
	authUser.POST("/submit", service.Submit)
	return r
}
