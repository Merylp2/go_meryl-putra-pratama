> PENGERTIAN DOCKER

Docker adalah platform open source yang memungkinkan pengembang dan sysadmin untuk membangun, mengemas, dan menjalankan aplikasi di dalam container. Container memungkinkan suatu aplikasi untuk diisolasi dengan lingkungan yang terpisah dan independen, sehingga dapat berjalan di berbagai lingkungan tanpa perlu menginstall dependensi yang berbeda-beda

> CONTAINER

- Pengurangan pada lapisan aplikasi
- Pengurangan pada lapisan aplikasi Kontainer menggunakan lebih sedikit ruang daripada VM (gambar kontainer biasanya berukuran puluhan MB)
- Menangani lebih banyak aplikasi dan membutuhkan lebih sedikit VM dan sistem Operationg

> VIRTUAL MACHINES

- Abstraksi hadrware fisik
- Setiap VM menyertakan salinan lengkap sistem operasi lambat untuk booting

> DOCKER BASIC

- Image
- Container
- Engine
- Registry
- Control Plane

> PERINTAH - PERINTAH DOCKER

- FROM = Mendapatkan citra dari registri docker
- RUN = Menjalankan perintah bash saat membangun kontainer
- ENV = Mengatur variabel di dalam kontainer
- ADD = Menyalin berkas dengan proses lain
- COPY = Menyalin berkas
- WORKDIR = Mengatur direktori berkas yang berfungsi
- ENTRYPOINT = Menjalankan perintah setelah selesai membangun kontainer
- CMD = Menjalankan perintah tetapi dapat ditimpa