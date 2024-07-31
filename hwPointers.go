package main

import (
	"fmt"
	"math"
)

const pi = math.Pi

func main() {
	//printCircleArea(10)
	//printRectangleArea(10, 10)
	printRectanglePerimeter(10, 10)
}

func printCircleArea(radius int) {
	//S = πr^2
	area := float32(radius) * float32(radius) * pi
	fmt.Println("Area of a circle:", area)
}

func printRectangleArea(lenght int, width int) {
	//S = lenght × width
	area := lenght * width
	fmt.Println("Area of a rectangle:", area)
}

func printRectanglePerimeter(lenght int, width int) {
	//P = 2(lenght + width)
	perimeter := 2 * (lenght + width)
	fmt.Println("Perimeter of a rectangle:", perimeter)
}
