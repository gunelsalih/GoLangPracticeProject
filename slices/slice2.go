package slices

import "fmt"

func Slice2() {
	sehirler := []string{"Ankara", "İstanbul", "Çanakkale"}
	fmt.Println(sehirler)
	sehirlerKopya := make([]string,len(sehirler))
	fmt.Println(sehirlerKopya)
	copy(sehirlerKopya,sehirler) 
	fmt.Println(sehirlerKopya)
	sehirler = append(sehirler,"Gümüşhane")
    fmt.Println(sehirler)
	fmt.Println(sehirlerKopya)


	fmt.Println(sehirler[1:4]) 
	fmt.Println(sehirler[1:])
	fmt.Println(sehirler[:2])  

}
