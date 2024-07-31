package main

import (
	"fmt"
	grade "task1/grd"
)

func test(){
	name, n_err := grade.Getstring("Enter your name:", `^[^\d]`)
	for n_err != nil{
		fmt.Println(n_err.Error())
		name, n_err = grade.Getstring("Enter your name:", `^[^\d]`)
	}
	num, i_err := grade.Getint("Enter the number of courses you take:")
	for i_err != nil{
		fmt.Println(i_err.Error())
		num, i_err = grade.Getint("Enter the number of courses you take:")
	}

	var chr grade.Chars
	chr.Input(name, float64(num))
	chr.Display()

}

func main(){
	test()
}

