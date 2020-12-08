package main

import (
	"fmt"
	"reflect"
)


func main(){
	var str string = "Hello world"


	fmt.Println("=========Find out the type of a variable=========")
	fmt.Println()
	fmt.Println("1) By Percent T Printf")
	fmt.Printf("str => %T", str)
	fmt.Println()
	fmt.Println()
	fmt.Println("2) Using reflect.TypeOf(variable)")
	fmt.Println("str => ", reflect.TypeOf(str))
	fmt.Println()
	fmt.Println("3) Using reflect.ValueOf(variable).Kind()")
	fmt.Println("str => ", reflect.ValueOf(str).Kind())
}