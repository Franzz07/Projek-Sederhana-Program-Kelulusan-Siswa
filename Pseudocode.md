For i = 1 to jumlah_siswa Do
    Output("--- Data Siswa Ke-" + i + " ---")
    
    //Input Nama
    Output("Masukkan Nama Siswa: ")
    Input(nama_siswa)

    //Input Absensi
 For 
    Output("Apakah absensi > 75%? (T/F): ")
    Input(kode_absen)

    //Cek Absensi
    If kode_absen == "T" or kode_absen == "t" Then
        absen <- true
        break
    Else if kode_absen == "F" || kode_absen == "f" Then
        absen <- false
        break
    Else Output("Inputan Absen Invalid. Harap Masukkan 'T' atau 'F'.")
    EndIf
 Endfor

    //Logika Utama
    If (absen == false) Then
        //Jika Absen < 75%
        Output("HASIL: " + nama_siswa + " TIDAK LULUS (Absensi Kurang)")
    Else
        //Jika Absen >= 75%, meminta inputan nilai
    For
        Output("Masukkan Nilai Akhir: ")
        Input(nilai_akhir)
    
        If (nilai_akhir >= 65) AND (nilai_akhir <= 100) Then
            Output("HASIL: " + nama_siswa + " dinyatakan LULUS")
            break
        Else if (nilai_akhir >= 0) AND (nilai_akhir <65) Then
            Output("HASIL: " + nama_siswa + " TIDAK LULUS (Nilai Kurang)")
            break
        Else
            output("Input Anda Invalid")
        EndIf
        EndFor
    EndIf

    Output("--------------------------------")
EndFor