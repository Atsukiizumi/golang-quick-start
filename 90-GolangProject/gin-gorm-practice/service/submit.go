package service

import (
	"GolangProject/gin-gorm-practice/define"
	"GolangProject/gin-gorm-practice/helper"
	"GolangProject/gin-gorm-practice/models"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
	"time"
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

// Submit
// @Schemes	test_case
// @Description 代码提交
// @Produce json
// @Tags 用户私有方法
// @Summary 代码提交
// @Param authorization header string true "token"
// @Param problem_identity query string true "problem_identity"
// @Param code body string true "code"
// @Success 200 {string} json{"code":"200","msg":""}
// @Router /user/submit [post]
func Submit(c *gin.Context) {
	problemIdentity := c.Query("problem_identity")
	code, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Submit ioutil.ReadAll Error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "读取代码失败",
		})
		return
	}

	u, _ := c.Get("user")
	userClaim := u.(*helper.UserClaims)

	//保存代码
	path, err := helper.CodeSave(code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "代码保存失败",
		})
		return
	}

	sb := &models.SubmitBasic{
		Identity:        helper.GetUUID(),
		ProblemIdentity: problemIdentity,
		UserIdentity:    userClaim.Identity,
		Path:            path,
	}

	// 代码判断
	pb := new(models.ProblemBasic)
	err = models.DB.Where("identity = ?", problemIdentity).Preload("TestCases").First(pb).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "题目不存在",
		})
		return
	}

	// 答案错误
	WA := make(chan int)
	// 运行超内存
	OOM := make(chan int)
	// 编译错误
	CE := make(chan int)
	// 通过的个数
	passCount := 0
	// 提示信息
	msg := ""

	var lock sync.Mutex

	for _, testCase := range pb.TestCases {
		testCase := testCase
		go func() {
			// 执行代码
			cmd := exec.Command("go", "run", "code-user/main.go")
			var out, stderr bytes.Buffer
			cmd.Stderr = &stderr
			cmd.Stdout = &out
			stdinPipe, err := cmd.StdinPipe()
			if err != nil {
				log.Fatalln(err)
			}
			io.WriteString(stdinPipe, testCase.Input)
			var bm runtime.MemStats
			runtime.ReadMemStats(&bm)

			// 根据测试的输入案例，进行运行，拿到输出结果和标准的输出的结果进行比对
			if err := cmd.Run(); err != nil {
				log.Println("err:", stderr.String())
				if err.Error() == "exit status 2" {
					msg = stderr.String()
					CE <- 1
					return
				}
			}
			var em runtime.MemStats
			runtime.ReadMemStats(&em)

			// 答案错误
			if testCase.Output != out.String() {
				msg = "答案错误"
				WA <- 1
				return
			}

			// 运行超内存
			if em.Alloc/1024-(bm.Alloc/1024) > uint64(pb.Max_Men) {
				msg = "运行超内存"
				OOM <- 1
				return
			}

			// 通过
			lock.Lock()
			passCount += 1

			lock.Unlock()

			println("err:", string(stderr.Bytes()))
			log.Println(out.String())
		}()
	}
	select {
	// -1:待判断 0:未执行 1:答案正确 2:答案错误 3:运行超时 4:运行超内存 5:编译错误
	case <-WA:
		sb.Status = 2
	case <-OOM:
		sb.Status = 4
	case <-time.After(time.Millisecond * time.Duration(pb.Max_runtime)):
		if passCount == len(pb.TestCases) {
			sb.Status = 1
		} else {
			sb.Status = 3
		}
	case <-CE:
		sb.Status = 5
	}

	err = models.DB.Create(sb).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "提交失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "提交成功",
		"data": map[string]interface{}{
			"status": sb.Status,
			"msg":    msg,
		},
	})
}
