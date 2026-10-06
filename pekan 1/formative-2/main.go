package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//soal 1
	fmt.Println("Bootcamp")

	var text string = "Digital"
	fmt.Println(text)

	text1 := "Skill"
	fmt.Println(text1)

	text2 := "academy"
	text2 = "Sanbercode"
	fmt.Println(text2)

	var text3 string = "Golang"
	fmt.Println(text3)

	//soal 2
	halo := "Halo Dunia"
	var find = "Dunia"
	var replace = "Golang"

	newText := strings.Replace(halo, find, replace, -1)
	fmt.Println(newText)

	//soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	var findS = "s"
	var replaceS = "S"

	var findR = "r"
	var replaceR = "R"

	var Golang = strings.ToUpper(kataKeempat)

	kataKedua = strings.Replace(kataKedua, findS, replaceS, 1)
	kataKetiga = strings.Replace(kataKetiga, findR, replaceR, 1)

	println(kataPertama + " " + kataKedua + " " + kataKetiga + " " + Golang)

	//soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var n1, _ = strconv.Atoi(angkaPertama)
	var n2, _ = strconv.Atoi(angkaKedua)
	var n3, _ = strconv.Atoi(angkaKetiga)
	var n4, _ = strconv.Atoi(angkaKeempat)
	fmt.Println(n1 + n2 + n3 + n4)

	//soal 5
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var p, _ = strconv.Atoi(panjangPersegiPanjang)
	var l, _ = strconv.Atoi(lebarPersegiPanjang)
	var a, _ = strconv.Atoi(alasSegitiga)
	var t, _ = strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang = p * l
	var kelilingPersegiPanjang = 2 * (p + l)
	var luasSegitiga = 0.5 * float32(a) * float32(t)

	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga:", luasSegitiga)
}
