package arrays

import "fmt"

func Demo2() {

	sayilar := [5]int{20, 30, 50, 10, 2}
	fmt.Println(sayilar)

	maxnumber := sayilar[0]
	for i := 0; i < len(sayilar); i++ {
		if  sayilar[i] > maxnumber {
              maxnumber = sayilar[i]
		}
	}

	fmt.Println("En Büyük Sayı" , maxnumber)
}
