package pointers

import "fmt"

func Pointer2(sayilar []int) {
	sayilar[0] = 100
	fmt.Println("Pointer2deki sayi:", sayilar[0])

}
