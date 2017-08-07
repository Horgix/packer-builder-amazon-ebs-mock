package amazonebsmock

import (
	"fmt"
	"testing"
)

func TestPrepare(t *testing.T) {
	builder := new(Builder)
	str, _ := builder.Prepare()
	fmt.Println(str)
}
