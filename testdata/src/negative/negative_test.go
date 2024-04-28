package negative

import (
	"fmt"
	"testing"
)

func TestWithZeroValueStruct(t *testing.T) {
	type User struct {
		Name string
	}

	u := User{}
	_ = fmt.Sprint(u)
}
