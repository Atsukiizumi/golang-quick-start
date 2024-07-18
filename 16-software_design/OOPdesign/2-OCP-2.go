package main

//开闭原则设计
import "fmt"

// 抽象的银行业务员
type AbstractBanker interface {
	//抽象的处理业务接口
	DoBusi()
}

type SaveBanker struct {
}

func (sa *SaveBanker) DoBusi() {
	fmt.Println("进行了存款")
}

type TransferBanker struct {
}

func (ta *TransferBanker) DoBusi() {
	fmt.Println("进行了转账")
}

type PayBanker struct {
}

func (pb *PayBanker) DoBusi() {
	fmt.Println("进行了支付")
}

// 实现架构层(基于抽象层进行业务封装-针对interface接口进行封装)
func BankerBussiness(banker AbstractBanker) {
	//通过接口来向下调用(多态)
	banker.DoBusi()
}

//开闭原则
/*
再看开闭原则定义:
开闭原则:一个软件实体如类、模块和函数应该对扩展开放，对修改关闭。
简单的说就是在修改需求的时候，应该尽量通过扩展来实现变化，而不是通过修改已有代码来实现变化。
*/
func main() {
	/*sb := SaveBanker{}
	sb.DoBusi()

	tb := TransferBanker{}
	tb.DoBusi()

	pb := PayBanker{}
	pb.DoBusi()*/

	BankerBussiness(&SaveBanker{})
	BankerBussiness(&TransferBanker{})
	BankerBussiness(&PayBanker{})
}
