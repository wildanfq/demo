# Deploy Backend Golang + PostgreSQL

Project ini mendemonstrasikan implementasi alur kerja **CI/CD** yang efisien dan modern untuk aplikasi backend berbasis **Golang** dan **PostgreSQL**. Tujuan utama dari sistem ini adalah membangun alur deployment yang sepenuhnya otomatis, konsisten, dan mudah dikelola sesuai dengan standar pengembangan perangkat lunak tahun 2026.

Aplikasi ini menggunakan **Golang native** untuk performa REST API yang optimal dan **PostgreSQL** sebagai basis data relasional. Untuk menjamin keseragaman lingkungan antara pengembangan dan produksi, seluruh komponen layanan dikemas menggunakan **Docker**, sehingga isu ketidakcocokan dependensi di server dapat dihindari sepenuhnya.

Sistem otomatisasi ini terintegrasi langsung dengan **GitHub Actions**. Setiap perubahan kode yang di-*push* ke *branch* utama akan memicu proses *build* Docker Image secara otomatis. Image yang telah berhasil disusun kemudian dikirim (*push*) ke **GitHub Container Registry (GHCR)** sebagai penyimpanan pusat yang aman dan terdistribusi.

Proses deployment pada sisi VPS kini jauh lebih efisien dengan dukungan **GitHub Self-Hosted Runner**. Setelah proses *build* di GitHub selesai, *runner* akan secara otomatis menerima instruksi untuk memperbarui container di server secara instan. Dengan arsitektur ini, VPS tidak lagi dibebani proses kompilasi kode, melainkan hanya menjalankan layanan terbaru yang telah dipersiapkan.

Arsitektur ini memberikan efisiensi sumber daya yang tinggi, isolasi lingkungan yang aman, serta pembaruan aplikasi yang cepat tanpa intervensi manual. Implementasi ini menjadi fondasi yang sangat stabil, skalabel, dan *production-ready* untuk pengembangan backend masa kini.
