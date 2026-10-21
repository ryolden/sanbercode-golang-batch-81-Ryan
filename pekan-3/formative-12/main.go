package main

import (
	"encoding/json"
	"net/http"
)

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string `json:"indeks_nilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

type InputNilaiMahasiswa struct {
	Nama       string `json:"nama"`
	MataKuliah string `json:"mata_kuliah"`
	Nilai      uint   `json:"nilai"`
}

func hitungIndeks(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	case nilai >= 50:
		return "D"
	default:
		return "E"
	}
}

func basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != "admin" || password != "admin123" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func postNilaiHandler(w http.ResponseWriter, r *http.Request) {
	var input InputNilaiMahasiswa
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	if input.Nilai > 100 {
		http.Error(w, "Nilai tidak boleh lebih dari 100", http.StatusBadRequest)
		return
	}

	idBaru := uint(len(nilaiNilaiMahasiswa) + 1)
	indeks := hitungIndeks(input.Nilai)

	dataBaru := NilaiMahasiswa{
		ID:          idBaru,
		Nama:        input.Nama,
		MataKuliah:  input.MataKuliah,
		Nilai:       input.Nilai,
		IndeksNilai: indeks,
	}

	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, dataBaru)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dataBaru)
}

func getNilaiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
}

func main() {
	http.HandleFunc("/nilai", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			basicAuth(postNilaiHandler)(w, r)
		case http.MethodGet:
			getNilaiHandler(w, r)
		default:
			http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
