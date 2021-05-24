package main

import (
	"github.com/gvencadze/swe-basic/increment"
)

func main() {
	inc := increment.New(0)
	inc.SetMaximumValue(10)
	inc.Increment()
}
