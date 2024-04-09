package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./text.txt", os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for {
		b := make([]byte, 12)
		n, err := file.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b), n)
	}

}
