package pointers

import "fmt"

func Pointer(sayi *int) {
	*sayi = *sayi + 1
	fmt.Println("Pointerdeki sayı:", *sayi)
}
