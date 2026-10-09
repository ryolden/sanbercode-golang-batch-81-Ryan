package main

import "fmt"

// Function untuk Soal 5
func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang, lebar, tinggi int) int {
	return panjang * lebar * tinggi
}

func main() {
	// SOAL 1
	for i := 1; i <= 7; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	// SOAL 2
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	fmt.Println(kalimat[2:])

	// SOAL 3
	var sayuran = []string{}

	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i, v := range sayuran {
		fmt.Printf("%d. %s\n", i+1, v)
	}

	// SOAL 4
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	volume := 1

	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)
		volume *= value
	}

	fmt.Printf("Volume Balok = %d\n", volume)

	// SOAL 5
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volumeBalok := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volumeBalok)
}
