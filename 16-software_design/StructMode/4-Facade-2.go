package main

import (
	"fmt"
)

// 电视机
type TV struct {
}

func (t *TV) On() {
	fmt.Println("TV On")
}

func (t *TV) Off() {
	fmt.Println("TV Off")
}

// 音响
type VoiceBox struct {
}

func (vb *VoiceBox) On() {
	fmt.Println("VoiceBox On")
}

func (vb *VoiceBox) Off() {
	fmt.Println("VoiceBox Off")
}

// 灯光
type Light struct {
}

func (l *Light) On() {
	fmt.Println("Light On")
}

func (l *Light) Off() {
	fmt.Println("Light Off")
}

// 游戏机
type Playstation5 struct {
}

func (p *Playstation5) On() {
	fmt.Println("Playstation5 On")
}

func (p *Playstation5) Off() {
	fmt.Println("Playstation5 Off")
}

type MicroPhone struct {
}

func (m *MicroPhone) On() {
	fmt.Println("MicroPhone On")
}

func (m *MicroPhone) Off() {
	fmt.Println("MicroPhone Off")
}

// 投影仪
type Projector struct {
}

func (p *Projector) On() {
	fmt.Println("Projector On")
}

func (p *Projector) Off() {
	fmt.Println("Projector Off")
}

// 家庭影院（外观）
type HomePlayerFacade struct {
	tv          TV
	vb          VoiceBox
	light       Light
	playstation Playstation5
	mp          MicroPhone
	projector   Projector
}

// KTV模式
func (hp *HomePlayerFacade) DoKTV() {
	fmt.Println("进入KTV模式...")
	hp.tv.On()
	hp.vb.On()
	hp.light.Off()
	hp.playstation.Off()
	hp.mp.On()
	hp.projector.On()
}

// 游戏模式
func (hp *HomePlayerFacade) DoGame() {
	fmt.Println("进入游戏模式...")
	hp.tv.On()
	hp.light.On()
	hp.playstation.On()
}

func main() {
	homePlayerFacade := &HomePlayerFacade{}
	homePlayerFacade.DoKTV()

	fmt.Println("-------------")
	homePlayerFacade.DoGame()
}
