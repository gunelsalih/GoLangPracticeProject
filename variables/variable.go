package variables

import "fmt"

func Variable() {

	var text string = "Hello World" // string bir değişken tipi ve metin içimn kullanılır.

	var number int = 5 // int bir değişken tipi tam sayılar için kullanılır.

	fmt.Println(text)
	fmt.Println(number)
	fmt.Println(100 + 100*number/100)

	var number2 float64 = 10.5 // float bir değişken tipi ondalıklı sayılar için kullanılır.
	fmt.Println(number2)
	fmt.Println(100 + 100*number2/100)

	number3 := 56 //basic deger atama
	fmt.Println(number3)
	fmt.Println(number3 * number3)

	number4 := 45.7
	fmt.Println(number4) // format ile yazdırma fmt.PrintF.

	// MANTIKSAL VERİ TİPLERİ

	var status bool = true
	fmt.Println(status)

	text4 := "Umut"
	text5 := "Demet"

	status = text4 == text5 // != eşit değildir "=="çift eşittir,eşittir.
	fmt.Println(status)
	fmt.Println(!status) // ! işareti her şekilde değildir demek.

}
