package loops

import "fmt"
//3 tahminde bilirse süper 4-6 arası bilirse güzel 6 dan fazla bilirse fena değil
func Workshop() {

	mainnumber := 50
	guessNumber := 0
	guessCount := 0


	fmt.Println("1-100 arasında bir sayı giriniz")
	
	for mainnumber != guessNumber {
		fmt.Scanln(&guessNumber)
		
		if mainnumber > guessNumber  {
			fmt.Println("Daha BÜYÜK sayı gir")			
		}

		if mainnumber < guessNumber {
			fmt.Println("Daha KÜÇÜK sayı gir")			
		}

		guessCount ++
	}
	fmt.Println("İyi iyi aferin.")		

	if guessCount <=3 && guessCount >=1 {
		fmt.Println("Süper")
	}
	if guessCount >=4 && guessCount <=6 {
		fmt.Println("Güzel")
	}
	if guessCount >6  {
		fmt.Println("Fena Değil")
		
	}

}