package main

import (
	"fmt"
	"math"
)

func main() {
	formatPrintf()

	var number float64 = 6
	dozen := doDouble(number)

	fmt.Println(dozen)
	fmt.Println(floatParts(1.51))

	amount := 6
	fmt.Println(&amount)
	myPointerTest()

	var myDouble = 6
	doubleByPointer(&myDouble)
	fmt.Printf("通过指针修改值后： %v \n", myDouble)

	testStruct()
	testTypeDefine()
	customType := testReturnCustomType("来100个")
	fmt.Printf("%#v", customType)
}

func doubleByPointer(numberPointer *int) {
	*numberPointer = *numberPointer * 2

}

func formatPrintf() {
	fmt.Println("打印的格式化动词")
	fmt.Printf("默认打印 %v %v %v ", "", "\t", "\n")
	fmt.Println("=====")
	fmt.Printf("打印不可见字符：%#v %#v %#v ", "", "\t", "\n")
	fmt.Println()
}

func doDouble(number float64) float64 {
	double := number * 2
	return double
}

func floatParts(number float64) (intPart int, fractionalPart float64) {
	myInt := math.Floor(number)
	return int(myInt), number - myInt
}

func myPointerTest() {
	number := 6
	numberPointer := &number

	fmt.Printf("the pointer is %v \n", numberPointer)
	fmt.Printf("the pointer stored value is %v \n", *numberPointer) // 指针前面加 * 表示

	*numberPointer = 8 // *用于通过指针给变量赋值
	fmt.Println(number)
}

func testStruct() {
	var myStruct struct {
		number int
		name   string
		myBool bool
	}
	myStruct.number = 1
	myStruct.name = "2233"
	myStruct.myBool = true
	fmt.Printf("%#v \n", myStruct)
}

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func testTypeDefine() {
	var myPart part
	myPart.description = "wheel"
	myPart.count = 4
	var byd car
	byd.name = "tang"
	byd.topSpeed = 200
	fmt.Printf("myPart is %#v \n", myPart)
	fmt.Printf("byd car is %#v \n", byd)
}

// 入参或返回值可以使用自定义 type
func testReturnCustomType(desc string) part {
	var p part
	p.description = desc
	p.count = 100
	return p
}
