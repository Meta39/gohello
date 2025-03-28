package gohello

import "fmt"

func GoHello() {
	fmt.Println("GoHello")
}

// SayHi 向指定人打招呼的函数
func SayHi(name string) {
	fmt.Printf("你好%s，我是Meta39。很高兴认识你。\n", name)
}
