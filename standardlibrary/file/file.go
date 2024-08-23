package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// io/ioutil 读取文件
	// ioUtilReadFile()

	// 以只读方式打开文件
	// f, err := os.Open("../time/time.go")
	// if err != nil {
	// 	fmt.Printf("open file failed: %v\n", err)
	// 	return
	// }
	// defer f.Close()

	// 一次读取文件
	// var tmp = make([]byte, 128)
	// n, err := f.Read(tmp)
	// if err == io.EOF {
	// 	fmt.Println("文件已经读取完了")
	// 	return
	// }
	// if err != nil {
	// 	fmt.Printf("read file failed: %v\n", err)
	// 	return
	// }
	// fmt.Printf("读取了%d字节数据", n)
	// fmt.Println(string(tmp[:n]))

	// 循环读取文件
	// var content []byte
	// var tmp = make([]byte, 128)
	// for {
	// 	n, err := f.Read(tmp)
	// 	if err == io.EOF {
	// 		fmt.Println("文件已经读取完了")
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Printf("read file failed: %v\n", err)
	// 		return
	// 	}
	// 	content = append(content, tmp[:n]...)
	// }
	// fmt.Println(string(content))

	// bufio 读取文件
	// bufio 是在 file 的基础上封装的一层 API
	// s := bufio.NewScanner(f)
	// for s.Scan() {
	// 	fmt.Println(s.Text())
	// }
	// if err := s.Err(); err != nil {
	// 	fmt.Printf("bufio method read file failed: %v\n", err)
	// }

	// 写入文件
	// writeFile()

	// 实践：将 file.go 文件中的内容读取出来然后写入到 test.txt 中
	rFile, err := os.Open("../time/time.go")
	wFile, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	w := bufio.NewWriter(wFile)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	var tmp = make([]byte, 128)
	for {
		n, err := rFile.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件已经读取完了")
			break
		}
		if err != nil {
			fmt.Printf("read file failed: %v\n", err)
			return
		}
		w.WriteString(string(tmp[:n]))
	}
	w.Flush()
}

func ioUtilReadFile() {
	b, err := ioutil.ReadFile("../time/time.go")
	if err != nil {
		fmt.Printf("read file failed err: %v\n", err)
		return
	}
	fmt.Println(string(b))
}

func writeFile() {
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	defer f.Close()

	// str := "hello 沙河"
	// f.Write([]byte(str))
	// f.WriteString(str)

	// bufio.NewWriter
	w := bufio.NewWriter(f)
	for i := 0; i < 100; i++ {
		// 将数据写入缓存
		w.WriteString("hello 流沙河\n")
	}
	w.Flush()
}
