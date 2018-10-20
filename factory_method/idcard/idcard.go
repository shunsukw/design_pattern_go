package idcard

import "github.com/shunsukw/design_pattern_go/factory_method/framework"

type IDCard struct {
	Owner string
}

func (idcard *IDCard) Use() string {
	return idcard.Owner
}

// Factoryの具象はFactoryInterfaceを実装し、さらにFactory構造体を委譲している
// これはJavaでは抽象クラスにメソッドとインターフェイスを両方定義できるのに対して
// Goではできないから
type IDCardFactory struct {
	*framework.Factory
	owners []*string
}

func (idCardFactory *IDCardFactory) CreateProduct(owner string) framework.Product {
	return &IDCard{
		Owner: owner,
	}
}

func (idCardFactory *IDCardFactory) RegisterProduct(product framework.Product) {
	owner := product.Use()
	idCardFactory.owners = append(idCardFactory.owners, &owner)
}
