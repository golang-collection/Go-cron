package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-01 11:24
* @Description:
**/

func main() {
	c := cron.New()
	if id, err := c.AddFunc("*/1 * * * *", func(){
		fmt.Println("hello")
	}); err != nil{
		fmt.Println(id, err)
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