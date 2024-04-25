package main

import (
	"bytes"
	"github.com/fsnotify/fsnotify"
	"log"
)

type Event struct {
	Name string
	Op   fsnotify.Op
}
type Op uint8

const (
	Create Op = 1 << iota
	Remove
	Write
	Rename
	Chmod
)

// fsnotify.Create : 表示创建了一个文件或目录
// fsnotify.Write : 表示修改了一个文件或目录
// fsnotify.Remove : 表示删除了一个文件或目录
// fsnotify.Rename : 表示重命名了一个文件或目录
// fsnotify.Chmod : 表示修改了一个文件或目录的权限

// fsnotify.go
func (op Op) String() string {
	// Use a buffer for efficient string concatenation
	var buffer bytes.Buffer
	if op&Create == Create {
		buffer.WriteString("|CREATE")
	}
	if op&Remove == Remove {
		buffer.WriteString("|REMOVE")
	}
	if op&Write == Write {
		buffer.WriteString("|WRITE")
	}
	if op&Rename == Rename {
		buffer.WriteString("|RENAME")
	}
	if op&Chmod == Chmod {
		buffer.WriteString("|CHMOD")
	}
	if buffer.Len() == 0 {
		return ""
	}
	return buffer.String()[1:] // Strip leading pipe
}

func main() {

	//创建一个fsnotify的监听器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	//添加想要添加监听的文件或目录
	err = watcher.Add("config.json")
	if err != nil {
		log.Fatal(err)
	}

	//启动一个循环，来处理监听器接收到的事件
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event", event)
				//判断事件是否为文件写入操作
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	<-done
}
