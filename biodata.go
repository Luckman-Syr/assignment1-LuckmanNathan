package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var temanKelas = []Teman{
	{1, "John Doe", "Jakarta", "Software Engineer", "Ingin mempelajari bahasa pemrograman Go."},
	{2, "Jane Smith", "Bandung", "Data Scientist", "Tertarik dengan kecepatan dan efisiensi Go."},
	{3, "David Johnson", "Surabaya", "Web Developer", "Mendapat rekomendasi kelas dari rekan kerja."},
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run biodata.go [no absen]")
		return
	}

	getTemanByAbsen(args[1])

}

func getTemanByAbsen(absen string) *Teman {

	for _, teman := range temanKelas {
		if fmt.Sprintf("%d", teman.Absen) == absen {
			fmt.Println("Data Teman :")
			fmt.Println("Absen \t\t\t\t:", teman.Absen)
			fmt.Println("Nama \t\t\t\t:", teman.Nama)
			fmt.Println("Alamat \t\t\t\t:", teman.Alamat)
			fmt.Println("Pekerjaan \t\t\t:", teman.Pekerjaan)
			fmt.Println("Alasan Memilih Kelas Golang \t:", teman.Alasan)
			return &teman
		}
	}

	fmt.Println("Tidak ditemukan teman dengan absen", absen)
	return nil
}
