package main

type Product struct {
	Name, Category string
	Price          float64
}

var Kayak = Product{
	Name:     "Kayak",
	Category: "Watersports",
	Price:    279,
}

type DiscountedProduct struct {
	*Product `json:",omitempty"`
	Discount float64 `json:"-"`
}

type Named interface{ GetName() string }
type Person struct{ PersonName string }

func (p *Person) GetName() string            { return p.PersonName }
func (p *DiscountedProduct) GetName() string { return p.Name }
