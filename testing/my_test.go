package test

import (
	"fmt"
	"os"
	"testing"
)

func TestT1(t *testing.T) {
	fmt.Println("t1")
	//t.Error("test TestT1 sadly failed")
}

func TestT2(t *testing.T) {
	fmt.Println("t2")
	//t.Error("test TestT2 sadly failed")
}

func TestT3(t *testing.T) {
	fmt.Println("t3")
	//t.Error("test TestT3 sadly failed")
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}