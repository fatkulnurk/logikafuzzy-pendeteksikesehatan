/******************************************************************
*
			Nama	: Fatkul Nur Koirudin
			Nrp		: 2110161019
			Matkul	: Kecerdasan Komputasional (AI)
*
*******************************************************************/
package main

import (
	"fmt"
)

const MAX  = 1000

type Data struct {
	key int
	label string
	min float32
	max float32
}

type Aturan struct {
	label1	string
	label2	string
	hasil	string
}

var Tinggi float32
var Berat float32

func main(){
	var tempNum 	float32
	var tempNum2 	float32
	var counter 	int
	//var lastMax		float32
	//var lastMin		float32
	//var tempBerat	float32
	//var tempBerat2	float32
	var tempKeyBerat    string
	var tempKeyTinggi   string
	var aturan = []Aturan{
		{label1:"sangat pendek", label2:"sangat kurus", hasil:"SS"},
		{label1:"sangat pendek", label2:"kurus", hasil:"S"},
		{label1:"sangat pendek", label2:"biasa", hasil:"AS"},
		{label1:"sangat pendek", label2:"berat", hasil:"TS"},
		{label1:"sangat pendek", label2:"sangat berat", hasil:"TS"},
		{label1:"pendek", label2:"sangat kurus", hasil:"S"},
		{label1:"pendek", label2:"kurus", hasil:"SS"},
		{label1:"pendek", label2:"biasa", hasil:"S"},
		{label1:"pendek", label2:"berat", hasil:"AS"},
		{label1:"pendek", label2:"sangat berat", hasil:"TS"},
		{label1:"sedang", label2:"sangat kurus", hasil:"AS"},
		{label1:"sedang", label2:"kurus", hasil:"SS"},
		{label1:"sedang", label2:"biasa", hasil:"SS"},
		{label1:"sedang", label2:"berat", hasil:"AS"},
		{label1:"sedang", label2:"sangat berat", hasil:"TS"},
		{label1:"tinggi", label2:"sangat kurus", hasil:"TS"},
		{label1:"tinggi", label2:"kurus", hasil:"S"},
		{label1:"tinggi", label2:"biasa", hasil:"SS"},
		{label1:"tinggi", label2:"berat", hasil:"S"},
		{label1:"tinggi", label2:"sangat berat", hasil:"TS"},
		{label1:"sangat tinggi", label2:"sangat kurus", hasil:"TS"},
		{label1:"sangat tinggi", label2:"kurus", hasil:"AS"},
		{label1:"sangat tinggi", label2:"biasa", hasil:"SS"},
		{label1:"sangat tinggi", label2:"berat", hasil:"S"},
		{label1:"sangat tinggi", label2:"sangat berat", hasil:"AS"},
	}

	var aturanTinggi = []Data{
		{key:1,label:"sangat pendek",min:0,max:120},
		{key:2,label:"pendek",min:115,max:145},
		{key:3,label:"sedang",min:140,max:165},
		{key:4,label:"tinggi",min:160,max:185},
		{key:5,label:"sangat tinggi",min:180,max:MAX},
	}
	var aturanBerat = []Data{
		{key:1,label:"sangat kurus",min:0,max:45},
		{key:2,label:"kurus",min:40,max:55},
		{key:3,label:"biasa",min:50,max:65},
		{key:4,label:"berat",min:60,max:85},
		{key:5,label:"sangat berat",min:80,max:MAX},
	}
	// fmt.Println("Hello",aturanBerat,aturanTinggi)

	fmt.Print("Masukan Tingi anda : ")
	fmt.Scanf("%f",&Tinggi)
	fmt.Println("tinggi anda adalah ",Tinggi)
	counter = 0
	for _,at := range aturanTinggi {
		if(counter == 1){
			break
		}
		if (Tinggi < at.max) && (Tinggi > at.min){

			tempNum = (at.max - Tinggi) / (at.max - (at.max-5))
			tempNum2 = (Tinggi - (at.max-5)) / (at.max - (at.max-5))

			tempKeyTinggi = at.label

			fmt.Println("ada pada label ",at.label)
			fmt.Println("Nilai tempnum adalah ",tempNum)
			fmt.Println("Nilai tempnum2 adalah ",tempNum2)
			counter++
		}
	}

	fmt.Println("----------------------------")
	fmt.Print("Masukan Berat anda : ")
	fmt.Scanf("%f",&Berat)


	counter = 0
	for _,ab := range aturanBerat {
		if(counter == 1){
			break
		}
		if (Berat < ab.max) && (Berat > ab.min){

			tempNum = (ab.max - Berat) / (ab.max - (ab.max-5))
			tempNum2 = (Berat - (ab.max-5)) / (ab.max - (ab.max-5))

			tempKeyBerat = ab.label

			fmt.Println("ada pada label ",ab.label)
			fmt.Println("Nilai tempnum adalah ",tempNum)
			fmt.Println("Nilai tempnum2 adalah ",tempNum2)
			counter++
		}
	}

	fmt.Println("----------------------------")

	for _,at := range aturan{
		if(at.label1 == tempKeyTinggi) && (at.label2 == tempKeyBerat){
			fmt.Println("Hasilnya adalah ",at.hasil)
		}
	}

	_=tempNum
	_=Tinggi
	_=Berat
}