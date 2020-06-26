package CustomModule

import (
	"encoding/json"
	"fmt"
)

func MakeJson() {
	var name string
	var addr string
	fmt.Print("Enter your name : ")
	fmt.Scan(&name)
	fmt.Print("Enter your address : ")
	fmt.Scan(&addr)
	nameAndAddr := make(map[string]string)
	nameAndAddr["name"] = name
	nameAndAddr["address"] = addr
	jsonObj, _ := json.Marshal(nameAndAddr)
	fmt.Println(string(jsonObj))
}
