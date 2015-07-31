package main

import (
	"fmt"
)

type Calculator struct {
	x, y int
}

func (c *Calculator) Multiply() int {
	return c.x * c.y
}

func (c *Calculator) Divide() int {
	return c.x / c.y
}

func (c *Calculator) Add() int {
	return c.x + c.y
}

func (c *Calculator) Subtract() int {
	return c.x - c.y
}
func (values ...int) adds int {
	total := 0
	_,value  := range values {
		total += value
	}
	return total
}

func main() {
	c := &Calculator{10, 80}
	fmt.Println(c.Add())

}
