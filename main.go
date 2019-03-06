package main

import (
    "fmt"
)

//関数を返す関数

func later() func(string) string {
	//１つ前に与えられた文字列を保存するための変数
	var store string
	//引数に文字列をとり文字列を返す関数を返す
	return func(next string) string {
		// var s string
 	// 	s = store
		s := store
		store = next
		return s
	}
}

func returnFunc() func() string{
	return func() string{
		return "Hello Golang"
	}
}

func intergers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	fmt.Println("Hello World!");

	f := later()
	ff := returnFunc()
	fff := intergers()
	ffff := intergers()

	ints := intergers()

	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(fff())
	fmt.Println(fff())
	fmt.Println(fff())
	fmt.Println(fff())
	fmt.Println(ffff())
    
    ff()
	fmt.Println(f("Golang")) //""
	fmt.Println(f("is")) // "Golang"
	fmt.Println(f("awesome!")) // "is"


}
