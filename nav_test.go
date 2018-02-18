package position

import (
	"fmt"
)

func ExampleNavigator_Right() {
	x, y, inside := 0, 0, true
	nav := NewXYNavigator(x, y, 0, 0, 2, 2)
	for ; inside; x, y, inside = nav.Right() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 0,0
	// 1,0
	// 2,0
	// 0,1
	// 1,1
	// 2,1
	// 0,2
	// 1,2
	// 2,2
}

func ExampleNavigator_Right_full() {
	x, y, inside := 0, 0, true
	nav := NewXYNavigator(x, y, x, y, 2, 2)
	for ; inside; x, y, inside = nav.Right() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 0,0
	// 1,0
	// 2,0
	// 0,1
	// 1,1
	// 2,1
	// 0,2
	// 1,2
	// 2,2
}

func ExampleNavigator_Up() {
	x, y, inside := 1, 1, true
	nav := NewXYNavigator(1, 1, 0, 0, 1, 1)
	for ; inside; x, y, inside = nav.Up() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 1,1
	// 1,0
	// 0,1
	// 0,0
}

func ExampleNavigator_Left() {
	nav := NewXYNavigator(2, 2, 1, 1, 2, 2)
	for x, y, more := nav.Left(); more; x, y, more = nav.Left() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 1,2
	// 2,1
	// 1,1
}

func ExampleNavigator_Down() {
	x, y, inside := 1, 1, true
	nav := NewXYNavigator(x, y, 1, 1, 2, 2)
	for ; inside; x, y, inside = nav.Down() {
		fmt.Printf("%v,%v\n", x, y)
	}
	// output:
	// 1,1
	// 1,2
	// 2,1
	// 2,2
}
