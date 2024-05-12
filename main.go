package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string //用户发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

// 保存在线用户 cliAddr ===> Client
var OnLineMap map[string]Client

// 通信的管道
var message = make(chan string)

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "{" + cli.Addr + "}" + cli.Name + ": " + msg
	return buf
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		msgs := msg + "\n"
		conn.Write([]byte(msgs))
	}

}

func HandlerConn(conn net.Conn) {
	//获取客户端网络地址
	cliAddr := conn.RemoteAddr().String()
	defer conn.Close()

	//创建一个结构体
	cli := Client{
		C:    make(chan string),
		Name: cliAddr,
		Addr: cliAddr,
	}

	//把结构体加入到map
	OnLineMap[cliAddr] = cli

	//新开一个协程，专门给当前客户端发送信息
	go WriteMsgToClient(cli, conn)

	//广播某个人在线
	//message <- "{" + cli.Addr + "}" + cli.Name + ": login"
	message <- MakeMsg(cli, "login\n")

	//提示我是谁
	cli.C <- MakeMsg(cli, "I am here\n")

	isQuit := make(chan bool)  //对方是否主动退出
	hasData := make(chan bool) //对方是否有数据发送

	//新开一个协程，接受用户发送过来的数据
	go func() {
		buf := make([]byte, 2048)

		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println("conn.Read = err", err)
				return
			}

			msg := string(buf[:n-1])

			//过滤数据
			if len(msg) == 3 && msg == "who" {
				//遍历map，给当前用户发送所有成员
				conn.Write([]byte("user list:\n"))
				for _, tmp := range OnLineMap {
					msg := tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				//"rename|mike"
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				OnLineMap[cliAddr] = cli
				conn.Write([]byte("rename OK" + "\n"))
			} else {
				message <- MakeMsg(cli, msg)
			}

			hasData <- true //代表有数据
		}

	}()

	for {
		//通过select检测这个channel的流动
		select {
		case <-isQuit:
			delete(OnLineMap, cliAddr)           //当前用户从map移除
			message <- MakeMsg(cli, "login out") //广播谁下线了
			return
		case <-hasData:

		case <-time.After(60 * time.Second):
			delete(OnLineMap, cliAddr)
			message <- MakeMsg(cli, "time out leave out")
			return
		}
	}
}

func Manager() {
	//给map分配空间
	OnLineMap = make(map[string]Client)

	for {
		msg := <-message //没有消息前，这里会阻塞
		//遍历map
		for _, cli := range OnLineMap {
			cli.C <- msg
		}
	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer listener.Close()

	//新开一个协程，转发消息,只要有消息就遍历map，给map每个成员都发送消息
	go Manager()

	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("accept err1 = ", err1)
			continue
		}

		go HandlerConn(conn) //处理用户连接的协程

	}

}
