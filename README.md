# Technical test answer for Qoin Digital Indonesia 

## Soal #1

### Rancangan database dapat dilihat di database_diagram.png
### Teknologi yang digunakan:

Teknologi yang digunakan tergantung dari seberapa besar penggunaan aplikasinya, jika untuk rumah makan sederhana dan request perminutes-nya tidak terlalu besar bisa pakai arsitektur MVC dalam satu framework seperti menggunakan PHP Laravel sehingga aplikasi dapat lebih mudah dimaintain dan cost yang dikeluarkan untuk mendevelop dan memaintain aplikasi juga bisa lebih efisien. Namun jika aplikasinya akan digunakan diperbagai platform (web dan mobile), atau akan digunakan oleh rumah makan dengan skala besar yang cabangnya tersebar di berbagai kota sehingga dalam satu waktu banyak sekali mendapatkan request, maka lebih baik jika memisahkan antara aplikasi server side (back end) dan client side-nya (fron end). Sehingga aplikasi lebih mudah untuk didevelop dan dimaintain karena masing" developer terfokus pada bagian tertentu. teknologi yang bisa digunakan untuk aplikasi dengan skala besar seperti ini bisa menggunakan golang untuk backend, dan react JS untuk front end-nya. Untuk database bisa pakai PostreSQL dan jika dibutuhkan untuk melakukan caching data untuk data yang besar atau querynya membutuhkan waktu yang lama bisa menggunakan Redis. dan untuk data laporan dalam jumlah besar bisa diquery secara otomatis menggunakan cronjob kemudian disimpan di document based database seperti MongoDB, sehingga tidak perlu melakukan query ke RDBMS setiap kali mau mencetak laporan.

### Tambahan Fitur:

1. Management pemesanan bahan baku ke supplier
2. Management User dan roles
3. Management membership pelanggan

## Soal #2

### Script code dapat dilihat pada qoin_digital_test.go