package main

import (
	"testing"
)

func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 1, "o")
	if b.get(1, 1) != "o" {
		t.Errorf("....")
	}
}

func TestDisplay01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 1, "o")
	if b.get(1, 1) != "o" {
		t.Errorf("....")
	}
	b.put(1, 2, "x")
	if b.get(1, 2) != "x" {
		t.Errorf("....")
	}
	b.put(2, 1, "o")
	if b.get(2, 1) != "o" {
		t.Errorf("....")
	}
	b.display()
}

func TestDrawDisplay01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "o")
	if b.get(0, 0) != "o" {
		t.Errorf("....")
	}
	b.put(0, 1, "x")
	if b.get(0, 1) != "x" {
		t.Errorf("....")
	}
	b.put(0, 2, "o")
	if b.get(0, 2) != "o" {
		t.Errorf("....")
	}
	b.put(1, 0, "x")
	if b.get(1, 0) != "x" {
		t.Errorf("....")
	}
	b.put(1, 1, "o")
	if b.get(1, 1) != "o" {
		t.Errorf("....")
	}
	b.put(1, 2, "x")
	if b.get(1, 2) != "x" {
		t.Errorf("....")
	}
	b.put(2, 0, "o")
	if b.get(2, 0) != "o" {
		t.Errorf("....")
	}
	b.put(2, 1, "x")
	if b.get(2, 1) != "x" {
		t.Errorf("....")
	}
	b.put(2, 2, "o")
	if b.get(2, 2) != "o" {
		t.Errorf("....")
	}
	b.display()
	if b.judge() != 3 {
		t.Errorf("JudgeError")
	}
}

func TestJudge01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "o")
	if b.get(0, 0) != "o" {
		t.Errorf("....")
	}
	b.put(1, 1, "o")
	if b.get(1, 1) != "o" {
		t.Errorf("....")
	}
	b.put(2, 2, "o")
	if b.get(2, 2) != "o" {
		t.Errorf("....")
	}
	b.display()
	if b.judge() != 1 {
		t.Errorf("JudgeError")
	}
}

func TestJudge02(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "o")
	if b.get(0, 0) != "o" {
		t.Errorf("....")
	}
	b.put(0, 1, "o")
	if b.get(0, 1) != "o" {
		t.Errorf("....")
	}
	b.put(0, 2, "o")
	if b.get(0, 2) != "o" {
		t.Errorf("....")
	}
	b.display()
	if b.judge() != 1 {
		t.Errorf("JudgeError")
	}
}

func TestJudge03(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "o")
	if b.get(0, 0) != "o" {
		t.Errorf("....")
	}
	b.put(1, 0, "o")
	if b.get(1, 0) != "o" {
		t.Errorf("....")
	}
	b.put(2, 0, "o")
	if b.get(2, 0) != "o" {
		t.Errorf("....")
	}
	b.display()
	if b.judge() != 1 {
		t.Errorf("JudgeError")
	}
}
