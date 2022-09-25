package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {

	// get input Argument
	var argsRaw = os.Args
	var num = argsRaw[1]

	// convert string menjadi int
	number, _ := strconv.Atoi(num)
	// fmt.Printf("-> %#v\n", number)

	// data siswa
	var dataSiswa = []Person{
		{Nama: "Evan", Alamat: "Jl telaga bogor", Pekerjaan: "Mahasiswa", Alasan: "Ingin tahu lebih banyak"},
		{Nama: "Faqih", Alamat: "Jl perkutut bintaro", Pekerjaan: "Mahasiswa", Alasan: "Supaya jadi programmer handal"},
		{Nama: "Alfirman", Alamat: "Jl Bulak indah jakbar", Pekerjaan: "Mahasiswa", Alasan: "Mau jadi backend engineer"},
		{Nama: "Indra", Alamat: "Jl Paal V Pondok indah", Pekerjaan: "Mahasiswa", Alasan: "Modal buat cari kerja"},
		{Nama: "Fahrizal", Alamat: "Jl Kasang Ciledug", Pekerjaan: "PNS", Alasan: "Memperdalam ilmu coding"},
		{Nama: "Rafli", Alamat: "Jl Kumpeh Banten", Pekerjaan: "Mahasiswa", Alasan: "Dapetin skill baru"},
		{Nama: "Samsul", Alamat: "Jl Senopati Jakarta", Pekerjaan: "Pengangguran", Alasan: "Nyari kesibukan"},
		{Nama: "Eko", Alamat: "Jl Kenangan Bandung barat", Pekerjaan: "Mahasiswa", Alasan: "Diajak teman"},
		{Nama: "Iwan", Alamat: "Jl Sudirman", Pekerjaan: "Mahasiswa", Alasan: "Disuruh orang tua"},
		{Nama: "Tedi", Alamat: "Jl Kebangsaan", Pekerjaan: "Freelancer", Alasan: "Dapetin sertifikat"},
	}

	// check data jika tidak ada
	if number > len(dataSiswa) {
		showError()
	} else {
		// jika data siswa ada panggil fungsi showdata
		showData(dataSiswa[number-1])
	}
}

func showData(siswa Person) {
	fmt.Println("Nama      : \t", siswa.Nama)
	fmt.Println("Alamat    : \t", siswa.Alamat)
	fmt.Println("Pekerjaan : \t", siswa.Pekerjaan)
	fmt.Println("Alasan    : \t", siswa.Alasan)
}

func showError() {
	fmt.Println("No Absen Tidak Ditemukan")
}

// Tugas
// Contohnya, ketika kalian menjalankan perintah go run biodata.go 1 maka data yang akan muncul adalah
// teman kalian dengan absen no 1. Data yang harus ditampilkan yaitu:
// ●Nama
// ●Alamat
// ●Pekerjaan
// ●Alasan memilih kelas Golang Gunakanlah struct dan function untuk menampilkan data tersebut.
// *Kalian bisa menggunakan os.Args untuk mendapatkan argument pada terminal
