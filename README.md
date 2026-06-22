# CI/CD Pipeline Backend Golang & PostgreSQL

Arsitektur CI/CD ini dirancang untuk backend berbasis Golang dan PostgreSQL dengan tujuan menciptakan proses deployment yang otomatis, konsisten, dan siap production. Seluruh alur pipeline dibuat untuk menghilangkan proses manual, meningkatkan stabilitas rilis, serta memastikan setiap perubahan kode dapat langsung masuk ke server melalui mekanisme yang terkontrol dan aman.

Backend menggunakan Golang untuk memastikan performa API tetap cepat, ringan, dan efisien. PostgreSQL digunakan sebagai database utama karena stabilitas, integritas data, dan dukungan transaksi yang kuat. Seluruh aplikasi dikemas menggunakan Docker agar environment development dan production tetap konsisten dan tidak bergantung pada konfigurasi server.

Proses CI/CD dijalankan menggunakan GitHub Actions yang akan memicu build otomatis setiap kali ada perubahan pada branch utama. Hasil build berupa Docker Image kemudian dipublikasikan ke GitHub Container Registry (GHCR). Dengan pendekatan ini, server produksi tidak lagi melakukan proses build sehingga seluruh beban komputasi dipindahkan ke pipeline CI.

Proses deployment dijalankan melalui GitHub Self-Hosted Runner yang terpasang langsung di VPS. Runner ini bertugas menarik image terbaru dari registry dan melakukan update container secara otomatis. Dengan model ini, VPS hanya berfungsi sebagai runtime environment sehingga lebih ringan, stabil, dan aman.

Secara keseluruhan, kombinasi Golang, PostgreSQL, Docker, GitHub Actions, dan Self-Hosted Runner menghasilkan sistem deployment yang otomatis, cepat, dan mudah di-scale dengan reliability tinggi untuk production system.

---

# Storage, Backup & Disaster Recovery Architecture

Arsitektur ini menggunakan dua server yang dipisahkan antara production dan backup. VPS utama berfungsi sebagai production server yang menjalankan seluruh aplikasi, termasuk backend Golang, PostgreSQL, dan Docker Compose. Semua traffic dan proses utama berjalan di server ini untuk menjaga performa dan latency tetap optimal.

Server kedua berfungsi sebagai backup server yang hanya digunakan untuk menyimpan seluruh data cadangan. Semua backup dari production direplikasi ke server ini sebagai perlindungan terhadap kegagalan total seperti crash, corruption, atau kehilangan akses server.

## Storage Layer (SeaweedFS)

File storage menggunakan SeaweedFS sebagai object storage berbasis S3-compatible API. Semua file seperti gambar, dokumen, dan media dikirim langsung dari backend ke SeaweedFS sehingga database tidak terbebani oleh data file.

SeaweedFS berjalan dalam container Docker dengan persistent volume untuk menjaga data tetap aman. Sistem ini mendukung replikasi dan scaling, sehingga cocok untuk menangani file dalam jumlah besar secara efisien.

## Database Backup (PostgreSQL + Databasus)

Backup PostgreSQL dikelola menggunakan Databasus yang berjalan di dalam container Docker. Sistem ini melakukan full backup secara terjadwal serta menyimpan WAL (Write-Ahead Logging) untuk mendukung Point-in-Time Recovery (PITR), sehingga database bisa dipulihkan ke kondisi waktu tertentu secara presisi.

Backup dikirim otomatis ke server backup tanpa mengganggu performa production. Sistem ini juga dilengkapi monitoring dan notifikasi real-time melalui Discord atau Telegram untuk setiap status backup.

## System Backup (Restic + Rclone)

Backup sistem mencakup seluruh konfigurasi penting seperti file environment, docker-compose, konfigurasi aplikasi, source cadangan, dan volume SeaweedFS.

Restic digunakan untuk melakukan backup terenkripsi dengan deduplikasi data agar lebih efisien. Proses ini berjalan otomatis melalui cron job. Hasil backup kemudian dikirim ke server backup menggunakan Rclone sebagai transport layer.

Kombinasi Restic dan Rclone memastikan seluruh sistem dapat dipulihkan secara lengkap dan konsisten kapan saja.

## Disaster Recovery Procedure

Jika terjadi kegagalan total pada server utama, VPS baru akan disiapkan dengan Docker dan Docker Compose sebagai dasar environment. Setelah itu, semua data dari server backup dipulihkan menggunakan Rclone, termasuk konfigurasi sistem dan storage SeaweedFS.

Database kemudian dipulihkan menggunakan Databasus dengan full backup terakhir dan WAL logs untuk mengembalikan kondisi paling terbaru sebelum kegagalan. Setelah semua service aktif kembali, sistem dijalankan menggunakan Docker Compose dan traffic dialihkan ke server baru.

---

# Monitoring & Observability

Sistem monitoring dirancang untuk memberikan visibilitas penuh terhadap seluruh infrastruktur, mulai dari server, container, hingga aplikasi. Tujuannya adalah mendeteksi masalah lebih awal sebelum berdampak ke pengguna.

Prometheus digunakan untuk mengumpulkan metrics dari VPS, Docker, database, dan backend Golang. Data yang dikumpulkan mencakup CPU, RAM, disk usage, latency API, throughput, dan status service.

Grafana digunakan untuk visualisasi data dari Prometheus dalam bentuk dashboard yang mudah dibaca, sehingga kondisi sistem dapat dipantau secara real-time dan historis.

Loki digunakan sebagai sistem logging terpusat untuk mengumpulkan log dari seluruh service. Dengan ini, proses debugging menjadi lebih cepat karena semua log berada dalam satu sistem terstruktur.

Alerting ditangani oleh Alertmanager yang terhubung dengan Prometheus. Setiap kondisi kritis akan dikirimkan secara real-time ke Telegram dan Discord agar langsung diketahui tanpa perlu pengecekan manual.

Secara keseluruhan, kombinasi Prometheus, Grafana, Loki, dan Alertmanager membentuk sistem observability yang lengkap, terintegrasi, dan siap untuk production dengan kemampuan monitoring, analisis, dan respons yang cepat.
