package main

import (
	"bufio"
	"fmt"
	ctrls "library/controllers"
	model "library/models"
	services "library/services"
	"os"
	"strings"

	"github.com/fatih/color"
)

var lib services.Library = services.Library{
	Books:   make(map[int]*model.Book),
	Members: make(map[int]*model.Member),
}
var reader bufio.Reader = *bufio.NewReader(os.Stdin)

func menu(id int) {
	exit := 0
	for exit == 0 {
		color.Yellow(strings.Repeat("-", 20))
		fmt.Printf("ID: %d, USER:%s\n", id, lib.Members[id].Name)
		ctrls.LISTAVAILABLE(&lib)

		fmt.Println("MENU:")
		color.Yellow(strings.Repeat("-", 20))
		fmt.Println("1-> ADDBOOK")
		fmt.Println("2-> REMOVEBOOK")
		fmt.Println("3-> BORROWBOOK")
		fmt.Println("4-> RETURNBOOK")
		fmt.Println("5-> LISTMYBORROWEDBOOKS")
		fmt.Println("6-> LOGOUT")

		ch, err := reader.ReadString('\n')
		ch = strings.TrimSpace(ch)
		for err != nil {
			fmt.Println(err.Error())
			ch, err = reader.ReadString('\n')
		}

		switch ch {
		case "1":
			ctrls.ADDBOOK(&lib)
			break

		case "2":
			ctrls.REMOVE(&lib)
			break

		case "3":
			ctrls.BORROW(&lib, id)
			break

		case "4":
			ctrls.RETURN(&lib, id)
			break

		case "5":
			ctrls.LISTBORROWED(&lib, id)
			break

		case "6":
			exit = 1
			break

		default:
			fmt.Println("Invalid INPUT")
			break

		}

	}

}

func main() {
	exit := 0
	for exit == 0 {
		color.Red(strings.Repeat("-", 20))

		fmt.Println("1-> LOGIN")
		fmt.Println("2-> SIGNUP")
		fmt.Println("0-> EXIT")

		ch, err := reader.ReadString('\n')
		ch = strings.TrimSpace(ch)
		for err != nil {
			fmt.Println(err.Error())
			ch, err = reader.ReadString('\n')
		}

		switch ch {
		case "1":
			ID := ctrls.LOGIN(&lib)
			menu(ID)
			break

		case "2":
			ID := ctrls.SIGNUP(&lib)
			menu(ID)
			break

		case "0":
			exit = 1
			break

		default:
			fmt.Println("Invalid INPUT")
			break

		}

	}

}
