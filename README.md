# CI/CD Pipeline for Golang & PostgreSQL Backend

Project ini menerapkan sistem **CI/CD modern** untuk backend berbasis **Golang** dan **PostgreSQL** dengan tujuan menciptakan alur deployment yang sepenuhnya otomatis, konsisten, dan siap produksi. Arsitektur ini dirancang untuk mengurangi intervensi manual, meningkatkan stabilitas deployment, serta memastikan setiap perubahan kode dapat langsung masuk ke lingkungan server secara aman dan terkontrol sesuai standar pengembangan sistem modern.

Backend dibangun menggunakan **Golang native** untuk memastikan performa API tetap ringan, cepat, dan efisien dalam menangani request. Sementara itu, **PostgreSQL** digunakan sebagai database utama untuk menjamin integritas data dan dukungan transaksi yang kuat. Seluruh komponen aplikasi dikemas menggunakan **Docker**, sehingga lingkungan development dan production selalu konsisten dan tidak bergantung pada konfigurasi sistem operasi server.

Proses CI/CD diintegrasikan dengan **GitHub Actions**, di mana setiap perubahan yang di-*push* ke branch utama akan secara otomatis memicu proses build Docker Image. Image yang dihasilkan kemudian dipublikasikan ke **GitHub Container Registry (GHCR)** sebagai pusat distribusi image yang terstandarisasi. Dengan mekanisme ini, proses build tidak lagi dilakukan di server produksi, sehingga beban komputasi sepenuhnya dipindahkan ke pipeline CI.

Tahap deployment dijalankan menggunakan **GitHub Self-Hosted Runner** yang terhubung langsung ke VPS. Runner ini bertugas mengeksekusi instruksi deployment setelah proses build selesai, termasuk menarik image terbaru dan melakukan update container secara otomatis. Pendekatan ini membuat VPS hanya berfungsi sebagai runtime environment, bukan sebagai tempat kompilasi, sehingga lebih ringan, stabil, dan aman.

Secara keseluruhan, arsitektur ini menghasilkan sistem deployment yang cepat, terotomatisasi, dan mudah di-scale. Kombinasi Golang, PostgreSQL, Docker, GitHub Actions, dan Self-Hosted Runner menjadikan pipeline ini siap digunakan untuk kebutuhan production-grade system dengan tingkat efisiensi dan reliability yang tinggi.

# Storage, Backup & Disaster Recovery Architecture

## 1. Arsitektur Infrastruktur Multi-Server

Sistem ini menggunakan arsitektur dua server untuk memisahkan environment produksi dan backup secara ketat. VPS utama (Google Cloud) berfungsi sebagai production server yang menjalankan seluruh komponen aplikasi, termasuk backend Golang, database PostgreSQL, serta seluruh stack Docker Compose. Semua proses bisnis, request pengguna, dan pemrosesan data berjalan di server ini untuk menjaga performa, latency rendah, dan stabilitas layanan.

Server kedua adalah VPS Backup yang tidak menjalankan aplikasi produksi sama sekali, melainkan difokuskan sebagai storage isolasi untuk seluruh data cadangan. Server ini menerima replikasi dari server utama dan menjadi lapisan terakhir perlindungan data jika terjadi kegagalan total pada production server. Dengan pemisahan ini, sistem tetap memiliki mekanisme pemulihan meskipun server utama mengalami crash, corruption, atau kehilangan akses.

## 2. Sistem Penyimpanan File (Storage Layer)

Sistem penyimpanan file menggunakan **SeaweedFS** sebagai object storage berbasis S3-compatible API, menggantikan penyimpanan lokal tradisional. Semua file seperti gambar, dokumen, dan media yang diunggah pengguna dikirim langsung dari backend Golang ke SeaweedFS melalui API, sehingga proses penyimpanan tidak membebani database PostgreSQL.

SeaweedFS dijalankan di dalam Docker dan menggunakan Docker Volume yang terisolasi untuk menyimpan data secara persisten. Arsitektur ini membuat sistem lebih fleksibel, mudah dipindahkan, serta mendukung skala besar untuk menangani jutaan file kecil secara efisien. Selain itu, SeaweedFS juga mendukung replikasi dan ekspor data yang memudahkan proses backup ke server lain.

## 3. Backup Database PostgreSQL

Backup database PostgreSQL dikelola oleh **Databasus** yang berjalan sebagai container khusus di dalam ekosistem Docker. Tool ini secara otomatis menjalankan full backup terjadwal serta mengelola WAL (Write-Ahead Logging) untuk mendukung Point-in-Time Recovery (PITR), sehingga database dapat dipulihkan ke kondisi waktu yang sangat spesifik.

Seluruh hasil backup dikompresi dan dikirim secara otomatis ke VPS Backup atau storage eksternal tanpa mengganggu performa server utama. Selain itu, Databasus dilengkapi sistem monitoring yang memberikan notifikasi real-time melalui Discord atau Telegram ketika proses backup berhasil maupun gagal, sehingga kondisi database selalu terpantau secara aktif.

## 4. Backup Sistem dan Konfigurasi (System Backup)

Backup sistem mencakup seluruh komponen penting di luar database, seperti file `.env`, konfigurasi `docker-compose.yml`, source cadangan, serta volume SeaweedFS. Proses ini ditangani oleh **Restic** yang berjalan secara otomatis melalui cron job. Restic menggunakan enkripsi end-to-end dan deduplikasi data, sehingga hanya perubahan yang benar-benar terjadi yang akan diproses, membuat backup lebih efisien dan aman.

Setelah proses backup selesai, **Rclone** berfungsi sebagai layer transport untuk mengirim hasil backup Restic ke VPS Backup. Kombinasi Restic dan Rclone memastikan seluruh konfigurasi sistem, environment Docker, dan data storage memiliki salinan lengkap di lokasi terpisah, sehingga sistem dapat direkonstruksi kapan saja dengan konsisten.

## 5. Prosedur Disaster Recovery

Ketika terjadi kegagalan total pada VPS utama, langkah pertama adalah menyiapkan VPS baru dan menginstal Docker serta Docker Compose sebagai fondasi environment. Setelah itu, seluruh data backup dari VPS Backup ditarik kembali menggunakan Rclone untuk memulihkan konfigurasi sistem, termasuk environment Docker dan volume SeaweedFS.

Tahap berikutnya adalah pemulihan database PostgreSQL menggunakan Databasus dengan mengembalikan full backup terakhir beserta WAL log untuk mencapai kondisi paling akhir sebelum kegagalan terjadi. Setelah database, storage, dan backend berhasil dipulihkan, sistem dijalankan kembali menggunakan Docker Compose. Ketika semua service sudah aktif dan saling terhubung, traffic kemudian dialihkan ke server baru, dan sistem kembali beroperasi normal tanpa kehilangan data yang signifikan.
