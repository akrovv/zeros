package invalid

import (
	"fmt"
)

type user struct {
	Name string
}

func zeroValue() {
	u := user{}
	_ = fmt.Sprint(u)
}

func allocationWithNew() {
	u := new(user)
	_ = fmt.Sprint(u)
}
