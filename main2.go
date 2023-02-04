package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
)

func main2() {
	simplePrint()
	typeOfDemo()
	varDemo()
	zeroInitVar()
	shortVarDefine()
	passFailTest()
}

func passFailTest() {
	fmt.Println("####  pass fail test func called, input your score: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}

func shortVarDefine() {
	fmt.Println("####  短变量声明")

	quantity := 4
	length, width := 11, 4.0
	fmt.Println(quantity)
	fmt.Println(length, width)

}

func zeroInitVar() {
	// 变量初始值是零值
	fmt.Println("#### 变量初始值是零值")
	var myString string
	var myBool bool
	fmt.Println(myString, myBool)

}

func varDemo() {
	fmt.Println("#### var demo called")

	var quantity, length int
	quantity = 2
	length = 11
	fmt.Println(quantity, length)
}

func typeOfDemo() {
	fmt.Println("#### type of demo called")
	fmt.Println(reflect.TypeOf(true))
}

func simplePrint() {
	floor := math.Floor(2.75)
	fmt.Println(floor)
	fmt.Println("hs!!!")
	fmt.Println('A') // Unicode 数
	fmt.Println('嗨')
}
