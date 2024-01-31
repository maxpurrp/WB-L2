package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern


Паттерн строитель - это порождающий шаблон проектирования, который испрользуется для пошагового
	создания сложного объекта с использованем одного и того же коннструктора. шаблон отделяет
	процесс конструирования от его представления, создавая различные представления объекта.
Плюсы - 1.  Изолирование кода, который реализует представление и конструирование
		2.  Позволяет изменять внутренее представление продукта
		3. более тонкий контроль над процессом конструирования

Минусы - 1. Процесс конструирования должен обеспечить различные представленияобъекта
		 2. алгоритм создания сложного объекта не должен зависеть от того
		 	из каких частей состоит объект и как они стыкуются между собой
*/

// Product
type ShopProduct struct {
	Category  string
	Brand     string
	Model     string
	Price     float64
	Available bool
}

// Builder
type ShopProductBuilder interface {
	SetCategory(category string)
	SetBrand(brand string)
	SetModel(model string)
	SetPrice(price float64)
	GetShopProduct() ShopProduct
}

// ConcreteBuilder
type SmartphoneBuilder struct {
	product ShopProduct
}

func (b *SmartphoneBuilder) SetCategory(category string) {
	b.product.Category = category
}

func (b *SmartphoneBuilder) SetBrand(brand string) {
	b.product.Brand = brand
}

func (b *SmartphoneBuilder) SetModel(model string) {
	b.product.Model = model
}

func (b *SmartphoneBuilder) SetPrice(price float64) {
	b.product.Price = price
}

func (b *SmartphoneBuilder) GetShopProduct() ShopProduct {
	return b.product
}

// Director
type ShopDirector struct {
	builder ShopProductBuilder
}

func NewShopDirector(builder ShopProductBuilder) *ShopDirector {
	return &ShopDirector{builder: builder}
}

func (d *ShopDirector) ConstructShopProduct() {
	d.builder.SetCategory("Electronics")
	d.builder.SetBrand("Apple")
	d.builder.SetModel("Iphone 144s")
	d.builder.SetPrice(9999.2)
}

// func main() {
// 	smartphoneBuilder := &SmartphoneBuilder{}
// 	shopDirector := NewShopDirector(smartphoneBuilder)

// 	shopDirector.ConstructShopProduct()
// 	smartphoneInShop := smartphoneBuilder.GetShopProduct()

// 	fmt.Println("Specs smartphone in the Shop :")
// 	fmt.Printf("Category: %s\nBrand: %s\nModel: %s\nPrice: $%.2f\n",
// 		smartphoneInShop.Category, smartphoneInShop.Brand, smartphoneInShop.Model,
// 		smartphoneInShop.Price)
// }
