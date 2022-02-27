package conditionals

import "fmt"

func Workshop() {

	var number1 = 5
	var number2 = 10
	var number3 = 18

	enBuyukSayi := number1

	if number2 > enBuyukSayi {
		enBuyukSayi = number2
	}
	if number3 > enBuyukSayi {
		enBuyukSayi = number3
	}
	fmt.Println("En büyük sayı: ", fmt.Sprintf("%v", enBuyukSayi))
}
