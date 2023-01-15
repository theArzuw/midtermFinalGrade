package main

import "fmt"

func main() {

	//Öğrencinin vize ve final notunu al
	var vizeNotu float32
	var finalNotu float32

	fmt.Println("Lütfen vize notunu giriniz:")
	fmt.Scanln(&vizeNotu)
	fmt.Println("Lütfen final notunu giriniz")
	fmt.Scanln(&finalNotu)

	//Vize ve final notunu topla ortalamasını al
	notTopla := vizeNotu + finalNotu
	ortalama := notTopla / 2

	//Ortalama notu ekrana yazdır
	fmt.Println("Ortalama:", ortalama)

	//Başka bir hesaplama yapmak isteyip istemediğini sor
	fmt.Println("Başka bir hesap yapmak ister misiniz? (evet/hayır)")
	var cevap string
	fmt.Scanln(&cevap)

	//Kullanıcı başka bir hesaplama yapmak isterse başa dön

	if cevap == "evet" {

		//Kullanıcı başka bir hesaplama yapmak istemiyorsa bitir

	} else if cevap == "hayır" {
		fmt.Println("İşlem tamamlandı. Teşekkür ederiz.")

	}
}
