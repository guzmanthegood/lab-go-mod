package main

import (
	"fmt"

	"github.com/guzmanweb/hellomod"
	"github.com/guzmanweb/lab-go-mod/mypkg"
)

func main() {
	fmt.Println("Hello world! Main function.")
	mypkg.MyHelloWorld()
	hellomod.SayHello()
}
