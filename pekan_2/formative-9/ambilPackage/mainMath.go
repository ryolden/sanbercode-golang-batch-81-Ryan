package mainMath

import (
	"fmt"
	mtk "formative-9/ambilPackage/matematika"
	"math"
	"strings"
)

// Struct soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

// Interface soal 1
type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

// Function soal 1
// Method untuk segitigaSamaSisi (hitungBangunDatar)
func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

// Method untuk persegiPanjang (hitungBangunDatar)
func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

// Method untuk tabung (hitungBangunRuang)
func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

// Method untuk balok (hitungBangunRuang)
func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return float64(2 * (b.panjang*b.lebar + b.panjang*b.tinggi + b.lebar*b.tinggi))
}

// Struct soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

// Interface soal 2
type phoneDetailer interface {
	displayData()
}

// Function soal 2
// Method untuk mengimplementasikan interface phoneDetailer
func (p phone) displayData() {
	fmt.Printf("name : %s\n", p.name)
	fmt.Printf("brand : %s\n", p.brand)
	fmt.Printf("year : %d\n", p.year)
	fmt.Printf("colors : %s\n", strings.Join(p.colors, ", "))
}

// Function soal 3
func luasPersegi(sisi int, tampilkanKalimat bool) interface{} {
	if sisi == 0 {
		if tampilkanKalimat {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	luas := sisi * sisi
	if tampilkanKalimat {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}

	return luas
}

func main() {
	//soal 1
	var bd1 hitungBangunDatar = segitigaSamaSisi{alas: 6, tinggi: 8}
	var bd2 hitungBangunDatar = persegiPanjang{panjang: 10, lebar: 5}

	fmt.Println("1. Segitiga Sama Sisi")
	fmt.Printf("   Luas     : %d\n", bd1.luas())
	fmt.Printf("   Keliling : %d\n\n", bd1.keliling())

	fmt.Println("2. Persegi Panjang")
	fmt.Printf("   Luas     : %d\n", bd2.luas())
	fmt.Printf("   Keliling : %d\n\n", bd2.keliling())

	var br1 hitungBangunRuang = tabung{jariJari: 7, tinggi: 10}
	var br2 hitungBangunRuang = balok{panjang: 4, lebar: 3, tinggi: 5}

	fmt.Println("1. Tabung")
	fmt.Printf("   Volume         : %.2f\n", br1.volume())
	fmt.Printf("   Luas Permukaan : %.2f\n\n", br1.luasPermukaan())

	fmt.Println("2. Balok")
	fmt.Printf("   Volume         : %.2f\n", br2.volume())
	fmt.Printf("   Luas Permukaan : %.2f\n", br2.luasPermukaan())

	//soal 2
	var myPhone phoneDetailer = phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	myPhone.displayData()

	//soal 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	//soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	strPrefix := prefix.(string)
	slice1 := kumpulanAngkaPertama.([]int)
	slice2 := kumpulanAngkaKedua.([]int)

	semuaAngka := append(slice1, slice2...)

	total := 0
	var teksAngka []string

	for _, angka := range semuaAngka {
		total += angka
		teksAngka = append(teksAngka, fmt.Sprintf("%d", angka))
	}

	fmt.Printf("%s%s = %d\n", strPrefix, strings.Join(teksAngka, " + "), total)

	//soal 5
	a, b := 10, 5

	totalPenjumlahan := mtk.Tambah(a, b)
	totalPerkalian := mtk.Kali(a, b)

	fmt.Printf("Hasil penjumlahan %d + %d = %d\n", a, b, totalPenjumlahan)
	fmt.Printf("Hasil perkalian %d * %d = %d\n", a, b, totalPerkalian)
}
