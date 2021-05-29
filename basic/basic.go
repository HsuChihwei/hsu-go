package main

import "fmt"

var (  // 可放在函数内或者直接放在包内，可以用var()进行集中定义变量
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue(){
	var a int  // 变量定义，var关键字，有默认值
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue(){
	var a, b int = 3, 4  // 可以一行多定义赋值
	var s string = "abc"
	fmt.Println(a,b,s)
}

func variableTypeDeduction(){
	var a, b, c, s = 3,4,true, "edf"  // 可以省略类型，编译器会自动推断类型
	fmt.Println(a,b,c,s)
}

func variableShorter(){
	a, b, c, s := 3,4,true, "edf"  // 可以省略var关键字，用:=代替定义赋值操作，但是只能在函数内使用
	b = 5
	fmt.Println(a,b,c,s)
}

func main(){
	fmt.Println("Hello Go!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)
}