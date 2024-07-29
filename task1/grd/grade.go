package grd

import (
	sys "fmt"
)

type Chars struct {
	grades  map[string]float64
	name    string
	courses float64
}

func (x *Chars) Input(n string, c float64) {
	x.name = n
	x.courses = c
	x.grades = make(map[string]float64)
	for c > 0 {
		var course string
		var grdd float64
		sys.Println("Enter your course and grade:")
		sys.Scan(&course, &grdd)
		for grdd < 0 || grdd > 100 {
			sys.Println("Please enter a valid grade:")
			sys.Scan(&grdd)
		}
		x.grades[course] = grdd
		c--
	}
}

func (x *Chars) calc_grdd() float64 {
	var avg float64
	for crs := range x.grades {
		avg += x.grades[crs]
	}

	return avg / x.courses
}

func (x *Chars) Display() {
	sys.Println(x.name, "Report Card:")
	for crs := range x.grades {
		sys.Println(crs, x.grades[crs])
	}
	avg := x.calc_grdd()
	sys.Println("Average", avg)

}
