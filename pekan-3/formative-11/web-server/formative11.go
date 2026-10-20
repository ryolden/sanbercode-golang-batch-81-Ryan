package main

import (
	"fmt"
	"math"
	"net/http"
)

// soal 4
func luasAlas(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}

func kelilingAlas(r float64) float64 {
	return 2 * math.Pi * r
}

func volumeTabung(r, t float64) float64 {
	return luasAlas(r) * t
}

func handler(w http.ResponseWriter, r *http.Request) {
	var rVal float64 = 7
	var tVal float64 = 10

	luas := luasAlas(rVal)
	keliling := kelilingAlas(rVal)
	volume := volumeTabung(rVal, tVal)

	responseText := fmt.Sprintf("jariJari : %.0f, tinggi: %.0f, volume : %.2f, luas alas: %.2f, keliling alas: %.2f", rVal, tVal, volume, luas, keliling)
	fmt.Fprint(w, responseText)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
