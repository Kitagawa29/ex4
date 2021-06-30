package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	b.put(1, 1, "o")
	if b.get(1, 1) != "o" {
		t.Errorf("test1")
	}
}

func TestShow01(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	stdout := os.Stdout
	os.Stdout = w

	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	b.put(1, 1, "o")
	b.put(2, 2, "x")
	b.show()

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	if buf.String() != "...\n.o.\n..x\n" {
		t.Errorf("test2")
	}
}

func TestJudge01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	b.put(0, 0, "o")
	b.put(0, 1, "o")
	b.put(0, 2, "o")
	b.put(1, 0, "x")
	b.put(1, 2, "x")
	b.put(2, 1, "x")
	if b.judge() != "o" {
		t.Errorf("test3")
	}
}

func TestJudge02(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	b.put(0, 0, "x")
	b.put(1, 0, "x")
	b.put(2, 0, "x")
	b.put(0, 2, "o")
	b.put(1, 1, "o")
	b.put(2, 1, "o")
	if b.judge() != "x" {
		t.Errorf("test4")
	}
}

func TestJudge03(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	b.put(0, 0, "o")
	b.put(1, 1, "o")
	b.put(2, 2, "o")
	if b.judge() != "o" {
		t.Errorf("test5")
	}
}
