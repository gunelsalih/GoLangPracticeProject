package loops

import (
	"fmt"
)

func Workshop3() {

	number := 220
	number2 := 284
	sumNum := 0
	sumNum2 := 0

	for i := 1; i < number; i++ {
		if number%i == 0 {
			sumNum = i + sumNum
		}
	}
	for i := 1; i < number2; i++ {

		if number2%i == 0 {
			sumNum2 = i + sumNum2
		}
	}
	if number == sumNum2 && number2 == sumNum {
		fmt.Println("Arkadas")

	}

}
