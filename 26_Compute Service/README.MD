> SYSTEM SOFTWARE DEPLOYMENT

Deployment adalah istilah yang mencakup semua proses yang terlibat dalam mendapatkan software atau hardware baru dan berjalan dengan baik di lingkungannya, termasuk instalasi, konfigurasi, pengoperasian, pengujian dan membuat perubahan yang diperlukan

> STRATEGY DEPLOYMENT

- Big-Bang Deployment Strategy
- Rollout Deployment Strategy
- Blue/Green Deployment Strategy
- Canary Deployment Strategy

> KELEBIHAN DAN KEKURANGAN

1. Big-Bang Deployment Strategy
Kelebihan:
- Mudah diimplementasikan
- Perubahan kepada sistem langsung 100% secara instan

Kekurangan:
- Terlalu beresiko, rata-rata downtime cukup lama

2. Rollout Deployment Strategy
Kelebihan:
- Lebih aman dan less downtime

Kekurangan:
- Akan ada 2 aplikasi berjalan secara bersamaan sampai semua server terdeploy

3. Blue/Green Deployment Strategy
Kelebihan:
- Perubahan sangat cepat
- Tidak ada issue beda versi pada service seperti yang pada rollout deployment

Kekurangan:
- Resource yang dibutuhkan lebih banyak
- Testing harus benar benar sangat diprioritaskan sebelum diswitch

4. Canary Deployment Strategy
Kelebihan:
- Cukup aman
- Mudah untuk rollback jika terjadi error/bug

Kekurangan:
- Untuk mencapai 100% cukup lama dibanding Blue/green deployment

> SIMPLE DEVPLOYMENT PROCESS

- Development env
- Test env
- UAT ent
- Production env

