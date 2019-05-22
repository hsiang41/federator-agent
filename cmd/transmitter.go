package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
)

type TestJob1 struct {
}

func (TestJob1) Run() {
	fmt.Println("test jog1 ...")
}

type TestJob2 struct {
}

func (TestJob2) Run() {
	fmt.Println("test job2 ...")
}

func main() {
	i := 0
	c := cron.New()

	// Add Func
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})

	// Add job{}
	c.AddJob(spec, TestJob1{})
	c.AddJob(spec, TestJob2{})

	c.Start()

	defer c.Stop()

	select {}
}
