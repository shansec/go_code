package main

import (
	"fmt"
	"jwt_study/generate"
)

func main() {
	// 生成 token
	token, err := generate.GenTokenByCustom("may")
	if err != nil {
		fmt.Printf("generate jwt failed err: %v\n", err)
		return
	}
	fmt.Printf("generate jwt success, token: %s\n", token)

	// 解析 token
	claims, err := generate.ParseRegisteredClaims(token)
	if err != nil {
		fmt.Printf("parsed token failed err: %v\n", err)
		return
	}
	fmt.Printf("parsed success, result: %v", claims)
}
