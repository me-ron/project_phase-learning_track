package main 

import (
	sys "fmt"
	grade "task1/grd"
)

func test(){
	var name string; var num float64
	sys.Println("Enter your name:")
	sys.Scan(&name)
	sys.Println("Enter the number of courses you take:")
	sys.Scan(&num)

	var chr grade.Chars
	chr.Input(name, num)
	chr.Display()

}

func main(){
	test()
}

