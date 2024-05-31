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

	var uu = user{}
	_ = fmt.Sprint(uu)
}

func allocationWithNew() {
	u := new(user)
	_ = fmt.Sprint(u)
}
