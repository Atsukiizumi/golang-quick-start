package main

type CheckCar interface {
	Run()
	Stop()
}

type Car struct {
	Brand string
	Color string
}

type truck any
