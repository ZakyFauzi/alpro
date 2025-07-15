# Luxury Watches - Website Penjualan Jam Tangan

Website penjualan jam tangan premium yang modern, responsif, dan user-friendly. Dibuat dengan HTML, CSS, dan JavaScript vanilla.

## 🚀 Fitur Utama

### ✨ Desain & UI/UX
- **Desain Modern**: Interface yang elegan dengan gradien dan animasi smooth
- **Responsif**: Optimal di desktop, tablet, dan mobile
- **Animasi**: Transisi dan hover effects yang menarik
- **Typography**: Font Poppins untuk keterbacaan yang baik

### 🛍️ E-commerce Features
- **Katalog Produk**: 8 produk jam tangan dengan kategori berbeda
- **Filter Produk**: Filter berdasarkan kategori (Luxury, Sport, Business, Casual)
- **Pencarian**: Search real-time berdasarkan nama, kategori, dan deskripsi
- **Keranjang Belanja**: 
  - Tambah/hapus produk
  - Update quantity
  - Kalkulasi total otomatis
  - Modal cart yang slide dari kanan
- **Wishlist**: Fitur favorit produk (placeholder)

### 📱 Responsive Design
- **Mobile-First**: Optimized untuk mobile devices
- **Hamburger Menu**: Menu navigasi yang collapsible di mobile
- **Touch-Friendly**: Button dan interactive elements yang mudah di-tap
- **Flexible Grid**: Layout yang adaptif di berbagai ukuran layar

### 🎯 User Experience
- **Smooth Scrolling**: Navigasi antar section yang halus
- **Loading Animations**: Animasi fade-in saat scroll
- **Notifications**: Toast notifications untuk feedback user
- **Keyboard Shortcuts**: 
  - `ESC` untuk tutup cart
  - `Ctrl/Cmd + K` untuk focus search
- **Form Validation**: Validasi form kontak

## 📁 Struktur File

```
luxury-watches/
├── index.html          # File HTML utama
├── styles.css          # File CSS styling
├── script.js           # File JavaScript functionality
└── README.md           # Dokumentasi
```

## 🛠️ Teknologi yang Digunakan

- **HTML5**: Semantic markup
- **CSS3**: 
  - Flexbox & Grid Layout
  - CSS Variables
  - Animations & Transitions
  - Media Queries
- **JavaScript (ES6+)**:
  - DOM Manipulation
  - Event Handling
  - Local Storage (untuk cart)
  - Intersection Observer API
  - Fetch API (untuk gambar)

## 🎨 Komponen Website

### 1. Header & Navigation
- Logo brand dengan icon jam
- Menu navigasi responsif
- Search box dengan real-time search
- Cart icon dengan counter
- Hamburger menu untuk mobile

### 2. Hero Section
- Background gradient yang menarik
- Call-to-action buttons
- Hero image dengan shadow effect

### 3. Categories Section
- 4 kategori utama dengan icons
- Hover animations
- Grid layout responsif

### 4. Products Section
- Product cards dengan hover effects
- Filter buttons
- Product badges (Best Seller, Limited, dll)
- Add to cart & wishlist buttons

### 5. About Section
- Informasi tentang brand
- Feature highlights dengan icons
- Split layout dengan image

### 6. Contact Section
- Contact information dengan icons
- Contact form dengan validation
- Responsive grid layout

### 7. Footer
- Social media links
- Quick links
- Copyright information

### 8. Shopping Cart Modal
- Slide-in animation dari kanan
- Product list dengan images
- Quantity controls
- Total calculation
- Checkout functionality

## 🚀 Cara Menjalankan

1. **Clone atau download** semua file ke direktori lokal
2. **Buka file `index.html`** di browser modern
3. **Atau gunakan live server**:
   ```bash
   # Jika menggunakan VS Code
   # Install extension "Live Server"
   # Right click index.html -> "Open with Live Server"
   
   # Atau menggunakan Python
   python -m http.server 8000
   
   # Atau menggunakan Node.js
   npx serve .
   ```

## 📱 Browser Support

- ✅ Chrome (versi terbaru)
- ✅ Firefox (versi terbaru)
- ✅ Safari (versi terbaru)
- ✅ Edge (versi terbaru)
- ⚠️ IE11 (tidak didukung)

## 🎯 Fitur Interaktif

### Keranjang Belanja
- Klik icon cart untuk membuka modal
- Tambah produk dengan tombol "Tambah ke Keranjang"
- Update quantity dengan tombol +/- 
- Hapus produk dengan tombol "Hapus"
- Total otomatis terupdate
- Checkout dengan alert konfirmasi

### Filter & Search
- Filter berdasarkan kategori
- Search real-time
- Reset filter dengan tombol "Semua"

### Mobile Navigation
- Hamburger menu untuk mobile
- Smooth transitions
- Touch-friendly interactions

## 🎨 Customization

### Mengubah Warna
Edit variabel CSS di `styles.css`:
```css
:root {
  --primary-color: #e74c3c;
  --secondary-color: #2c3e50;
  --accent-color: #27ae60;
}
```

### Menambah Produk
Edit array `products` di `script.js`:
```javascript
const products = [
  {
    id: 9,
    name: "Nama Jam Tangan",
    category: "luxury", // luxury, sport, business, casual
    price: 50000000,
    image: "url-gambar",
    badge: "New", // optional
    description: "Deskripsi produk"
  }
];
```

### Mengubah Konten
Edit teks di `index.html` sesuai kebutuhan brand Anda.

## 🔧 Performance Optimizations

- **Lazy Loading**: Images load saat visible
- **CSS Animations**: Hardware-accelerated animations
- **Minimal JavaScript**: Vanilla JS tanpa framework
- **Optimized Images**: Compressed images dari Unsplash
- **Efficient DOM**: Minimal DOM manipulation

## 📞 Support

Untuk pertanyaan atau feedback, silakan hubungi:
- Email: info@luxurywatches.com
- Phone: +62 21 1234 5678

## 📄 License

Proyek ini dibuat untuk tujuan edukasi dan demonstrasi. Silakan modifikasi sesuai kebutuhan Anda.

---

**Dibuat dengan ❤️ untuk toko jam tangan premium**
