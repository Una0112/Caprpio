package main

import (
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"strings"
	"log"
	"os"
	"bufio"
	"fmt"
	"time"
)

//go:generate cqcfg -c .
// cqp: 名称:教学助手机器人
// cqp: 版本: 1.0.4:5
// cqp: 作者: zyj
// cqp: 简介: ddl提醒机器

var (
	newFile *os.File
	err     error
)

func main() {
	newFile, err = os.Create("D:/test.txt")
	if err!=nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}

var groups []int64
var marks []int64

func init() {
	cqp.AppID = "bot.ddl.helpers"
	cqp.GroupMsg = OnGroupMsg
	cqp.PrivateMsg= onPrivateMsg
	groups=append(groups,942315244)
	groups=append(groups,829856110)
}

func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	num:=1;
	for {
		ticker:=time.NewTicker(time.Second);
		num=num+1
		if num > 360000000 {
			ticker.Stop();
			break;
		}
		t2 := int(time.Now().Month())
		t3 := time.Now().Day()
		t4 := time.Now().Hour()
		t5 := time.Now().Minute()
		t6 := time.Now().Second()
		file, err := os.Open("D:/test.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		i := 0
		for {
			i = i + 1
			str, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			if marks[i]!= 1 {
				a := int(str[5]*10 + str[6])
				b := int(str[8]*10 + str[9])
				c := int(str[11]*10 + str[12])
				d := int(str[14]*10 + str[15])
				e := int(str[17]*10 + str[18])
				if a == t2 && b == t3 && c == t4 && d == t5 && e == t6 {
					str := "NOTICE:" + str
					cqp.SendPrivateMsg(fromQQ, str)
				}
			}
		}
		if strings.Contains(msg, "添加") {
			cqp.SendPrivateMsg(fromQQ, "请按以下格式输入：2020-01-01,20:00:00,name 其中标点为英文格式")
			return 0
		}

		if strings.Contains(msg, "2020") {
			fileName := "D:/test"
			file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
			if err != nil {
				fmt.Println("error", err)
				os.Exit(1)
			}
			defer file.Close()
			file.Seek(0, 2)
			file.WriteString(msg)
			return 0
		}

		if strings.Contains(msg, "删除") {
			file, err := os.Open("D:/test.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			reader := bufio.NewReader(file)
			i := 0
			t:=0
			for {
				i = i + 1
				t=t+1
				str, err := reader.ReadString('\n')
				if err != nil {
					break
				}
				if marks[i]!= 1  {
					cqp.SendPrivateMsg(fromQQ,str)
				}
				if marks[i]==1{
					t=t-1
				}
			}
			cqp.SendPrivateMsg(fromQQ, "请输入删除+任务编号")
			return 0
		}

		if strings.Contains(msg, "删除+") {
			n := int(msg[2] - '0')
			marks[n]=1
		}

		if strings.Contains(msg, "查询") {
			file, err := os.Open("D:/test.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			reader := bufio.NewReader(file)
			i := 0
			t:=0
			for {
				i = i + 1
				t=t+1
				str, err := reader.ReadString('\n')
				if err != nil {
					return 0
				}
				if marks[i]!= 1  {
					cqp.SendPrivateMsg(fromQQ,str)
					return 0
				}
				if marks[i]==1{
					t=t-1
					return 0
				}
			}
			return 0
		}
	}

	return 0
}

func OnGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32)int32 {
	if strings.Contains(msg,"晚安") {
		cqp.SendGroupMsg(fromGroup,"good night")
		return 0
	}
	return 0
}
