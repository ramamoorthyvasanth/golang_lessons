package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	c := &Calculator{10.35, 20}
	if c.Add() != 30 {
		t.Error("OOPS, Addition failed")
	}

}
