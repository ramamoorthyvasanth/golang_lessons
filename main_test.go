package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	c := &Calculator{10, 20}
	if c.Add() != 30 {
		t.Error("OOPS, Addition failed")
	}

	if c.Multiply() != 200 {
		t.Error("OOPS, Addition failed")
	}

}
