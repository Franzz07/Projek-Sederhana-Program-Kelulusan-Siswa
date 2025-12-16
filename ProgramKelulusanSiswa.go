package main

import "fmt"

func main() {
	var jumlah_siswa int
	var i int
	var nama_siswa string
	var kode_absen string
	var absen bool
	var nilai_akhir float64

	// inputan jumlah siswa
	fmt.Print("Masukan Jumlah Siswa: ")
	fmt.Scan(&jumlah_siswa)

	//
	for i = 1; i <= jumlah_siswa; i++ {
		fmt.Println("\n--- Data Siswa Ke-", i, "---")

		// input nama
		fmt.Print("Masukan Nama Siswa (satu kata): ")
		fmt.Scan(&nama_siswa)

		// input absen
		for {
			fmt.Print("Apakah Absen >= 75% ? (T/F): ")
			fmt.Scan(&kode_absen)

			if kode_absen == "T" || kode_absen == "t" {
				absen = true
				break // Keluar dari loop karena input valid (T/t)

			} else if kode_absen == "F" || kode_absen == "f" {
				absen = false
				break // Keluar dari loop karena input valid (F/f)

			} else {
				// Input tidak valid
				fmt.Println("Inputan Absen Invalid. Harap Masukkan 'T' atau 'F'.")
			}
		}

		// jika absen kurang dari 75%
		if absen == false {
			fmt.Printf("Maaf, %s dinyatakan TIDAK LULUS karena absensi kurang.\n", nama_siswa)

			// jika absen lebih dari 75%
		} else {
			fmt.Printf("Masukan Nilai Akhir: ")
			fmt.Scan(&nilai_akhir)

			// cek nilai
			if nilai_akhir >= 65 && nilai_akhir <= 100 {
				fmt.Printf("Selamat, %s dinyatakan LULUS,\n", nama_siswa)

			} else if nilai_akhir >= 0 && nilai_akhir < 65 {
				fmt.Printf("Maaf, %s dinyatakan TIDAK LULUS karena nilai akhir kurang.\n", nama_siswa)

			} else {
				fmt.Println("Inputan Nilai Anda Invalid")
			}
		}

	}
}
