package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

// Function soal 1
func introduce(nama, jenisKelamin, pekerjaan, usia string) (result string) {
	sebutan := "Pak"
	if jenisKelamin == "perempuan" {
		sebutan = "Bu"
	}

	result = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", sebutan, nama, pekerjaan, usia)
	return
}

// Function soal 2
func buahFavorit(nama string, buah ...string) string {
	var buahDikutip []string
	for _, item := range buah {
		buahDikutip = append(buahDikutip, fmt.Sprintf(`"%s"`, item))
	}

	daftarBuah := strings.Join(buahDikutip, ", ")
	return fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah %s", nama, daftarBuah)
}

// Function soal 4
func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 1; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

// function soal 5
func hitungLingkaran(luas *float64, keliling *float64, r float64) {
	*luas = math.Pi * math.Pow(r, 2)
	*keliling = 2 * math.Pi * r
}

func main() {
	//soal 1
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Output: Pak John adalah seorang penulis yang berusia 30 tahun

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Output: Bu Sarah adalah seorang model yang berusia 28 tahun

	//soal 2
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)

	//soal 3
	var dataFilm = []map[string]string{}

	// Closure function untuk menambahkan data film ke slice dataFilm
	tambahDataFilm := func(title, jam, genre, tahun string) {
		film := map[string]string{
			"title": title,
			"jam":   jam,
			"genre": genre,
			"tahun": tahun,
		}
		dataFilm = append(dataFilm, film)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

	//soal 4
	fmt.Println(factorial(5))
	fmt.Println(factorial(7))
	fmt.Println(factorial(30))

	//soal 5
	var luasLingkaran float64
	var kelilingLingkaran float64
	var jariJari float64 = 7.0

	// Mengirimkan alamat memori variabel (&)
	hitungLingkaran(&luasLingkaran, &kelilingLingkaran, jariJari)

	fmt.Printf("Luas Lingkaran    : %.2f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran: %.2f\n", kelilingLingkaran)
}
