package main

import (
	"fmt"

	"github.com/abhinav-0401/weep-8/weep"
)

func main() {
	fmt.Println("Hi, I'm Weep-8!")
	var system = weep.New()
	// system.Pc = 10
	// fmt.Println(system.Pc)
	// system.LoadRom("spitball/test.txt")
	// var time = time.Now().Unix()
	fmt.Println(system.GenRand())
	fmt.Println(system.GenRand())
	fmt.Println(system.GenRand())
}