package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Manager struct {
	FullName       string
	Position       string
	Age            int32
	YearsInCompany int32
}

func EncodeManager(manager *Manager) (io.Reader, error) {
	encoder := json.NewEncoder(os.Stdout)
	c := struct {
		FullName       string `json:"full_name"`
		Position       string `json:"position"`
		Age            int32  `json:"age"`
		YearsInCompany int32  `json:"years_in_company"`
	}{
		FullName:       manager.FullName,
		Position:       manager.Position,
		Age:            manager.Age,
		YearsInCompany: manager.YearsInCompany,
	}
	encoder.Encode(c)
	r := strings.NewReader("My Reader")
	return r, 
}

func main() {
	fmt.Println("bismillah")

}
