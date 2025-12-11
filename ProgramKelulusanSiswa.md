Program SistemPenilaianSiswa

Kamus
    jumlah_siswa  : Integer
    nama_siswa    : String
    kode_absen    : String   
    absen         : Boolean
    nilai_akhir   : Real

Algoritma
    Output("Masukkan jumlah siswa: ")
    Input(jumlah_siswa)

    For i = 1 to jumlah_siswa Do
        Output("--- Data Siswa Ke-" + i + " ---")
        
        //Input Nama
        Output("Masukkan Nama Siswa: ")
        Input(nama_siswa)

        //Input Absensi
        Output("Apakah absensi > 75%? (y/t): ")
        Input(kode_absen)

        //Cek Absensi
        If (kode_absen == "y") OR (kode_absen == "Y") Then
            absen <- true
        Else
            absen <- false
        EndIf

        //Logika Utama
        If (absen == false) Then
            //Jika Absen < 75%
            Output("HASIL: " + nama_siswa + " TIDAK LULUS (Absensi Kurang)")
        Else
            //Jika Absen >= 75%, meminta inputan nilai
            Output("Masukkan Nilai Akhir: ")
            Input(nilai_akhir)

            If (nilai_akhir >= 65) AND (nilai_akhir <= 100) Then
                Output("HASIL: " + nama_siswa + " dinyatakan LULUS")
            Else if (nilai_akhir >= 0) AND (nilai_akhir <65) Then
                Output("HASIL: " + nama_siswa + " TIDAK LULUS (Nilai Kurang)")
            Else
                output("Input
                
                
                
                
                
                
                Anda Invalid")
            EndIf
        EndIf

        Output("--------------------------------")
    EndFor

EndProgram