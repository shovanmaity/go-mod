package main

import (
	"fmt"

	bitbucket "bitbucket.org/shovanmaity/go-pvt-module"
	github "github.com/shovanmaity/go-pvt-module"

	gitlab "gitlab.com/shovanmaity/go-pvt-module"
)

func main() {
	fmt.Println("-----------------------------")
	fmt.Println(bitbucket.Message)
	fmt.Println(github.Message)
	fmt.Println(gitlab.Message)
	fmt.Println("-----------------------------")
}
