package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Inventory struct {
	Total   int      `json: "total`
	Devices []string `json: "devices"`
}

func main() {
	data, _ := ioutil.ReadFile("file.json")

	inv := Inventory{}
	_ = json.Unmarshal([]byte(data), &inv)

	fmt.Printf("Total: %d\nDevices: %q\n", inv.Total, inv.Devices)

}
