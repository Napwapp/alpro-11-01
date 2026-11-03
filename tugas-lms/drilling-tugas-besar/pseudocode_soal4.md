```text
program penjualan_tiket_wisata
kamus
   kuota, jumlah, harga, total : integer
   terjual, pendapatan, transaksi, kategori : integer
   hari, jenis_hari, nama_kategori : string
algoritma
   input(kuota)
   terjual <- 0
   pendapatan <- 0
   transaksi <- 0
   for kuota > 0 do
      output("Sisa kuota: ", kuota)

      for do
         output("Kategori tiket untuk (1 = Anak, 2 = Dewasa): ")
         input(kategori)
         if kategori == 1 or kategori == 2 then
            break
         endif
         output("kategori tidak valid (Masukkan angka 1 atau 2)")
      endfor

      for do
         output("Masukkan hari (Senin - Minggu): ")
         input(hari)
         if hari == "Senin" or hari == "Selasa" or hari == "Rabu" or hari == "Kamis" or hari == "Jumat" or hari == "Sabtu" or hari == "Minggu" then
            break
         endif
         output("Hari tidak valid (Senin, Selasa, Rabu, Kamis, Jumat, Sabtu, Minggu)")
      endfor

      for do
         output("Masukkan jumlah tiket: ")
         input(jumlah)
         if jumlah <= 0 then
            output("jumlah tiket harus lebih dari 0")
            continue
         endif
         if jumlah > kuota then
            output("kuota tidak mencukupi")
            continue
         endif
         break
      endfor

      switch hari
         case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
            jenis_hari <- "hari biasa"
         case "Sabtu", "Minggu":
            jenis_hari <- "akhir pekan"
      endswitch

      switch kategori
         case 1:
            nama_kategori <- "anak"
            if jenis_hari == "hari biasa" then
               harga <- 20000
            else
               harga <- 30000
            endif
         case 2:
            nama_kategori <- "dewasa"
            if jenis_hari == "hari biasa" then
               harga <- 35000
            else
               harga <- 50000
            endif
      endswitch

      total <- harga * jumlah
      kuota <- kuota - jumlah
      terjual <- terjual + jumlah
      pendapatan <- pendapatan + total
      transaksi <- transaksi + 1

      output("Bukti pembelian ke-", transaksi)
      output("Kategori: ", nama_kategori)
      output("Hari: ", hari)
      output("Jumlah tiket: ", jumlah)
      output("Harga satuan: ", harga)
      output("Total bayar: ", total)
      output("Sisa kuota: ", kuota)
   endfor

   output("Tiket habis terjual")
   output("Sisa kuota tiket: ", kuota)
   output("Jumlah tiket terjual: ", terjual)
   output("Total pendapatan: ", pendapatan)
endprogram
```