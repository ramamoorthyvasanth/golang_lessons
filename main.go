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

func main() {
	c := &Calculator{10, 80}
	fmt.Println(c.Add())

}
