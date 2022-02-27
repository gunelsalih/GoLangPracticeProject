package conditionals

import "fmt"

func CondiIfElse() {

	var hesap float64 = 1000
	var cekilmekistenen float64 = 1000
	if cekilmekistenen > hesap {
		fmt.Println("Para Yetersiz")

	} else if hesap==cekilmekistenen { 

		fmt.Println("Paranız Hazırlanıyor")
		fmt.Println("Hesabınızda Para Kalmayacak")
		hesap = hesap - cekilmekistenen


	}else {

		fmt.Println("Paranız Hazırlanıyor")
		hesap = hesap - cekilmekistenen

		
	}
	fmt.Println("Kalan Bakiyeniz: ", fmt.Sprintf("%v",hesap))
}
