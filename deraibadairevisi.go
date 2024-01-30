package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Studio struct {
	Nama      string
	Kapasitas int
	Harga     float64
	Status    string
}

type Pemesanan struct {
	Tanggal string
	Durasi  int
	Penyewa string
	Studio  Studio
	Total   float64
}

var daftarStudio = map[string]Studio{
	"A": {"Studio A", 4, 50000.00, "Tersedia"},
	"B": {"Studio B", 6, 80000.00, "Tersedia"},
}

var daftarPemesanan = make(map[string]Pemesanan)

func main() {
	fmt.Println("=== Selamat Datang di Studio Derai Badai ===")
	for {
		fmt.Println("=== Menu ===")
		fmt.Println("1. Lihat Informasi Studio")
		fmt.Println("2. Pesan Studio")
		fmt.Println("3. Status Pemesanan")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu (1-4): ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		menu := scanner.Text()

		switch menu {
		case "1":
			lihatInformasiStudio()
		case "2":
			pesanStudio()
		case "3":
			statusPemesanan()
		case "4":
			fmt.Println("Terima kasih. Sampai jumpa!")
			os.Exit(0)
		default:
			fmt.Println("Pilihan menu tidak valid.")
		}
	}
}

func lihatInformasiStudio() {
	fmt.Println("Studio Information:")
	for _, studio := range daftarStudio {
		fmt.Println("----------------------------")
		fmt.Printf("Studio: %s\n", studio.Nama)
		fmt.Printf("Kapasitas: %d orang\n", studio.Kapasitas)
		fmt.Printf("Harga per jam: Rp %.2f\n", studio.Harga)
		fmt.Printf("Status: %s\n", studio.Status)
	}
	fmt.Println("----------------------------")
}

func pesanStudio() {
	fmt.Println("Pesan Studio:")
	fmt.Print("Pilih Studio (A/B): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pilihanStudio := strings.ToUpper(scanner.Text())

	studio, found := daftarStudio[pilihanStudio]
	if !found {
		fmt.Println("Pilihan Studio tidak valid.")
		return
	}

	fmt.Printf("Masukkan tanggal pemesanan (format: 02/01/2006): ")
	scanner.Scan()
	tanggal := scanner.Text()

	fmt.Printf("Masukkan durasi pemesanan (dalam jam): ")
	scanner.Scan()
	durasiStr := scanner.Text()
	durasi, err := strconv.Atoi(durasiStr)
	if err != nil {
		fmt.Println("Durasi tidak valid.")
		return
	}

	fmt.Printf("Masukkan nama penyewa: ")
	scanner.Scan()
	penyewa := scanner.Text()

	pemesanan := Pemesanan{
		Tanggal: tanggal,
		Durasi:  durasi,
		Penyewa: penyewa,
		Studio:  studio,
		Total:   float64(durasi) * studio.Harga,
	}

	// Tambahkan pemesanan ke daftarPemesanan
	kodePemesanan := generateKodePemesanan()
	daftarPemesanan[kodePemesanan] = pemesanan

	fmt.Printf("Pemesanan untuk %s\n", pemesanan.Studio.Nama)
	fmt.Printf("Tanggal: %s\n", pemesanan.Tanggal)
	fmt.Printf("Durasi: %d JAM\n", pemesanan.Durasi)
	fmt.Printf("Penyewa: %s\n", pemesanan.Penyewa)
	fmt.Printf("Total Harga: Rp %.2f\n", pemesanan.Total)

	fmt.Println("----------------------------")
	fmt.Println("Menggunakan integrasi pembayaran...")
	fmt.Print("Pilihan pembayaran (cash/qris): ")

	scanner.Scan()
	pilihanPembayaran := strings.ToLower(scanner.Text())

	if pilihanPembayaran == "qris" {
		fmt.Printf("Silahkan scan kode QRIS yang berada di depan pintu masuk %s.\n", pemesanan.Studio.Nama)
	}

	fmt.Println("Terima kasih atas pemesanan Anda!")
}

func statusPemesanan() {
	fmt.Println("=== Status Pemesanan ===")
	if len(daftarPemesanan) == 0 {
		fmt.Println("Belum ada studio yang dipesan.")
		return
	}

	for kode, pemesanan := range daftarPemesanan {
		fmt.Printf("Kode: %s\n", kode)
		fmt.Printf("Tanggal: %s\n", pemesanan.Tanggal)
		fmt.Printf("Durasi: %d JAM\n", pemesanan.Durasi)
		fmt.Printf("Penyewa: %s\n", pemesanan.Penyewa)
		fmt.Printf("Total Harga: Rp %.2f\n", pemesanan.Total)
		fmt.Println("----------------------------")
	}

	fmt.Print("Masukkan kode pemesanan yang ingin dibatalkan (0 untuk kembali): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	kodePemesanan := scanner.Text()

	if kodePemesanan == "0" {
		return
	}

	// Membatalkan pemesanan
	_, found := daftarPemesanan[kodePemesanan]
	if found {
		delete(daftarPemesanan, kodePemesanan)
		fmt.Println("Pemesanan berhasil dibatalkan.")
	} else {
		fmt.Println("Kode pemesanan tidak valid.")
	}
}

// Fungsi untuk menghasilkan kode unik untuk setiap pemesanan
func generateKodePemesanan() string {
	return strconv.FormatInt(time.Now().UnixNano(), 36)
}
