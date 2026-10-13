package main

import (
	"fmt"
)

func main() {
	//soal 1
	var sentence string

	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	//soal 2
	listBuah := []string{
		"Jeruk",
		"Semangka",
		"Mangga",
		"Strawberry",
		"Durian",
		"Manggis",
		"Alpukat",
	}
	// pointer penambahan buah
	for _, b := range listBuah {
		buahBuahan(&daftarBuah, b)
	}
	// looping urutan buah
	for i, b := range daftarBuah {
		fmt.Printf("%d. %s\n", i+1, b)
	}

	//soal 3
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)
	//menampilkan data film
	for i, film := range dataFilm {
		fmt.Printf("%d. judul : %s\n", i+1, film["judul"])
		fmt.Printf("   durasi : %s\n", film["durasi"])
		fmt.Printf("   genre  : %s\n", film["genre"])
		fmt.Printf("   tahun  : %s\n\n", film["tahun"])
	}

	//soal 4
	var arr = [5]int{177, 153, 190, 199, 184}
	fmt.Println("Sebelum diubah:", arr)
	ubahGenap(&arr)
	fmt.Println("Setelah diubah:", arr)

	//soal 5
	var buah1 = buah{
		nama:       "Nanas",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      9000.00,
	}
	fmt.Printf("Buah 1:\n")
	fmt.Printf("Nama: %s\n", buah1.nama)
	fmt.Printf("Warna: %s\n", buah1.warna)
	fmt.Printf("Ada Bijinya: %t\n", buah1.adaBijinya)
	fmt.Printf("Harga: %.2f\n", buah1.harga)

	var buah2 = buah{
		nama:       "Jeruk",
		warna:      "Jingga",
		adaBijinya: true,
		harga:      8000.00,
	}
	fmt.Printf("Buah 2:\n")
	fmt.Printf("Nama: %s\n", buah2.nama)
	fmt.Printf("Warna: %s\n", buah2.warna)
	fmt.Printf("Ada Bijinya: %t\n", buah2.adaBijinya)
	fmt.Printf("Harga: %.2f\n", buah2.harga)

	var buah3 = buah{
		nama:       "Semangka",
		warna:      "Hijau & Merah",
		adaBijinya: true,
		harga:      10000.00,
	}
	fmt.Printf("Buah 3:\n")
	fmt.Printf("Nama: %s\n", buah3.nama)
	fmt.Printf("Warna: %s\n", buah3.warna)
	fmt.Printf("Ada Bijinya: %t\n", buah3.adaBijinya)
	fmt.Printf("Harga: %.2f\n", buah3.harga)

	var buah4 = buah{
		nama:       "Pisang",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      5000.00,
	}
	fmt.Printf("Buah 4:\n")
	fmt.Printf("Nama: %s\n", buah4.nama)
	fmt.Printf("Warna: %s\n", buah4.warna)
	fmt.Printf("Ada Bijinya: %t\n", buah4.adaBijinya)
	fmt.Printf("Harga: %.2f\n", buah4.harga)
}

// function soal 1
func introduce(sentence *string, name, gender, job, age string) {
	if gender == "laki-laki" {
		*sentence = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", name, job, age)
		return
	}

	*sentence = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", name, job, age)
}

// function soal 2
var daftarBuah = []string{}

func buahBuahan(slice *[]string, namaBuah string) {
	*slice = append(*slice, namaBuah)

}

// func soal 3
var dataFilm = []map[string]string{}

func tambahDataFilm(judul, durasi, genre, tahun string, data *[]map[string]string) {
	film := map[string]string{
		"judul":  judul,
		"durasi": durasi,
		"genre":  genre,
		"tahun":  tahun,
	}
	*data = append(*data, film)
}

// function soal 4
func ubahGenap(arr *[5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
}

// function soal 5
type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      float64
}
