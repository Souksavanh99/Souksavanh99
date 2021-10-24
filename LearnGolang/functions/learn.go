package functions

import (
	"fmt"
	"strings"
)

// var fullname string
func add(x, y int) int {

	return x * y
}

func converter(titel, email string) (string, string) {
	s1 := strings.ToUpper(titel)
	return s1, email
}

func sum(numbers ...int) {
	// fmt.Println(numbers) //Array

	total := 0
	for _, v := range numbers {
		total += v

	}
	fmt.Println(total)

}

func Learn() {

	fmt.Println(add(100, 5))
	fmt.Println(converter("Hello", "Hello@gmail.com"))
	sum(10, 20, 30, 40, 50)
}
