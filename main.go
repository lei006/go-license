package main

import (
	"fmt"
	"time"

	"github.com/lei006/go-daemon/daemontool"
)

func main() {

	daemonTool := daemontool.DefDaemonTool
	/*
		daemonTool.OnInit("tet_app", "测试的...")
		daemonTool.SetModel(false)
		daemonTool.ServerRun(func() {
			Run()
		})
	*/
	daemonTool.Run("test_app11", "desc 111测试333", Run)
}

func Run() {

	for true {
		fmt.Println("xxxxxxx", daemontool.DefDaemonTool.GetWordPath())
		time.Sleep(time.Second)
	}

}
