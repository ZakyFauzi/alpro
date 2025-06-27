# Aplikasi Pelacak Gaya Hidup Ramah Lingkungan

Aplikasi Pelacak Gaya Hidup Ramah Lingkungan adalah aplikasi berbasis konsol yang membantu pengguna memantau jejak karbon dari aktivitas sehari-hari. Aplikasi ini menghitung total jejak karbon bulanan, memberikan skor keberlanjutan, dan menawarkan rekomendasi untuk gaya hidup yang lebih ramah lingkungan. Dibuat sebagai tugas besar mata kuliah algoritma dan pemrograman menggunakan bahasa Go (Golang).

## Fitur Utama

- **Manajemen Aktivitas**: Menambah, mengedit, dan menghapus aktivitas harian yang berdampak pada lingkungan.
- **Perhitungan Jejak Karbon**: Menghitung total jejak karbon bulanan berdasarkan aktivitas yang dicatat.
- **Laporan Bulanan**: Menampilkan total jejak karbon, skor keberlanjutan, serta aktivitas dengan dampak tertinggi dan terendah.
- **Pencarian Aktivitas**:
  - **Sequential Search**: Mencari aktivitas berdasarkan kategori secara berurutan.
  - **Binary Search**: Mencari aktivitas berdasarkan kategori (data diurutkan terlebih dahulu).
- **Pengurutan Data**:
  - **Selection Sort**: Mengurutkan aktivitas berdasarkan dampak karbon.
  - **Insertion Sort**: Mengurutkan aktivitas berdasarkan frekuensi.
- **Struktur Data**:
  - Menggunakan **struct** untuk menyimpan detail aktivitas (ID, kategori, deskripsi, dampak karbon, frekuensi).
  - Menggunakan **array statis** dengan kapasitas 100 aktivitas.
- **Antarmuka Berbasis Teks**: Menu interaktif dengan input numerik untuk navigasi yang mudah.

## Prasyarat

Untuk menjalankan aplikasi ini, kamu memerlukan:
- **Go (Golang)** versi 1.16 atau lebih baru. Unduh di [golang.org](https://golang.org/dl/).
- Terminal atau command prompt untuk menjalankan program.

## Cara Menjalankan

### Opsi 1: Menjalankan di Lokal Device
Jika kamu memiliki file `main.go` di komputer:
1. **Simpan File**: Pastikan file `main.go` tersimpan di direktori tertentu di komputer kamu.
2. **Buka Terminal**:
   - Windows: Buka Command Prompt atau PowerShell.
   - macOS/Linux: Buka Terminal.
3. **Navigasi ke Direktori**: Masuk ke direktori tempat `main.go` tersimpan menggunakan perintah:
   ```bash
   cd /path/to/your/folder
   ```
   Contoh: `cd Documents/TubesNgantuk`
4. **Jalankan Program**:
   ```bash
   go run main.go
   ```
5. **Gunakan Aplikasi**: Menu utama akan muncul di konsol. Pilih opsi dengan memasukkan angka sesuai perintah.

### Opsi 2: Menjalankan dari Repository
Jika kamu ingin mengunduh dari GitHub:
1. **Kloning Repository**:
   ```bash
   git clone https://github.com/FtBdra/TubesNgantuk.git
   cd TubesNgantuk
   ```
2. **Jalankan Program**:
   ```bash
   go run main.go
   ```
3. **Gunakan Aplikasi**: Ikuti petunjuk di konsol untuk mengelola aktivitas atau melihat laporan.

## Struktur Kode

- **`main.go`**: Berisi logika utama, termasuk menu interaktif dan pemanggilan fungsi.
- **Struktur Data**:
  - Struct `Aktivitas` untuk menyimpan data seperti ID, kategori, deskripsi, dampak karbon, dan frekuensi.
  - Array statis `[100]Aktivitas` untuk menyimpan data aktivitas.
- **Subprogram**:
  - Fungsi untuk menambah, mengedit, menghapus, mencari, dan mengurutkan aktivitas.
  - Prosedur untuk menghasilkan laporan bulanan dan rekomendasi.

## Contoh Penggunaan

1. **Menambah Aktivitas**:
   - Pilih opsi "Tambah Aktivitas" dari menu.
   - Masukkan detail seperti kategori (contoh: "Transportasi"), deskripsi (contoh: "Naik mobil"), dampak karbon (kg CO2), dan frekuensi.
2. **Melihat Laporan**:
   - Pilih opsi "Laporan Bulanan" untuk melihat total jejak karbon dan rekomendasi.
3. **Mencari Aktivitas**:
   - Gunakan opsi pencarian untuk menemukan aktivitas berdasarkan kategori.

## Catatan

Aplikasi ini dibuat untuk keperluan tugas besar dan fokus pada penerapan konsep algoritma dasar. Untuk pengembangan lebih lanjut, fitur seperti antarmuka grafis (GUI) atau penyimpanan data ke file dapat ditambahkan.