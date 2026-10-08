package main

import (
	"fmt"
	"strings"
)

func main() {
	//soal 1
	kalimat := "halo halo bandung"
	angka := 2021
	var find = "halo"
	var replace = "Hi"

	newText := strings.Replace(kalimat, find, replace, -1)
	fmt.Println(newText + " " + fmt.Sprintf("%d", angka))

	//soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	var indeksJohn string
	if nilaiJohn >= 80 {
		indeksJohn = "A"
	} else if nilaiJohn >= 70 {
		indeksJohn = "B"
	} else if nilaiJohn >= 60 {
		indeksJohn = "C"
	} else if nilaiJohn >= 50 {
		indeksJohn = "D"
	} else {
		indeksJohn = "E"
	}

	var indeksDoe string
	if nilaiDoe >= 80 {
		indeksDoe = "A"
	} else if nilaiDoe >= 70 {
		indeksDoe = "B"
	} else if nilaiDoe >= 60 {
		indeksDoe = "C"
	} else if nilaiDoe >= 50 {
		indeksDoe = "D"
	} else {
		indeksDoe = "E"
	}

	fmt.Println("Indeks Nilai John:", indeksJohn)
	fmt.Println("Indeks Nilai Doe:", indeksDoe)

	//soal 3
	var tanggal = 20
	var bulan = 5
	var tahun = 2000

	var namaBulan string
	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	}

	fmt.Printf("%d %s %d\n", tanggal, namaBulan, tahun)

	//soal 4
	tahunLahir := 2000
	var generasi string

	if tahunLahir >= 1944 && tahunLahir <= 1964 {
		generasi = "Baby boomer"
	} else if tahunLahir >= 1965 && tahunLahir <= 1979 {
		generasi = "Generasi X"
	} else if tahunLahir >= 1980 && tahunLahir <= 1994 {
		generasi = "Generasi Y (Millenials)"
	} else if tahunLahir >= 1995 && tahunLahir <= 2015 {
		generasi = "Generasi Z"
	} else {
		generasi = "Generasi Tidak Terdefinisi"
	}

	fmt.Println("Generasi:", generasi)

	//soal 5
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%2 != 0 {
			fmt.Printf("%d - I Love Coding\n", i)
		} else if i%2 == 0 {
			fmt.Printf("%d - Berkualitas\n", i)
		} else {
			fmt.Printf("%d - Santai\n", i)
		}
	}
}
