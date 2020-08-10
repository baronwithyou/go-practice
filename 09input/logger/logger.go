package logger

import (
	"fmt"
	"io"
	"os"
)

// type Type string
// type Message string

type log struct {
	w io.Writer
}

// type logObj struct {
// 	t Type
// 	m Message
// }

var l *log

func init() {
	l = &log{w: os.Stdin}
}

func SetOutput(w io.Writer) {
	l.w = w
}

func Info(v interface{}) {
	l.Info(v)
}

func Error(v interface{}) {
	l.Error(v)
}

func (l *log) Error(v interface{}) {
	fmt.Fprintln(l.w, v)
}

func (l *log) Info(v interface{}) {
	fmt.Fprintln(l.w, v)
}
