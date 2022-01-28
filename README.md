# Rest API dengan menggunakan Golang dan Mysql

Full Rest API product dengan bahasa pemrograman Go dan RDMS Mysql

Semua Endpoint pada aplikasi ini diuji menggukan Aplikasi Postman (https://www.getpostman.com/).

## Berikut beberapa langkah untuk menjalankan project Go tersebut.

**1. Mengunduh project**

```bash
git clone https://github.com/kelvinha/go_product_restapi

```

**2. Import Tabel**

```bash
- Buatlah sebuah database pada PhpMyAdmin dengan nama database sesuai dengan yang ada pada folder config.
- Setelah itu import tabel .sql yang sudah di sediakan di folder database
- Import tabel selesai.

```

**3. Run Project**

```bash
Jalankan perintah go run main.go untuk menjalankan project dan instalisasi package / library yang dibutuhkan

```

# Endpoint untuk Get Data Produk

| Method | Route | Description | params |
| --- | --- | --- | --- |
| `GET` | `/produk` | Menampilkan seluruh data `produk` secara urut | None |
| `GET` | `/produk/?sortingBy=ASC` | Menampilkan seluruh data `produk` dengan urutan nama produk dari A hingga Z | None |
| `GET` | `/produk/?sortingBy=DESC` | Menampilkan seluruh data `produk` dengan urutan nama produk dari Z hingga A | None |
| `GET` | `/produk/?sortingBy=high-price` | Menampilkan seluruh data `produk` dengan urutan harga produk dari yang termahal hingga termurah | None |
| `GET` | `/produk/?sortingBy=low-price` | Menampilkan seluruh data `produk` dengan urutan harga produk dari yang termurah hingga termahal | None |

---
# Endpoint untuk Update Data Produk

| Method | Route | Description | params |
| --- | --- | --- | --- |
| `PUT` |`/update` | Melakukan update sebuah data `produk` berdasarkan id produk | id_produk |

---


<br>
