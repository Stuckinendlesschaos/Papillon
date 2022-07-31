package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"papillon/conf"
	"papillon/event"
	"papillon/event/merge"
	"papillon/event/push"
	"papillon/event/release"
	"papillon/event/tag"
	"papillon/output"
	"papillon/output/robot_type"

	"github.com/gin-gonic/gin"
)

func handleEvent(EventType string, body *[]byte) error {
	var (
		msg *event.MessageNotify
		err error
	)

	//根据请求头的事件获取内容的不同解析
	//handles body at embedded handle method sorted by the kind of events
	switch EventType {
	case "Push Hook":
		msg, err = push.Handle(body)
	case "Merge Request Hook":
		msg, err = merge.Handle(body)
	case "Tag Push Hook":
		msg, err = tag.Handle(body)
	case "Release Hook":
		msg, err = release.Handle(body)
	}

	if err != nil {
		return fmt.Errorf("Grab gitlab event has failed,err is:%w", err)
	}

	//根据不同的终端，用不同的机器人发送到终端
	//send to destination after handling body
	var terminal output.Terminal

	terminal = &robot_type.WecomRobot{}
	new_robot := terminal.CreateRobot(msg)

	err = new_robot.SendSpecifiedAddress()
	if err != nil {
		return fmt.Errorf("send embedded address has happened error: %w", err)
	}
	return nil
}

// handleGitlab handles the embedded request from gitlab
// 处理具体请求的句柄
func handleGitlab(c *gin.Context) {
	//从gitlab获取不同的事件消息
	//obtains different events on POST request header when github incurs specified event
	gitLabEvent := c.Request.Header.Get("X-Gitlab-Event")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalf("read message has failed, err is:%s", err)
	}

	go handleEvent(gitLabEvent, &body)
	c.String(200, "Success")
}

func main() {
	//read config file path via flag
	//flag命令行工具解析配置文件路径
	configFile := flag.String("c", "./backend.yaml", "the necessary config file for launching the whole programe")
	flag.Parse()

	//加载backend.yaml文件和加载userid.csv文件
	//flag cli loads specfied configure
	conf.LoadConfig(*configFile)
	//loads content at userid.csv into map for using
	conf.LoadUserData()

	//使用gin web框架在本地服务做个转发代理
	//gin web accepts POST method for receiving msg from gitlab
	gin.DefaultWriter = io.MultiWriter(os.Stdout)

	route := gin.Default()
	route.POST("/gitlab", handleGitlab)

	route.Run(conf.Config.Addr)
}
