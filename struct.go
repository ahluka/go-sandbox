package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type Employee struct {
	ID        int `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"addr"`
	Position  string `json:"pos"`
	Salary    int `json:"salary"`
	ManagerID int `json:"mgrid"`
}

func (e Employee) String() string {
	return fmt.Sprintf("%d - %s - %s - £%d\n", e.ID, e.Name, e.Position, e.Salary)
}

func main() {
	lee := Employee{0, "Lee Thomas", "13 Somerton Road", "Admin", 18525, 0}
	fmt.Println(lee)

	j, err := json.MarshalIndent(lee, "", "\t")
	if err != nil {
		log.Fatalf("json.Marshal: %v", err)
	}

	fmt.Println(string(j))
}
