package conditionals

import "fmt"

func Condi() {

	var hesap float64 = 1000
	var cekilmekistenen float64 = 900
	if cekilmekistenen > hesap {
		fmt.Println("Para Yetersiz")

	}
	if cekilmekistenen <= hesap {
		fmt.Println("Paranız Hazırlanıyor")
		hesap = hesap - cekilmekistenen

	}
	fmt.Println("Kalan Bakiyeniz: ", fmt.Sprintf("%v",hesap))
}
