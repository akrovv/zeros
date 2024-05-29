package valid

import (
	"fmt"
)

type user struct {
	Name string
}

func withoutZeroValue() {
	u := user{Name: "Artyom"}
	_ = fmt.Sprint(u)
}
