package main

import (
	"fmt"

	"github.com/shunsukw/design_pattern_go/factory_method/idcard"
)

func main() {
	factory := &idcard.IDCardFactory{}

	card1 := factory.Create(factory, "渡辺")
	card2 := factory.Create(factory, "竣介")

	fmt.Println(card1)
	fmt.Println(card2)
}
