package main

import "fmt"

type Vehicle interface {
	Start()
	Stop()
}

type BMW struct {
	id     int
	engine string
}

func (b BMW) Start() {
	fmt.Println("BMW start")
}

func (b BMW) Stop() {
	fmt.Println("BMW stop")
}

type Mercedes struct {
	id     int
	engine string
}

func (m Mercedes) Start() {
	fmt.Println("Mercedes start")
}

func (m Mercedes) Stop() {
	fmt.Println("Mercedes stop")
}

func main() {
	bmw := BMW{1, "BMW"}
	mercedes := Mercedes{1, "Mercedes"}

	Vehicle.Start(bmw)
	Vehicle.Start(mercedes)
}
