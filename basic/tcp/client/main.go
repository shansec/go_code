package main

import (
	"fmt"
	"net"
)

// func main() {
// 	conn, err := net.Dial("tcp", "127.0.0.1:8080")
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// 	defer conn.Close()
// 	inputReader := bufio.NewReader(os.Stdin)
// 	for {
// 		input, _ := inputReader.ReadString('\n')
// 		inputInfo := strings.Trim(input, "\r\n")
// 		if strings.ToUpper(inputInfo) == "Q" {
// 			return
// 		}
// 		_, err := conn.Write([]byte(inputInfo))
// 		if err != nil {
// 			return
// 		}
// 		buff := [512]byte{}
// 		n, err := conn.Read(buff[:])
// 		if err != nil {
// 			fmt.Println("recv failed, err:", err)
// 			return
// 		}
// 		fmt.Println(string(buff[:n]))
// 	}
// }

// tcp 黏包案例
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	for i := 0; i < 10; i++ {
		msg := "Hello, Hello, How are you ?"
		conn.Write([]byte(msg))
	}
}
