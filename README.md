# Go & React Management App

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org)
[![React Version](https://img.shields.io/badge/react-18.2.0-blue.svg)](https://reactjs.org)

Sebuah aplikasi full-stack yang menampilkan RESTful API yang dibangun dengan **Golang** sebagai backend dan antarmuka pengguna yang interaktif dibangun dengan **React.js** sebagai frontend.

---

## ✨ Fitur Utama

- **Backend Cepat & Efisien**: Dibangun dengan Go menggunakan framework [Fiber](https://gofiber.io/) untuk performa tinggi.
- **Frontend Modern**: Antarmuka yang responsif dan dinamis menggunakan [React.js](https://reactjs.org/).
- **RESTful API**: Desain API yang bersih dan terstruktur untuk operasi CRUD (Create, Read, Update, Delete) pada data.
- **Database**: Menggunakan [PostgreSQL](https://www.postgresql.org/) dengan [GORM](https://gorm.io/) sebagai ORM untuk interaksi database yang mudah.
- **Konfigurasi Terpusat**: Pengelolaan konfigurasi melalui file `.env`.

---

## 📚 Tumpukan Teknologi (Tech Stack)

| Bagian       | Teknologi    | Deskripsi                                   |
| :----------- | :----------- | :------------------------------------------ |
| **Backend**  | Golang       | Bahasa pemrograman utama untuk API.         |
|              | Gin Gonic    | Framework web untuk routing dan middleware. |
|              | GORM         | ORM untuk interaksi dengan database.        |
|              | PostgreSQL   | Sistem database relasional.                 |
| **Frontend** | React.js     | Library JavaScript untuk membangun UI.      |
|              | Axios        | HTTP client untuk berkomunikasi dengan API. |
|              | Tailwind CSS | Framework CSS untuk styling cepat.          |

---

## 📂 Struktur Direktori

```bash

/projeqtor-api-go/
├── config/ # Konfigurasi aplikasi (database, env)
├── controllers/ # Menangani request & response HTTP
├── database/ # Pengaturan koneksi, migrasi & seeding data
│ ├── migrations/
│ └── seed/
├── docs/ # Dokumentasi API (misalnya Swagger)
├── middleware/ # Middleware untuk request (misal: auth, logger)
├── models/ # Struct/entitas yang merepresentasikan data
├── repositories/ # Logika akses data (interaksi dengan database)
├── routes/ # Definisi semua rute/endpoint API
├── services/ # Logika bisnis utama aplikasi
├── utils/ # Fungsi-fungsi bantuan (helpers)
├── .env # File konfigurasi (tidak di-commit)
├── .gitignore # File dan folder yang diabaikan oleh Git
├── env.example # Contoh file konfigurasi
├── go.mod # Definisi modul dan dependensi Go
├── go.sum # Checksum untuk integritas dependensi
├── main.go # Titik masuk utama aplikasi
└── README.md # File ini
```

---

## 🛠️ Prasyarat

Pastikan perangkat Anda sudah terinstal perangkat lunak berikut:

- [Go](https://golang.org/dl/) versi 1.21+
- [Node.js](https://nodejs.org/) versi 18.0+
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)

---

## ⚙️ Instalasi & Konfigurasi

Ikuti langkah-langkah ini untuk menjalankan proyek secara lokal.

### 1. Clone Repository

```bash
git clone [https://github.com/raddva/projeqtor-api-go.git](https://github.com/raddva/projeqtor-api-go.git)
cd projeqtor-api-go
```

### 2. Konfigurasi Backend (Server)

1. Masuk ke direktori server:

```bash
cd server
```

2. Copy env.example dan isi dengan konfigurasi database Anda

3. Instal semua dependensi Go:

```bash
go mod tidy
```

## 🚀 Menjalankan Aplikasi

Menjalankan Server Backend

```bash
go run main.go
```

Server API akan berjalan di http://localhost:3030.
