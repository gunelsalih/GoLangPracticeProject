package arrays

import "fmt"

func Demo() {
	var citys [6]string
	citys[0] = "Ankara"
	citys[1] = "İstanbul"
	citys[2] = "İzmir"
	citys[3] = "Adana"
	citys[4] = "Çanakkale"
	citys[5] = "Bursa"
	fmt.Println(citys)
	
	for i := 0; i < 6; i++ {
		fmt.Println(citys[i])
		


	}


}
