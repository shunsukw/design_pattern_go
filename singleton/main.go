package main

import (
	"fmt"
	"reflect"

	"github.com/shunsukw/design_pattern_go/singleton/singleton"
)

func main() {
	i1 := singleton.GetInstance()
	i2 := singleton.GetInstance()

	fmt.Println(reflect.DeepEqual(i1, i2))
	fmt.Println(reflect.DeepEqual(&i1, &i2))
}
