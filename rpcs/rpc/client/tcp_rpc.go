package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	X, Y int
}

func main() {
	dial, err := rpc.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Printf("dial error: %v\n", err)
		return
	}
	// 同步调用
	args := Args{10, 20}
	var reply int
	err = dial.Call("ServiceA.Add", args, &reply)
	if err != nil {
		fmt.Printf("call error: %v\n", err)
	}
	fmt.Printf("reply: %d\n", reply)
}
