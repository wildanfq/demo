
# Deploy Backend Golang + PostgreSQL

Project ini mendemonstrasikan implementasi alur kerja **CI/CD** yang efisien untuk aplikasi backend berbasis **Golang** dan **PostgreSQL**. Tujuan utamanya adalah menciptakan sistem deployment yang otomatis, konsisten, dan mudah dikelola sesuai standar industri modern tahun 2026.

Sistem ini dibangun menggunakan **Golang native** untuk REST API dan **PostgreSQL** sebagai basis datanya. Agar lingkungan pengembangan dan produksi selalu seragam serta terhindar dari isu kompatibilitas, seluruh layanan dikemas secara utuh menggunakan **Docker**.

Alur otomatisasi berjalan ketika kode di-push ke GitHub. **GitHub Actions** akan secara otomatis mem-build aplikasi, menyusun Docker Image, lalu mengunggahnya ke **GitHub Container Registry (GHCR)**. Tahapan ini memastikan selalu ada versi aplikasi terbaru yang siap dirilis ke server tanpa proses manual.

Pada sisi server, deployment dikendalikan menggunakan **Docker Compose**. VPS tidak dibebani tugas komputasi untuk mem-build kode, melainkan hanya menarik (*pull*) image terbaru dari GHCR melalui **Watchtower**. Sistem ini memungkinkan aplikasi di server untuk memperbarui dirinya sendiri secara otomatis dan stabil setiap kali ada perubahan kode terbaru.

Arsitektur ini dirancang untuk memberikan efisiensi tinggi, isolasi lingkungan yang aman, dan proses pembaruan yang otomatis, menjadikannya fondasi yang realistis untuk aplikasi backend *production-ready*.
