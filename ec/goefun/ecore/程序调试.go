package ecore

import "fmt"
import "github.com/kr/pretty"

func E调试输出(a ...interface{}) {
	pretty.Println(a...)
}

func E调试输出格式化(s string, a ...interface{}) {
	fmt.Printf(s, a...)
}
