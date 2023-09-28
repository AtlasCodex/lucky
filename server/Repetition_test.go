package server

import (
	"fmt"
	"testing"
)

func TestRepetition(t *testing.T) {
	code := "dlt"
	number := []interface{}{"04", "09", "11", "13", "17", "07", "08"}
	result := Calculate(code, number)
	fmt.Println(result)

}
