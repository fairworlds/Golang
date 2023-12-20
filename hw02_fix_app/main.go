package main

import (
	"fmt"

	"github.com/fairworlds/Golang/hw02_fix_app/printer"
	"github.com/fairworlds/Golang/hw02_fix_app/reader"
)

func main() {
	path := "/home/fair/Golang/hw02_fix_app/data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	///var err error
	///var staff []types.Employee

	if len(path) == 0 {
		path = "/home/fair/Golang/hw02_fix_app/data.json"
	}

	staff, err := reader.ReadJSON(path)
	if err != nil {
		fmt.Print(err)
	}
	///fmt.Print(err)

	printer.PrintStaff(staff)
}
