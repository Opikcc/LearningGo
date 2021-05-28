package main

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/opikcc/calc"
)

func main() {
	x := 5
	y := 3

	fmt.Println(calc.Add(x, y))
	fmt.Println(calc.Subtract(x, y))

	var array [5]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	fmt.Println(array)

	arrayLiteral := [5]int{10, 11, 12, 13, 14}

	fmt.Println(arrayLiteral)

	arrayMultiLine := [5]int{
		100,
		200,
		300,
		400,
		500,
	}

	fmt.Println(arrayMultiLine)

	arrayDotDotDot := [...]int{
		1000,
		2000,
		3000,
	}

	fmt.Println(arrayDotDotDot)

	arrayIndexed := [5]int{2: 125, 4: 350}

	fmt.Println(arrayIndexed)

	var sliceMake []int
	sliceMake = make([]int, 5)

	sliceMake = append(sliceMake, 1)
	sliceMake = append(sliceMake, 2)
	sliceMake = append(sliceMake, 3)
	sliceMake = append(sliceMake, 4)
	sliceMake = append(sliceMake, 5)

	fmt.Println(sliceMake)

	sliceAppend := []int{1, 22, 333}

	sliceAppended := append(sliceAppend, 4444, 55555)

	fmt.Println(sliceAppended)

	sliceX := make([]int, 2, 5)
	sliceX[0] = 10
	sliceX[1] = 20
	fmt.Println("Length of sliceX : ", len(sliceX))
	fmt.Println("Capacity of sliceX : ", cap(sliceX))
	sliceX = append(sliceX, 30, 40, 50)
	fmt.Println("Length of sliceX : ", len(sliceX))
	fmt.Println("Capacity of sliceX : ", cap(sliceX))

	for k, v := range sliceX {
		fmt.Printf("Index: %d Value: %d\n", k, v)
	}

	dict := make(map[string]string)
	dict["go"] = "Golang"
	dict["cs"] = "CSharp"
	dict["rb"] = "Ruby"
	dict["py"] = "Python"
	dict["js"] = "JavaScript"
	for k, v := range dict {
		fmt.Printf("Key: %s Value: %s\n", k, v)
	}

	fmt.Println("Starting to panic")
	doPanic()
	fmt.Println("Program regains control after panic recover")

	f, err := os.Open("readme.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(f)
}

func doPanic() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recover with: ", e)
		}
	}()
	panic("Just panicking for the sake of demo")
	fmt.Println("This will never be called")
}
