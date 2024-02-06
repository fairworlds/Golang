package fixapp

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Employee struct {
	UserID       int    `json:"userId"`
	Age          int    `json:"age"`
	Name         string `json:"name"`
	DepartmentID int    `json:"departmentId"`
}

func (e Employee) String() string {
	return fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d; ", e.UserID, e.Age, e.Name, e.DepartmentID)
}

func ReadJSON(filePath string) ([]Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	var data []Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	res := data

	return res, nil
}

func PrintStaff(staff []Employee) {
	for i := 0; i < len(staff); i++ {
		fmt.Println(staff[i])
	}
}
