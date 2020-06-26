package CustomModule

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	name    string
	surname string
}

func ReadFile() {
	var fileName string
	var person Person
	personSlice := make([]Person, 0, 1)
	fmt.Println("Enter the file name:")
	fmt.Scan(&fileName)
	fileDesc, err := os.Open(fileName)
	scanner := bufio.NewScanner(fileDesc)
	if err == nil {
		fmt.Println("File opened!")
		var line string
		for scanner.Scan() {
			line = scanner.Text()
			nameAndSurname := strings.Split(line, " ")
			person.name = nameAndSurname[0]
			person.surname = nameAndSurname[1]
			personSlice = append(personSlice, person)
		}
		fmt.Println("File Closed!")
		fileDesc.Close()
	} else {
		fmt.Println("File not found!")
	}

	for _, e := range personSlice {
		fmt.Println(e.name + " " + e.surname)
	}
}
