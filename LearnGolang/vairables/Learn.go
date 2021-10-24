package vairables

import (
	"fmt"
	"strconv"
)

var fullname string // global
var email string = "Souksavanh.pc@gmail.com"
var c, python bool = true, false

const vat int = 7

func Learn() {
	fullname = "Souksavanh"
	fmt.Println(fullname)
	fmt.Printf("Hello %s Email %s \n", fullname, email)
	fmt.Println(c, python)

	companyname := "Souksavanh"
	isShow := true
	age := 10
	fmt.Println(companyname, isShow, age)
	fmt.Printf("%T \n", age)
	fmt.Printf("%T \n", isShow)

	f := float64(age)
	fmt.Printf("%T \n", f)

	s := strconv.Itoa(vat)
	fmt.Println("vat is" + s)

}
