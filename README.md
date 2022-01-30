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
| `GET` | `/produk` | Menampilkan seluruh data `produk` secara urut | - |
| `GET` | `/produk/?sortingBy=ASC` | Menampilkan seluruh data `produk` dengan urutan nama produk dari A hingga Z | - |
| `GET` | `/produk/?sortingBy=DESC` | Menampilkan seluruh data `produk` dengan urutan nama produk dari Z hingga A | - |
| `GET` | `/produk/?sortingBy=high-price` | Menampilkan seluruh data `produk` dengan urutan harga produk dari yang termahal hingga termurah | - |
| `GET` | `/produk/?sortingBy=low-price` | Menampilkan seluruh data `produk` dengan urutan harga produk dari yang termurah hingga termahal | - |

---

# Endpoint untuk Create Data Produk

| Method | Route | Description | params |
| --- | --- | --- | --- |
| `POST` |`/create` | Melakukan input sebuah data `produk` baru | - |

---

# Endpoint untuk Update Data Produk

| Method | Route | Description | params |
| --- | --- | --- | --- |
| `PUT` |`/update` | Melakukan update sebuah data `produk` berdasarkan id produk | id_produk |

---

# Endpoint untuk Delete Data Produk

| Method | Route | Description | params |
| --- | --- | --- | --- |
| `DELETE` |`/delete?product_id=` | Melakukan hapus sebuah data `produk` berdasarkan id produk | id_produk |

---

## Cara pengetesan Endpoint pada Postman <br> menggunakan raw data untuk metode POST dan PUT
```bash
- Buka aplikasi Postman
- Lalu pilih Body dan pilih raw
- Piiih formatnya menjadi JSON
```
# Format JSON untuk POST Data Produk
```JSON
{
    "product_name": "Nike Vapor Max 2019",
    "product_price": 1530000,
    "product_description": "Nike Vapor Max 2019 adalah sepatu sneakers Unisex model Low Cut dengan ujung sepatu berbentuk Almond. Dibuat dengan material insole dari Polyster dan outsole dari Rubber, sepatu sneakers Nike ini bisa membuat penampilan kamu terlihat semakin keren. Langsung beli dan dapatkan harga Nike Vapor Max 2019 yang termurah mulai dari IDR1449000 hanya di iPrice Indonesia.",
    "product_quantity": 10
}

```




<br>
