package channels

import (
	"fmt"
)

func CiftSayilar(ciftSayiCn chan int) {
	toplam := 0
	for i := 0; i < 10; i += 2 {
		fmt.Println("Çift Sayı=", i)
		toplam = toplam + i

	}

	ciftSayiCn <- toplam
}
func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i < 10; i += 2 {
		fmt.Println("Tek Sayı=", i)
		toplam = toplam + i
		tekSayiCn <- toplam

	}

}
