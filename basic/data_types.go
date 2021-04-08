package main

import (
	"fmt"
	"math"
)

func main() {
	var b bool = true
	fmt.Println(b)

	var uint8Num uint8 = 8
	fmt.Println(uint8Num)

	var uint16Num uint16 = 16
	fmt.Println(uint16Num)

	var uint32Num uint32 = 32
	fmt.Println(uint32Num)

	var uint64Num uint64 = 64
	fmt.Println(uint64Num)

	var int8Num int8 = 8
	fmt.Println(int8Num)

	var int16Num int16 = 16
	fmt.Println(int16Num)

	var int32Num int32 = 32
	fmt.Println(int32Num)

	var int64Num int64 = 64
	fmt.Println(int64Num)

	var floatNum32 float32 = 1.2
	fmt.Println(floatNum32)

	var floatNum64 float64 = 2.4
	fmt.Println(floatNum64)

	var complex64Num complex64 = 64 + 64i
	fmt.Println(complex64Num)

	var complex128Num complex128 = 128 + 128i
	fmt.Println(complex128Num)

	//类似于uint8
	var byteNum byte = 8
	fmt.Println(byteNum)

	//类似于int32
	var runeNum rune = 32
	fmt.Println(runeNum)

	//32位或者64位
	var uintNum uint = 3264
	fmt.Println(uintNum)

	//与uint一样大小
	var intNum int = 3264
	fmt.Println(intNum)

	//无符号整形 用于存放一个指针
	//var uintptrPtr uintptr = & intNum

	//类型转换
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	//类型推导
	var i int
	j := i
	i = 42
	f = 3.142
	g := 0.867 + 0.5i
	fmt.Println(j, i, f, g)

}
