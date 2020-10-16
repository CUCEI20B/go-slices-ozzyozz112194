package main

import "fmt"

func main() {

	var number1 int
	var number2 int
	var number3 int
	var total int = 0

	fmt.Scanln(&number1)
	number3 = number1 - 1

	//slice
	s := make([]int, 0, number3)
	i := 0

	for number3 >= i {

		fmt.Scanln(&number2)
		total += number2
		s = append(s, number2)

		i++

	}

	for _, v := range s {
		fmt.Println(v)
	}
	print(total)

}
