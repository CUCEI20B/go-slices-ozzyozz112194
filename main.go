package main

import "fmt"

func main() {

	var number1 int
	var number2 int

	var total int = 0

	fmt.Scanln(&number1)

	//slice
	s := make([]int, 0, number1)
	i := 0

	for number1 >= i {

		println("Entrada:")
		fmt.Scanln(&number2)
		total += number2
		s = append(s, number2)

		i++

	}

	for _, v := range s {
		fmt.Println(v)
	}
	println("Salida:")
	println(total)

}
