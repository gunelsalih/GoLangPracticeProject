package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Bilgisayar", 500, "XYZ",20})

	fmt.Println(product{"Bilgisayar", 500, "ABC",20})

	fmt.Println(product{name: "Bilgisayar", unitPrice: 500, brand:"XYZ"})
}

type product struct {
	name         string
	unitPrice    float64
	brand        string
	discountRate int
}
