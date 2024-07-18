package main

import "fmt"

type Car interface {
	Run()
}

type Driver interface {
	Driver(car Car)
}

type BenZ struct {
}

func (benz *BenZ) Run() {
	fmt.Println("BenZ is running...")
}

type Bmw struct {
}

func (bmw *Bmw) Run() {
	print(fmt.Println("Bmw is running..."))
}

type Zhang struct {
}

func (zhang *Zhang) Driver(car Car) {
	fmt.Println("Zhang is running...")
	car.Run()
}

type Li struct {
}

func (li *Li) Driver(car Car) {
	fmt.Println("Li is running...")
	car.Run()
}

func main() {
	var bwm Car
	bwm = new(BenZ)

	var zhang Driver
	zhang = &Zhang{}
	zhang.Driver(bwm)

	var benz Car
	benz = new(BenZ)

	var li Driver
	li = &Li{}
	li.Driver(benz)
}
