package loops

import (
	"fmt"
)

//kullanıcıdan alınan sayının asal olup olmadığını bulan program

func Workshop2() {
	number := 0
	fmt.Println("Bir sayı giriniz")
	fmt.Scanln(&number)

	asalMi := true

	for i := 2; i < number; i++ {

		if number%i == 0 {
			asalMi = false
		}

	}
	if asalMi {
		fmt.Println("Asal")

	} else {
		fmt.Println("Asal Değil")

	}
}
