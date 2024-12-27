package main

import "fmt"

// 医生-命令接收者
type Doctor struct {
}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼镜")
}

func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// 抽象命令
type Command interface {
	Treat()
}

// 治疗眼睛的病单
type CommandTreatEye struct {
	doctor *Doctor
}

func (c *CommandTreatEye) Treat() {
	c.doctor.treatEye()
}

// 治疗鼻子的病单
type CommandTreatNose struct {
	doctor *Doctor
}

func (c *CommandTreatNose) Treat() {
	c.doctor.treatNose()
}

// 护士-调用命令者
type Nurse struct {
	CmdList []Command //收集的命令集合
}

func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}

	for _, cmd := range n.CmdList {
		cmd.Treat() //执行病单绑定的命令(这里会调用病单已经绑定的医生的诊断方法)
	}
}

func main() {
	//依赖病单，通过填写病单，让医生看病
	doctor := &Doctor{}
	//治疗眼镜的病单
	cmdEye := &CommandTreatEye{doctor}
	//治疗鼻子的病单
	cmdNose := &CommandTreatNose{doctor}

	//护士
	nurse := new(Nurse)
	//收集管理病单

	nurse.CmdList = append(nurse.CmdList, cmdEye)
	nurse.CmdList = append(nurse.CmdList, cmdNose)

	//执行病单指令
	nurse.Notify()
}
