package slices

import "fmt"

func Slice() {
	isimler := make([]string, 3)

	fmt.Println(isimler)
	isimler[0] = "Sinan"
	isimler[1] = "Ahmet"
	isimler[2] = "Rasim"
	isimler = append(isimler, "Salih")

	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
