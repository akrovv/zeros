package tests

import (
	"fmt"
	"testing"
)

func TestWithoutZeroValueStruct(t *testing.T) {
	type User struct {
		Name string
	}

	u := User{Name: "Artyom"}
	_ = fmt.Sprint(u)
}
