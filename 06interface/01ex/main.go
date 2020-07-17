package main

// 使用接口的方式实现一个既可以往终端写日志也可以往文件写日志的简易日志哭
import (
	"fmt"
	"os"
)

type Logger interface {
	Log(string)
}

type TerminalLogger struct{}
type FileLogger struct{}

func (t *TerminalLogger) Log(s string) {
	fmt.Println(s)
}

func (f *FileLogger) Log(s string) {
	file, _ := os.OpenFile("./log.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	defer file.Close()

	file.WriteString(s)
}

func main() {

	var logger Logger

	logger = &FileLogger{}

	logger.Log("helloworld")
}
