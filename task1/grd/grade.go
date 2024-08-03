package grd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	sconv "strconv"
	"strings"
)

type Chars struct {
	grades  map[string]float64
	name    string
	courses float64
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func Getstring(prompt string, regex string) (string, error) {
	fmt.Println(prompt)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	matched, _ := regexp.MatchString(regex, input)
	if matched {
		return input, nil
	}
	return "", errors.New("invalid input")

}

func Getint(prompt string) (int, error) {
	fmt.Println(prompt)
	sInput, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	sInput = strings.TrimSpace(sInput)
	input, err := sconv.Atoi(sInput)
	if err == nil {
		return input, nil
	}

	return 0, errors.New("invalid integer")
}

func (x *Chars) Input(n string, c float64) {
	x.name = n
	x.courses = c
	x.grades = make(map[string]float64)
	for c > 0 {

		course, n_err := Getstring("Enter your course name:", `^[^\d]`)
		for n_err != nil{
			fmt.Println(n_err.Error())
			course, n_err = Getstring("Enter your course name:", `^[^\d]`)
		}
		grd, i_err := Getint("Enter your grade:")
		for i_err != nil{
			fmt.Println(i_err.Error())
			grd, i_err = Getint("Enter your grade:")
		}
		grdd := float64(grd)
		
		for grdd < 0 || grdd > 100 {
			fmt.Println("Please enter a valid grade:")
			
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
	fmt.Println(x.name, "Report Card:")
	for crs := range x.grades {
		fmt.Println(crs, x.grades[crs])
	}
	avg := x.calc_grdd()
	fmt.Println("Average", avg)

}
