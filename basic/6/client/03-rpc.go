package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Req struct {
	NumOne int
	NumTwo int
}

type Res struct {
	Sum int
}

func main() {
	req := Req{
		NumTwo: 1,
		NumOne: 2,
	}
	var res Res
	client, err := rpc.DialHTTP("tcp", "localhost:8888")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	call := client.Go("Server.Add", req, &res, nil)
	for {
		select {
		case <-call.Done:
			fmt.Println(res)
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("你快点，我等着呢")
		}
	}
	fmt.Println(res)
}
