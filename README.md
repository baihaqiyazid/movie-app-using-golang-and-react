# Movie App


Aplikasi sederhana CRUD Movie List dengan Golang dan ReactJS yang dihubungkan dengan API adalah sebuah aplikasi yang memungkinkan pengguna untuk melakukan operasi dasar CRUD (Create, Read, Update, Delete) terhadap daftar film menggunakan antarmuka pengguna yang dibangun dengan ReactJS. Aplikasi ini menggunakan Golang sebagai backend untuk menyediakan API yang berfungsi sebagai jembatan antara frontend ReactJS dan database.

Berikut adalah deskripsi umum tentang bagaimana aplikasi tersebut dapat beroperasi:

1. Tampilan Daftar Film: Aplikasi akan menampilkan daftar film yang tersedia dalam antarmuka pengguna ReactJS. Informasi yang ditampilkan untuk setiap film dapat mencakup judul, sinopsis, dan genre. Data ini akan diperoleh dari backend melalui API.

2. Penambahan Film: Pengguna dapat menambahkan film baru ke dalam daftar menggunakan formulir yang disediakan. Informasi film seperti judul, sinopsis, dan genre akan dimasukkan ke dalam formulir. Setelah pengguna mengirimkan formulir, data akan dikirimkan ke backend melalui API untuk ditambahkan ke database.

3. Pembaruan Film: Pengguna dapat mengedit informasi film yang ada dalam daftar. Ketika pengguna memilih untuk mengedit film, aplikasi akan menampilkan formulir pra-diisi dengan informasi film yang ada. Pengguna dapat mengubah informasi apa pun dan menyimpan perubahan yang akan diteruskan ke backend melalui API untuk memperbarui data film dalam database.

4. Penghapusan Film: Pengguna dapat menghapus film dari daftar dengan menekan tombol hapus yang sesuai pada tampilan film. Aplikasi akan mengirimkan permintaan penghapusan ke backend melalui API, dan film yang dipilih akan dihapus dari database.

5. Penyimpanan Data Film: Backend Golang akan bertanggung jawab untuk mengelola data film dan menyimpannya dalam database MySql. Golang akan menyediakan API endpoint yang menerima permintaan dari frontend ReactJS dan menjalankan operasi CRUD yang sesuai pada database.

6. Koneksi Antara Frontend dan Backend: Frontend ReactJS akan berkomunikasi dengan backend Golang melalui HTTP API. Permintaan HTTP seperti GET, POST, PUT, dan DELETE akan digunakan untuk memperoleh, menambahkan, memperbarui, dan menghapus data film. Backend akan mengirimkan respon yang sesuai dengan permintaan yang diterima oleh frontend.

7. Penting untuk dicatat bahwa deskripsi di atas adalah gambaran umum tentang cara kerja aplikasi CRUD Movie List dengan Golang dan ReactJS. Implementasi spesifik dan rincian teknis dapat bervariasi tergantung pada preferensi dan kebutuhan proyek yang spesifik.

## How to run the program?

1. go to path `server/cmd/api` and write `go run main.go` in terminal
2. go to path `client` and write `npm start` in terminal
3. open your browser localhot:3000
4. success
