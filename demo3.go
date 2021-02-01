package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-01 12:32
* @Description:
**/

type MyJob struct {
	ID int
}

func (job MyJob) Run(){
	fmt.Println("hello", job.ID)
}

func main() {
	job1 := MyJob{1}
	job2 := MyJob{2}
	job3 := MyJob{3}
	job4 := MyJob{4}
	c := cron.New()
	if _, err := c.AddJob("@every 1s", job1); err != nil{
		fmt.Println(err)
	}
	if _, err := c.AddJob("@every 2s", job2); err != nil{
		fmt.Println(err)
	}
	if _, err := c.AddJob("@every 3s", job3); err != nil{
		fmt.Println(err)
	}
	if _, err := c.AddJob("@every 4s", job4); err != nil{
		fmt.Println(err)
	}
	c.Start()
	t1 := time.NewTimer(time.Minute * 2)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Minute * 2)
		}
	}
}