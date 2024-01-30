package main

import (
	"fmt"
	"strings"
	"time"
)

type Studio struct {
	Nama       string
	Kapasitas  int
	HargaPerJam int
	Status     string
}

type Pemesanan struct {
	Tanggal    string
	Durasi     time.Duration
	Penyewa    string
	AlatMusik  string // Menambahkan field untuk menyimpan nama alat musik yang disewa
	TotalHarga int
}

// Menambahkan fitur "sewa alat musik" pada Studio
type AlatMusik struct {
	Nama  string
	Harga int
}

var alatMusikTersedia = map[string]AlatMusik{
	"Gitar":    {"Gitar", 10000},
	"Keyboard": {"Keyboard", 15000},
	"Drum":     {"Drum", 20000},
}

var studios = map[string]Studio{
	"A": {"Studio A", 4, 50000, "Tersedia"},
	"B": {"Studio B", 6, 80000, "Tersedia"},
}

var pemesananAktif = make(map[string]Pemesanan) // Menambahkan map untuk menyimpan pemesanan yang aktif

// Menambahkan satu fitur untuk melihat studio yang tersedia untuk pemesanan
func lihatStudioTersedia() {
	fmt.Println("Studio Tersedia:")
	fmt.Println("----------------------------")
	for key, studio := range studios {
		if studio.Status == "Tersedia" {
			fmt.Printf("Studio %s - Kapasitas: %d orang\n", key, studio.Kapasitas)
		}
	}
	fmt.Println("----------------------------")
}

// Menambahkan satu fitur untuk melihat alat musik yang tersedia
func lihatAlatMusikTersedia() {
	fmt.Println("Alat Musik Tersedia:")
	fmt.Println("----------------------------")
	for key, alat := range alatMusikTersedia {
		fmt.Printf("%s - Harga: Rp %d\n", key, alat.Harga)
	}
	fmt.Println("----------------------------")
}

// Menambahkan satu fitur untuk melihat barang yang disewa
func lihatBarangSewaan() {
	fmt.Println("Barang yang Disewa:")
	fmt.Println("----------------------------")
	for _, pemesanan := range pemesananAktif {
		fmt.Printf("Penyewa: %s\nAlat: %s\nDurasi: %d hari\n", pemesanan.Penyewa, pemesanan.AlatMusik, int(pemesanan.Durasi.Hours()))
		fmt.Println("----------------------------")
	}
}

func main() {
	menu()
}

func menu() {
	for {
		fmt.Println("=== Menu ===")
		fmt.Println("1. Lihat Informasi Studio")
		fmt.Println("2. Booking Studio")
		fmt.Println("3. Sewa Alat Musik")
		fmt.Println("4. Lihat Barang Sewaan")
		fmt.Println("5. Keluar")

		fmt.Print("Pilih menu (1-5): ")
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			lihatInformasiStudio()
		case 2:
			pesanStudio()
		case 3:
			sewaAlatMusik()
		case 4:
			lihatBarangSewaan()
		case 5:
			fmt.Println("Terima kasih sudah menggunakan Studio kami, sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan menu tidak valid.")
		}
	}
}

func lihatInformasiStudio() {
	// Menambahkan pemanggilan fungsi lihatStudioTersedia
	lihatStudioTersedia()

	fmt.Println("Studio Information:")
	fmt.Println("----------------------------")
	for _, studio := range studios {
		fmt.Printf("Studio: %s\nKapasitas: %d orang\nHarga per jam: Rp %d\nStatus: %s\n", studio.Nama, studio.Kapasitas, studio.HargaPerJam, studio.Status)
		fmt.Println("----------------------------")
	}
}

func pesanStudio() {
	// ...
}

// Menambahkan fungsi untuk menyewa alat musik
func sewaAlatMusik() {
	lihatAlatMusikTersedia()

	fmt.Println("Sewa Alat Musik:")
	var pilihanAlat string
	for {
		fmt.Print("Pilih Alat Musik: ")
		fmt.Scanln(&pilihanAlat)
		pilihanAlat = strings.Title(strings.ToLower(pilihanAlat))
		if _, exists := alatMusikTersedia[pilihanAlat]; exists {
			break
		}
		fmt.Println("Pilihan Alat Musik tidak valid.")
	}

	fmt.Printf("Masukkan durasi penyewaan alat musik (dalam hari) : ")
	var durasiSewa int
	fmt.Scanln(&durasiSewa)

	fmt.Printf("Masukkan nama penyewa: ")
	var namaPenyewa string
	fmt.Scanln(&namaPenyewa)

	// Menyimpan informasi pemesanan yang aktif
	pemesanan := Pemesanan{
		Tanggal:    time.Now().Format("02/01/2006"),
		Durasi:     time.Duration(durasiSewa) * time.Hour,
		Penyewa:    namaPenyewa,
		AlatMusik:  pilihanAlat,
		TotalHarga: alatMusikTersedia[pilihanAlat].Harga * durasiSewa,
	}

	pemesananAktif[namaPenyewa] = pemesanan

	fmt.Printf("Total Harga Sewa Alat Musik: Rp %d\n", pemesanan.TotalHarga)
	fmt.Println("----------------------------")
}

