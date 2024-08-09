package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Account struct {
	family *Family
	flag   int
}

func (this *Account) menu() bool {
	var flag int
	fmt.Println("-------家庭收支记账软件-------")
	fmt.Println("	1.收支明细")
	fmt.Println("	2.登记收入")
	fmt.Println("	3.登记支出")
	fmt.Println("	4.退出软件")
	fmt.Print("请选择（1-4）：")
	fmt.Scanln(&flag)

	if flag >= 1 && flag <= 4 {
		this.flag = flag
		return true
	} else {
		fmt.Println(">>>>>> 请输入有效范围 <<<<<<")
		return false
	}
}

func (this *Account) accountingLog() {
	fmt.Println("-------当前收支明细记录-------")
	fmt.Println("收支\t账户金额\t收支金额\t说明")
	if len(this.family.Billing) == 0 {
		fmt.Println(">>>>>> 暂无数据 <<<<<<")
		return
	}
	for _, v := range this.family.Billing {
		fmt.Printf("%s\t%.2f\t%.2f\t\t%s\n", v.AccType, v.Accbalance, v.RevenueNum, v.Comment)
	}
}

func (this *Account) revenue() {
	var rate float64
	var comment string
	//TODO 判断输入金额
	fmt.Print("本次收入金额：")
	fmt.Scanln(&rate)
	fmt.Print("本次收入说明：")
	inputreader := bufio.NewReader(os.Stdin)
	input, err := inputreader.ReadString('\n')
	if err != nil {
		fmt.Println("Some errors reading input...")
		return
	}

	comment = input

	this.family.AddBalance("收入", rate, comment)

}

func (this *Account) expenditures() {
	var rate float64
	var comment string
	//TODO 判断输入金额
	fmt.Println("本次支出金额：")
	fmt.Scanln(&rate)
	//TODO 判断是否有余额？
	fmt.Println("本次支出说明：")
	inputreader := bufio.NewReader(os.Stdin)
	input, err := inputreader.ReadString('\n')
	if err != nil {
		fmt.Println("Some errors reading input...")
		return
	}

	comment = input

	this.family.AddBalance("支出", rate, comment)
}

func (this *Account) Run() {
	for this.flag != 4 {
		for this.menu() != true {
		}

		switch this.flag {
		case 1:
			this.accountingLog()
			break
		case 2:
			this.revenue()
			break
		case 3:
			this.expenditures()
			break
		case 4:
			//添加确认退出逻辑
			flag := false
			var yon string
			for flag == false {
				fmt.Println("确认要退出本程序吗？ y(Y)/n(N)")
				fmt.Scanln(&yon)
				switch strings.ToUpper(yon) {
				case "Y":
					fmt.Println("感谢使用！")
					flag = true
					break
				case "N":
					flag = true
					this.flag = 0
					break
				default:
					fmt.Println(">>>>>> 请输入正确的指令 <<<<<<")
					flag = false
					break
				}
			}
			break
		}
	}
}

func NewAccount() *Account {
	nacc := &Account{
		family: NewFamily(),
	}

	return nacc
}

func main() {
	nacc := NewAccount()
	nacc.Run()
}
