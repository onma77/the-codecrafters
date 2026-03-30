package main

import (
	"fmt"
)

func main() {
	for {
var onma1 int
fmt.Println("enter onma: ")
fmt.Scanln(&onma1)
var onma2 int
fmt.Println("enter onma2: ")
fmt.Scanln(&onma2)
fmt.Println("1. addition")
fmt.Println("2. subtraction")
fmt.Println("3. multiplication")
fmt.Println("4. division")

var operation int
fmt.Println("choose your option")
fmt.Scanln(&operation)
if operation != 1 && operation != 2 && operation != 3 && operation != 4 {
fmt.Print("invalid option, try again")
continue
}


switch operation {
case 1:
fmt.Println(onma1 + onma2)
case 2:
fmt.Println(onma1 - onma2)
case 3:
fmt.Println(onma1 * onma2)
case 4:
fmt.Println(onma1 / onma2)
}
break
}
}
