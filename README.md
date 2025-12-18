# Sistem Penilaian Kelulusan Siswa

> Program sederhana berbasis CLI (Command Line Interface) untuk menentukan kelulusan siswa.

---

## Tentang Program

Program ini dirancang untuk membantu guru atau dosen dalam memproses nilai siswa secara massal.

1.  **Cek Absensi Dulu:** Jika absensi di bawah 75%, siswa otomatis gugur (tidak perlu input nilai).
2.  **Cek Nilai:** Jika absensi aman, baru nilai dicek (Minimal 65).

## Contoh Tampilan

```text
Masukkan jumlah siswa: 2

--- Data Siswa Ke-1 ---
Nama Siswa: Budi
Absensi > 75%? (T/F): F
Maaf, Budi TIDAK LULUS karena absensi kurang.

--- Data Siswa Ke-2 ---
Nama Siswa: Ani
Absensi > 75%? (T/F): T
Masukkan Nilai Akhir: 85
Selamat, Ani dinyatakan LULUS.
```

Berikut adalah logika jalannya program ini:
<p align="center">
  <img src="./FlowChart.svg" alt="Flowchart Sistem" width="90%">
</p>


