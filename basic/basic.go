package main

import "fmt"

func variableInitialValue(){
	var a int = 3
	var s string = "abc"
	fmt.Println(a,s)
}

func main(){
	fmt.Println("Hello Go!")
	variableInitialValue()
}