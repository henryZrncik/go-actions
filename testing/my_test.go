package test

import (
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"log"
	"os"
	"testing"
)

func TestT1(t *testing.T) {
	t.Log("Test zacina")
	var err error
	fmt.Println("t1")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path := dir + "/compose-kafka-zookeeper.yml"
	var tc = testcontainers.NewLocalDockerCompose([]string{path}, "someName2")
	var res = tc.WithCommand([]string{"up", "-d"}).Invoke()
	if res.Error != nil {
		log.Println("DORITI")
		fmt.Println("do you feel it")
		t.Error("container didn't start")
	}
	log.Println("docker compose networka created")

	tc.Down()
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