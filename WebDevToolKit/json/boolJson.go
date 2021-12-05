package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var data bool

	rcvd := `true`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}
