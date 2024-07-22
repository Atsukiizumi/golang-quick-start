package main

import "fmt"

// 抽象主题
type BeautyWoman interface {
	//对男人抛媚眼
	MakeEyesWithMan()
	//和男人浪漫约会
	RomanDateWithMan()
}

// 具体主题
type Women struct {
}

func (w *Women) MakeEyesWithMan() {
	fmt.Println("给男人抛了个媚眼")
}
func (w *Women) RomanDateWithMan() {
	fmt.Println("与男人共度了浪漫的约会")
}

type Matchmaker struct {
	women BeautyWoman
}

func NewProxy_2(women BeautyWoman) BeautyWoman {
	return &Matchmaker{women}
}

func (m *Matchmaker) MakeEyesWithMan() {
	m.women.MakeEyesWithMan()
}

func (m *Matchmaker) RomanDateWithMan() {
	m.women.RomanDateWithMan()
}

func main() {
	matchmaker := NewProxy_2(new(Women))
	matchmaker.RomanDateWithMan()
	matchmaker.MakeEyesWithMan()
}
