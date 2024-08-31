package main

import "fmt"

// Add 整数加法
func Add(testA int, testB int) int {
	return testA + testB
}

type animal struct {
	data int
	cli  string
}

func (*animal) bark() {
	fmt.Print("wangwang")
}

//
func main() {
	fmt.Println("hello world")
	fmt.Println(Add(12, 24))
	duck := new(animal)
	duck.bark()
}
