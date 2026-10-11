### Studi Kasus: Petugas Parkir Kampus
Seorang petugas parkir bertugas mencatat kendaraan yang masuk ke area parkir kampus setiap pagi. Petugas mulai mencatat dari kendaraan pertama, dan terus mencatat satu per satu setiap kendaraan yang datang, hingga akhirnya area parkir penuh (mencapai kapasitas maksimum 50 kendaraan). Begitu kendaraan ke-50 tercatat, petugas berhenti mencatat dan memasang papan "Parkir Penuh".


### Penjelasan
Proses perulangan yang terjadi pada studi kasus di atas adalah, `Petugas parkir` terus mencatat setiap motor yang masuk ke `area parkir` **berulang** kali selama `area parkir` masih belum penuh, dengan kapasitas maksimum 50 kendaraan. Artinya, selama `area parkir` belum terisi oleh 50 kendaraan (kapasitas maksimum area parkir), setiap kendaraan yang masuk ke dalam `area parkir` akan selalu dicatat oleh `Petugas parkir`. `Area parkir` tersebut awalnya tidak terisi oleh kendaraan apapun (0 kendaraan), setelah terdapat 1 kendaraan masuk, `petugas parkir` mencatat dan memperbarui jumlah kendaraan yang telah parkir di `area parkir`, proses pencatatan tersebut akan terus diulang oleh `petugas parkir` hingga `area parkir` mencapai batas maksimum nya/penuh yaitu ketika mencapai 50 kendaraan yang telah parkir di `area parkir`.

- Nilai awal: `Area parkir` = 0 kendaraan
- Proses yang berulang: `Petugas parkir mencatat kendaraan yang masuk ke area parkir`
- Kondisi berhenti: Ketika `area parkir` telah penuh (50 kendaraan telah masuk ke `area parkir`)