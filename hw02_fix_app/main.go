package main

import (
	"fmt"
	"os"

	"github.com/angbuh/golang-homeworks/hw02_fix_app/printer"
	"github.com/angbuh/golang-homeworks/hw02_fix_app/reader"
	"github.com/angbuh/golang-homeworks/hw02_fix_app/types"
)

func main() {
	var err error
	var staff []types.Employee
	var path string

	fmt.Println("Enter data file path: ")
	fmt.Scanln(&path)

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(printer.StaffToString(staff))
}
