# Go_Calculator
# Kalkulator Sederhana di Go

## Deskripsi
Program ini adalah kalkulator sederhana yang dibuat dengan bahasa pemrograman Go. Kalkulator ini mendukung operasi matematika dasar seperti penjumlahan, pengurangan, perkalian, dan pembagian. Pengguna dapat memasukkan operasi secara berulang hingga mengetikkan '#' untuk keluar dari program.

## Fitur
- Operasi yang didukung:
  - Penjumlahan (`+`)
  - Pengurangan (`-`)
  - Perkalian (`*`)
  - Pembagian (`/`)
- Memeriksa kesalahan input, termasuk pembagian oleh nol
- Memungkinkan operasi berturut-turut hingga pengguna keluar

## Cara Menjalankan
1. Pastikan Anda telah menginstal Go di sistem Anda.
2. Simpan file program sebagai `kalkulator.go`.
3. Buka terminal atau command prompt, lalu navigasikan ke direktori tempat file disimpan.
4. Jalankan perintah berikut untuk menjalankan program:
   ```sh
   go run main.go
   ```

## Cara Menggunakan
1. Masukkan operasi dengan format:
   ```
   angka1 operator angka2
   ```
   Contoh:
   ```
   5 + 3
   ```
   Output:
   ```
   Hasil: 8
   ```
2. Untuk melanjutkan, cukup masukkan operator dan angka berikutnya:
   ```
   * 2
   ```
   Output:
   ```
   Hasil: 16
   ```
3. Untuk keluar dari program, ketik `#`.

## Catatan
- Pastikan input sesuai dengan format yang benar agar program tidak error.
- Operator harus berada di antara dua angka, misalnya `5 + 3`, bukan `5+3`.
- Pembagian dengan nol akan menghasilkan pesan kesalahan "Input tidak valid".


