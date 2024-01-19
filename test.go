package main

import (
	"fmt"
	"module3/pack1"
	"module3/pack2"
	"module3/pack3"
	"module3/pack4"
)

func main() {
	fmt.Println(pack1.Add(10, 20))
	fmt.Println(pack2.Sub(10, 20))
	fmt.Println(pack3.Multiply(10, 20))
	quo, rem := pack4.Division(10, 20)
	fmt.Println(quo)
	fmt.Println(rem)
	fmt.Println("The assignment is completed")
}
