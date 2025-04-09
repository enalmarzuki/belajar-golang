package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func Test(t *testing.T) {
	logJson("Jhon")
	logJson(1)
	logJson(true)
	logJson([]string{"Jhon", "Doe", "Vale"})
}
