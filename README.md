# 🎓 Sistem Penilaian Kelulusan Siswa

> Program sederhana berbasis CLI (Command Line Interface) untuk menentukan kelulusan siswa.

---

## Tentang Program

Program ini dirancang untuk membantu guru atau dosen dalam memproses nilai siswa secara massal. Sistem menggunakan logika **bersarang (nested logic)** untuk efisiensi:

1.  **Cek Absensi Dulu:** Jika absensi di bawah 75%, siswa otomatis gugur (tidak perlu input nilai).
2.  **Cek Nilai:** Jika absensi aman, baru nilai dicek (Minimal 65).

---

##  Cara Menggunakan

Pastikan Anda sudah menginstall [Go](https://go.dev/dl/).

1.  **Clone atau Download** file program ini.
2.  Buka terminal/CMD di folder tersebut.
3.  Jalankan perintah:

    ```bash
    go run ProgramKelulusan.go
    ```

4.  Ikuti instruksi di layar:
    * Masukkan jumlah siswa.
    * Input nama, status kehadiran, dan nilai sesuai permintaan.

---

## 📸 Contoh Tampilan

```text
Masukkan jumlah siswa: 2

--- Data Siswa Ke-1 ---
Nama Siswa: Budi
Absensi > 75%? (y/t): t
Maaf, Budi TIDAK LULUS karena absensi kurang.

--- Data Siswa Ke-2 ---
Nama Siswa: Ani
Absensi > 75%? (y/t): y
Masukkan Nilai Akhir: 85
Selamat, Ani dinyatakan LULUS.
