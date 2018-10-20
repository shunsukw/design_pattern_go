package framework

type Product interface {
	Use() string
}

type Factory struct {
}

func (factory *Factory) Create(creater FactoryInterface, owner string) Product {
	user := creater.CreateProduct(owner)
	creater.RegisterProduct(user)
	return user
}

type FactoryInterface interface {
	CreateProduct(owner string) Product
	RegisterProduct(product Product)
}
