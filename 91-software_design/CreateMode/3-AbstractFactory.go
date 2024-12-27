package main

import (
	"fmt"
)

// === 抽象层 ===
type AbstractToy interface {
	ShowToy()
}

type AbstractCar interface {
	ShowCar()
}

type AbstractTV interface {
	ShowTV()
}

// 抽象工厂
type AbstractFactory interface {
	CreateToy() AbstractToy
	CreateCar() AbstractCar
	CreateTV() AbstractTV
}

// == 实现层 ==
type ChinaProduct struct{}

func (ca *ChinaProduct) ShowToy() {
	fmt.Println("Toy made in China")
}

func (ca *ChinaProduct) ShowCar() {
	fmt.Println("Car made in China")
}

func (ca *ChinaProduct) ShowTV() {
	fmt.Println("TV made in China")
}

type ChinaFactory struct{}

func (cf *ChinaFactory) CreateToy() AbstractToy {
	var toy AbstractToy
	toy = new(ChinaProduct)

	return toy
}

func (cf *ChinaFactory) CreateCar() AbstractCar {
	var car AbstractCar
	car = new(ChinaProduct)

	return car
}

func (cf *ChinaFactory) CreateTV() AbstractTV {
	var tv AbstractTV
	tv = new(ChinaProduct)

	return tv
}

type ItalyProduct struct{}

func (ci *ItalyProduct) ShowCar() {
	fmt.Println("Car made in Italy")
}

func (ci *ItalyProduct) ShowTV() {
	fmt.Println("TV made in Italy")
}

func (ci *ItalyProduct) ShowToy() {
	fmt.Println("Toy made in Italy")
}

type ItalyFactory struct{}

func (xf *ItalyFactory) CreateToy() AbstractToy {
	var toy AbstractToy
	toy = new(ItalyProduct)

	return toy
}

type IndiaProduct struct{}

func (ci *IndiaProduct) ShowCar() {
	fmt.Println("Car made in India")
}
func (ci *IndiaProduct) ShowTV() {
	fmt.Println("TV made in India")
}

func (ci *IndiaProduct) ShowToy() {
	fmt.Println("Toy made in India")
}

func main() {
	chinaF := new(ChinaFactory)
	ctoy := chinaF.CreateToy()
	ccar := chinaF.CreateCar()
	ctv := chinaF.CreateTV()

	ctv.ShowTV()
	ccar.ShowCar()
	ctoy.ShowToy()

}
