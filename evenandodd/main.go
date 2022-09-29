package main

import "fmt"

type myType []int

func main() {
	ints := newMyType()

	for _, myInt := range ints {
		if isEven(myInt) {
			fmt.Printf("%d is even\n", myInt)
		} else {
			fmt.Printf("%d is odd\n", myInt)
		}
	}

}

func newMyType() myType {
	ints := []int{}

	for i := 0; i <= 10; i++ {
		ints = append(ints, i)
	}

	fmt.Println(ints)

	return ints
}

func isEven(myNo int) bool {
	return myNo%2 == 0
}
