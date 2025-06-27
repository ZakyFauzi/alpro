package main

import "fmt"

// --- Tipe Bentukan (Struct) ---
type Aktivitas struct {
	ID           int
	Kategori     string
	Deskripsi    string
	DampakKarbon float64
	Frekuensi    int
}

const MAX_AKTIVITAS = 100

var daftarAktivitas [MAX_AKTIVITAS]Aktivitas
var jumlahAktivitas int

// --- Fungsi Bantu (untuk String, karena tanpa "strings" package) ---
// Mengonversi string ke huruf kecil
func toLower(s string) string {
	var result string = ""
	var i int
	var c byte

	for i = 0; i < len(s); i++ {
		c = s[i]
		// Jika karakter adalah huruf besar (A-Z)
		if c >= 'A' && c <= 'Z' {
			// Ubah ke huruf kecil dengan menambahkan 32
			c += 32
		}
		result = result + string(c)
	}
	return result
}

// --- Pencarian: Cari ID Aktivitas (untuk internal) ---
func findAktivitasIndexByID(id int) int {
	var i int
	for i = 0; i < jumlahAktivitas; i++ {
		if daftarAktivitas[i].ID == id {
			return i
		}
	}
	return -1
}

// --- Menu: 1. Tambah Aktivitas ---
func tambahAktivitas(id int, kategori, deskripsi string, dampak float64, frekuensi int) {
	var i_cek int
	var aktivitasBaru Aktivitas

	if jumlahAktivitas >= MAX_AKTIVITAS {
		fmt.Println("Maaf, kapasitas aktivitas sudah penuh (maksimal 100).")

	}

	for i_cek = 0; i_cek < jumlahAktivitas; i_cek++ {
		if daftarAktivitas[i_cek].ID == id {
			fmt.Println("Gagal menambahkan: ID sudah ada. Gunakan ID lain atau edit aktivitas yang sudah ada.")

		}
	}

	aktivitasBaru = Aktivitas{
		ID:           id,
		Kategori:     kategori,
		Deskripsi:    deskripsi,
		DampakKarbon: dampak,
		Frekuensi:    frekuensi,
	}

	daftarAktivitas[jumlahAktivitas] = aktivitasBaru
	jumlahAktivitas++
	fmt.Println("Aktivitas berhasil ditambahkan!")
}

// --- Menu: 2. Cari Aktivitas (Sequential & Binary Search) ---

func cariSequential(kategoriCari string) []int {
	var hasilIndeks []int
	var kategoriCariLower string
	kategoriCariLower = toLower(kategoriCari)
	var i int
	for i = 0; i < jumlahAktivitas; i++ {
		if toLower(daftarAktivitas[i].Kategori) == kategoriCariLower {
			hasilIndeks = append(hasilIndeks, i)
		}
	}
	return hasilIndeks
}

func selectionSortKategori() {
	var i int
	for i = 0; i < jumlahAktivitas-1; i++ {
		var minIndex int
		minIndex = i
		var j int
		for j = i + 1; j < jumlahAktivitas; j++ {
			if toLower(daftarAktivitas[j].Kategori) < toLower(daftarAktivitas[minIndex].Kategori) {
				minIndex = j
			}
		}
		daftarAktivitas[i], daftarAktivitas[minIndex] = daftarAktivitas[minIndex], daftarAktivitas[i]
	}
}

func cariBinary(kategoriCari string) []int {
	var hasilIndeks []int
	var kategoriCariLower string
	var low int
	var high int
	var firstOccurenceIndex int
	var i_loop int
	var mid int

	kategoriCariLower = toLower(kategoriCari)
	low = 0
	high = jumlahAktivitas - 1
	firstOccurenceIndex = -1
	for low <= high {
		mid = (low + high) / 2
		var currentKategoriLower string
		currentKategoriLower = toLower(daftarAktivitas[mid].Kategori)

		if currentKategoriLower == kategoriCariLower {
			firstOccurenceIndex = mid
			high = mid - 1
		} else if currentKategoriLower < kategoriCariLower {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if firstOccurenceIndex == -1 {
		return hasilIndeks
	}

	hasilIndeks = append(hasilIndeks, firstOccurenceIndex)

	for i_loop = firstOccurenceIndex + 1; i_loop < jumlahAktivitas; i_loop++ {
		if toLower(daftarAktivitas[i_loop].Kategori) == kategoriCariLower {
			hasilIndeks = append(hasilIndeks, i_loop)
		}
	}
	return hasilIndeks
}

// --- Menu: 3. Edit Aktivitas & Hapus Aktivitas ---
func editAktivitas(id int, kategori, deskripsi string, dampak float64, frekuensi int, isDelete bool) {
	var foundIndex int
	var i int
	foundIndex = findAktivitasIndexByID(id)

	if foundIndex != -1 {
		if isDelete {
			for i = foundIndex; i < jumlahAktivitas-1; i++ {
				daftarAktivitas[i] = daftarAktivitas[i+1]
			}
			daftarAktivitas[jumlahAktivitas-1] = Aktivitas{}
			jumlahAktivitas--
			fmt.Println("Aktivitas berhasil dihapus!")
		} else {
			daftarAktivitas[foundIndex].Kategori = kategori
			daftarAktivitas[foundIndex].Deskripsi = deskripsi
			daftarAktivitas[foundIndex].DampakKarbon = dampak
			daftarAktivitas[foundIndex].Frekuensi = frekuensi
			fmt.Println("Aktivitas berhasil diupdate!")
		}
	} else {
		fmt.Println("Aktivitas dengan ID tersebut tidak ditemukan.")
	}
}

// --- Menu: 4. Urutkan (Selection Sort & Insertion Sort) ---
func selectionSortDampak(ascending bool) {
	var i int
	var j int
	for i = 0; i < jumlahAktivitas-1; i++ {
		var extremeIndex int
		extremeIndex = i
		for j = i + 1; j < jumlahAktivitas; j++ {
			if ascending {
				if daftarAktivitas[j].DampakKarbon < daftarAktivitas[extremeIndex].DampakKarbon {
					extremeIndex = j
				}
			} else {
				if daftarAktivitas[j].DampakKarbon > daftarAktivitas[extremeIndex].DampakKarbon {
					extremeIndex = j
				}
			}
		}
		daftarAktivitas[i], daftarAktivitas[extremeIndex] = daftarAktivitas[extremeIndex], daftarAktivitas[i]
	}
}

func insertionSortFrekuensi(ascending bool) {
	var i int
	var key Aktivitas
	var j int
	for i = 1; i < jumlahAktivitas; i++ {
		key = daftarAktivitas[i]
		j = i - 1
		if ascending {
			for j >= 0 && daftarAktivitas[j].Frekuensi > key.Frekuensi {
				daftarAktivitas[j+1] = daftarAktivitas[j]
				j = j - 1
			}
		} else {
			for j >= 0 && daftarAktivitas[j].Frekuensi < key.Frekuensi {
				daftarAktivitas[j+1] = daftarAktivitas[j]
				j = j - 1
			}
		}
		daftarAktivitas[j+1] = key
	}
}

// --- Menu: 5. Tampilkan Daftar Aktivitas ---
func tampilkanDaftar() {
	var i int
	var a Aktivitas
	fmt.Println("\n--- Daftar Aktivitas ---")
	if jumlahAktivitas == 0 {
		fmt.Println("Belum ada aktivitas yang ditambahkan.")

	}
	fmt.Printf("%-5s | %-15s | %-20s | %-10s | %-10s\n", "ID", "Kategori", "Deskripsi", "Dampak (kg)", "Frekuensi")
	fmt.Println("----------------------------------------------------------------------")
	for i = 0; i < jumlahAktivitas; i++ {
		a = daftarAktivitas[i]
		fmt.Printf("%-5d | %-15s | %-20s | %-10.2f | %-10d\n", a.ID, a.Kategori, a.Deskripsi, a.DampakKarbon, a.Frekuensi)
	}
}

// --- Fungsi untuk menampilkan hasil pencarian ---
func tampilkanHasilPencarian(hasilIndeks []int) {
	var i int
	var idx int
	var a Aktivitas

	if len(hasilIndeks) == 0 {
		fmt.Println("Aktivitas tidak ditemukan!")
	}
	fmt.Println("Aktivitas ditemukan:")
	fmt.Printf("%-5s | %-15s | %-20s | %-10s | %-10s\n", "ID", "Kategori", "Deskripsi", "Dampak (kg)", "Frekuensi")
	fmt.Println("----------------------------------------------------------------------")

	for i = 0; i < len(hasilIndeks); i++ { // Loop menggunakan i
		idx = hasilIndeks[i] // Ambil nilai indeks dari slice hasilIndeks
		a = daftarAktivitas[idx]
		fmt.Printf("%-5d | %-15s | %-20s | %-10.2f | %-10d\n", a.ID, a.Kategori, a.Deskripsi, a.DampakKarbon, a.Frekuensi)
	}
}

// --- Menu: 6. Laporan Bulanan & Rekomendasi ---
func laporanBulanan() {
	var totalJejakKarbon float64
	var i int
	var a Aktivitas
	var aktivitasTerbanyakDampak Aktivitas
	var aktivitasTerdikitDampak Aktivitas
	var skor float64

	if jumlahAktivitas == 0 {
		fmt.Println("Belum ada aktivitas untuk membuat laporan.")

	}

	fmt.Println("\n--- Laporan Bulanan Jejak Karbon ---")

	totalJejakKarbon = 0.0
	aktivitasTerbanyakDampak = Aktivitas{DampakKarbon: -1.0}
	aktivitasTerdikitDampak = Aktivitas{DampakKarbon: 999999999.0}

	for i = 0; i < jumlahAktivitas; i++ {
		a = daftarAktivitas[i]
		totalJejakKarbon += a.DampakKarbon * float64(a.Frekuensi)

		if a.DampakKarbon > aktivitasTerbanyakDampak.DampakKarbon {
			aktivitasTerbanyakDampak = a
		}
		if a.DampakKarbon < aktivitasTerdikitDampak.DampakKarbon {
			aktivitasTerdikitDampak = a
		}
	}

	fmt.Printf("Total Jejak Karbon Bulan Ini: %.2f kg CO2e\n", totalJejakKarbon)

	if aktivitasTerbanyakDampak.DampakKarbon != -1.0 {
		fmt.Printf("Aktivitas dengan Dampak Karbon Terbesar: '%s' (%.2f kg CO2e)\n", aktivitasTerbanyakDampak.Deskripsi, aktivitasTerbanyakDampak.DampakKarbon)
	}
	if aktivitasTerdikitDampak.DampakKarbon != 999999999.0 {
		fmt.Printf("Aktivitas dengan Dampak Karbon Terkecil: '%s' (%.2f kg CO2e)\n", aktivitasTerdikitDampak.Deskripsi, aktivitasTerdikitDampak.DampakKarbon)
	}

	skor = 100.0 - (totalJejakKarbon / 100)
	if skor < 0 {
		skor = 0
	} else if skor > 100 {
		skor = 100
	}
	fmt.Printf("Skor Keberlanjutan Anda: %.2f/100\n", skor)

	fmt.Println("\nRekomendasi untuk Mengurangi Jejak Karbon:")
	if totalJejakKarbon > 500 {
		fmt.Println("- Pertimbangkan untuk mengurangi perjalanan menggunakan kendaraan pribadi.")
		fmt.Println("- Fokus pada makanan nabati dan lokal.")
	} else if totalJejakKarbon > 200 {
		fmt.Println("- Kurangi konsumsi listrik (matikan lampu yang tidak perlu, cabut charger).")
		fmt.Println("- Pilah sampah dengan lebih baik dan kurangi penggunaan plastik sekali pakai.")
	} else {
		fmt.Println("- Terus pertahankan gaya hidup ramah lingkungan Anda!")
		fmt.Println("- Edukasi orang lain tentang pentingnya keberlanjutan.")
	}

	fmt.Println("\nSistem memberikan skor keberlanjutan berdasarkan pola hidup pengguna. Skor Anda dihitung dari total jejak karbon bulanan. Semakin rendah jejak karbon, semakin tinggi skornya.")
}

// --- Main Program (Menu Interaktif) ---
func main() {
	var pilihan, subPilihan, id, frekuensi int
	var kategori, deskripsi string
	var dampak float64
	var ascending bool

	daftarAktivitas[0] = Aktivitas{ID: 1, Deskripsi: "Motor", Kategori: "Transportasi", DampakKarbon: 2.3, Frekuensi: 20}
	daftarAktivitas[1] = Aktivitas{ID: 2, Deskripsi: "Kereta", Kategori: "Transportasi", DampakKarbon: 1.0, Frekuensi: 10}
	daftarAktivitas[2] = Aktivitas{ID: 3, Deskripsi: "Bus", Kategori: "Transportasi", DampakKarbon: 0.5, Frekuensi: 15}
	daftarAktivitas[3] = Aktivitas{ID: 4, Deskripsi: "Ayam", Kategori: "Makanan", DampakKarbon: 1.5, Frekuensi: 30}
	daftarAktivitas[4] = Aktivitas{ID: 5, Deskripsi: "Sayur", Kategori: "Makanan", DampakKarbon: 0.2, Frekuensi: 15}
	daftarAktivitas[5] = Aktivitas{ID: 6, Deskripsi: "Daging", Kategori: "Makanan", DampakKarbon: 3.0, Frekuensi: 5}
	daftarAktivitas[6] = Aktivitas{ID: 7, Deskripsi: "Lampu", Kategori: "Energi", DampakKarbon: -0.2, Frekuensi: 60}
	jumlahAktivitas = 7

	for {
		fmt.Println("\n=== Aplikasi Pelacak Gaya Hidup Ramah Lingkungan ===")
		fmt.Println("1. Tambah Aktivitas")
		fmt.Println("2. Cari Aktivitas")
		fmt.Println("3. Edit Aktivitas")
		fmt.Println("4. Urutkan")
		fmt.Println("5. Tampilkan Daftar Aktivitas")
		fmt.Println("6. Laporan Bulanan")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu (berupa angka, contoh: 1): ")
		fmt.Scan(&pilihan)

		if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan Aplikasi Pelacak Gaya Hidup Ramah Lingkungan!")

		}

		if pilihan == 1 {
			fmt.Print("Masukkan ID (berupa angka, contoh: 1): ")
			fmt.Scan(&id)
			fmt.Print("Masukkan Kategori (satu kata, contoh: Transportasi): ")
			fmt.Scan(&kategori)
			fmt.Print("Masukkan Deskripsi (satu kata, contoh: Mobil): ")
			fmt.Scan(&deskripsi)
			fmt.Print("Masukkan Dampak Karbon (kg CO2 per aktivitas, contoh: 0.5): ")
			fmt.Scan(&dampak)
			fmt.Print("Masukkan Frekuensi (kali per bulan, contoh: 20): ")
			fmt.Scan(&frekuensi)
			tambahAktivitas(id, kategori, deskripsi, dampak, frekuensi)
		} else if pilihan == 2 {
			fmt.Print("Masukkan Kategori yang dicari (satu kata, contoh: Transportasi): ")
			fmt.Scan(&kategori)
			fmt.Println("Pilih metode pencarian:")
			fmt.Println("1. Sequential Search")
			fmt.Println("2. Binary Search")
			fmt.Print("Pilih metode (berupa angka, contoh: 1): ")
			fmt.Scan(&subPilihan)
			if subPilihan == 1 {
				var hasilIndeks []int
				hasilIndeks = cariSequential(kategori)
				tampilkanHasilPencarian(hasilIndeks)
			} else if subPilihan == 2 {
				selectionSortKategori()
				var hasilIndeks []int
				hasilIndeks = cariBinary(kategori)
				tampilkanHasilPencarian(hasilIndeks)
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		} else if pilihan == 3 {
			fmt.Print("Masukkan ID aktivitas (berupa angka, contoh: 1): ")
			fmt.Scan(&id)
			fmt.Println("Pilih aksi:")
			fmt.Println("1. Edit Aktivitas")
			fmt.Println("2. Hapus Aktivitas")
			fmt.Print("Pilih aksi (berupa angka, contoh: 1): ")
			fmt.Scan(&subPilihan)
			if subPilihan == 1 {
				fmt.Print("Masukkan Kategori baru (satu kata, contoh: Transportasi): ")
				fmt.Scan(&kategori)
				fmt.Print("Masukkan Deskripsi baru (satu kata, contoh: Bus): ")
				fmt.Scan(&deskripsi)
				fmt.Print("Masukkan Dampak Karbon baru (contoh: 0.2): ")
				fmt.Scan(&dampak)
				fmt.Print("Masukkan Frekuensi baru (contoh: 15): ")
				fmt.Scan(&frekuensi)
				editAktivitas(id, kategori, deskripsi, dampak, frekuensi, false)
			} else if subPilihan == 2 {
				editAktivitas(id, "", "", 0, 0, true)
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		} else if pilihan == 4 {
			fmt.Println("Pilih kriteria pengurutan:")
			fmt.Println("1. Dampak Karbon (Selection Sort)")
			fmt.Println("2. Frekuensi (Insertion Sort)")
			fmt.Print("Pilih kriteria (berupa angka, contoh: 1): ")
			fmt.Scan(&subPilihan)
			fmt.Print("Urutkan ascending? (true/false, contoh: true): ")
			fmt.Scan(&ascending)
			if subPilihan == 1 {
				selectionSortDampak(ascending)
				tampilkanDaftar()
			} else if subPilihan == 2 {
				insertionSortFrekuensi(ascending)
				tampilkanDaftar()
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		} else if pilihan == 5 {
			tampilkanDaftar()
		} else if pilihan == 6 {
			laporanBulanan()
		} else {
			if pilihan != 0 {
				fmt.Println("Pilihan tidak valid!")
			}
		}
	}
}
