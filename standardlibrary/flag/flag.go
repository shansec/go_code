package main

import (
	"flag"
	"fmt"
)

func main() {
	// flag.Type
	//name := flag.String("name", "刘备", "姓名")
	//age := flag.Int("age", 18, "年龄")

	// flag.TypeVar
	//var name string
	//var age int
	//flag.StringVar(&name, "name", "刘备", "姓名")
	//flag.IntVar(&age, "age", 18, "年龄")

	//flag.Arg() 返回命令行参数后的其他参数，以[]string类型
	//flag.NArg() 返回命令行参数后的其它参数

	var name string
	var age int
	flag.StringVar(&name, "name", "刘备", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")

	flag.Parse()
	fmt.Println(name, age)

	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
