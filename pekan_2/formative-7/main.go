package main

import (
	"fmt"
)

func main() {
	//soal 1
	type mahasiswa struct {
		nama string
		nim  string
		usia int
	}
	var mhs = mahasiswa{
		nama: "asep",
		nim:  "29901",
		usia: 20,
	}
	fmt.Println("Nama: ", mhs.nama)
	fmt.Println("NIM: ", mhs.nim)
	fmt.Println("Usia: ", mhs.usia)

	//soal 2
	type segitiga struct {
		alas, tinggi int
	}
	var sgtg = segitiga{
		alas:   13,
		tinggi: 19,
	}
	fmt.Println("Alas: ", sgtg.alas)
	fmt.Println("Tinggi: ", sgtg.tinggi)

	type persegi struct {
		sisi int
	}
	var prsg = persegi{
		sisi: 5,
	}
	fmt.Println("Sisi: ", prsg.sisi)

	type persegiPanjang struct {
		panjang, lebar int
	}
	var prsgPjg = persegiPanjang{
		panjang: 7,
		lebar:   9,
	}
	fmt.Println("Panjang: ", prsgPjg.panjang)
	fmt.Println("Lebar: ", prsgPjg.lebar)

	//soal 3
	myPhone := phone{
		name:   "Samsung galax bgt",
		brand:  "Samsung",
		colors: []string{"black", "yellow"},
	}
	fmt.Println("myPhone colors before: ", myPhone.colors)
	//ubah warna
	myPhone.addColor("blue")
	fmt.Println("myPhone colors after: ", myPhone.colors)

	//soal 4
	var dataFilm = []movies{}
	tambahDataFilm("LOTR", 2, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 2, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 2, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 2, "horror", 2004, &dataFilm)

	//loop menampilkan data film
	for i, movie := range dataFilm {
		fmt.Printf("%d. judul : %s\n", i+1, movie.title)
		fmt.Printf("   durasi : %d jam\n", movie.duration)
		fmt.Printf("   genre  : %s\n", movie.genre)
		fmt.Printf("   tahun  : %d\n\n", movie.year)
	}

	//soal 5
	kucing := Kucing{}
	anjing := Anjing{}

	// Pemanggilan langsung melalui fungsi penerima interface
	cetakSuara(kucing)
	cetakSuara(anjing)
}

// struct untuk soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

// pointer untuk menambahkan warna baru ke dalam slice colors soal 3
func (p *phone) addColor(color string) {
	p.colors = append(p.colors, color)
}

// struct untuk soal 4
type movies struct {
	title, genre   string
	duration, year int
}

// func pointer untuk menambahkan film baru ke dalam no 4
func tambahDataFilm(title string, duration int, genre string, year int, data *[]movies) {
	newMovie := movies{
		title:    title,
		duration: duration,
		genre:    genre,
		year:     year,
	}
	*data = append(*data, newMovie)
}

// interface untuk soal 5
type Hewan interface {
	Suara() string
}

type Kucing struct{}

func (k Kucing) Suara() string {
	return "Meong"
}

type Anjing struct{}

func (a Anjing) Suara() string {
	return "Woof"
}

func cetakSuara(h Hewan) {
	fmt.Println(h.Suara())
}
