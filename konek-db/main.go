package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Definisikan struct Biodata
type Biodata struct {
	Nim    int    `json:"nim"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
}

// Fungsi untuk mengambil data dari tabel biodata
func selectData(db *sql.DB) []Biodata {
	// Menjalankan query ke database
	selectDB, err := db.Query("SELECT * FROM biodata")
	if err != nil {
		panic(err.Error()) // Menggunakan panic untuk error handling
	}
	defer selectDB.Close() // Pastikan resource database ditutup setelah digunakan

	var bioArray []Biodata // Inisialisasi slice untuk menyimpan hasil query

	// Iterasi melalui hasil query
	for selectDB.Next() {
		var bio Biodata
		// Scan hasil query ke dalam struct Biodata
		err = selectDB.Scan(&bio.Nim, &bio.Nama, &bio.Alamat)
		if err != nil {
			panic(err.Error()) // Menggunakan panic untuk error handling
		}
		// Tambahkan struct Biodata ke slice
		bioArray = append(bioArray, bio)
	}

	// Periksa kesalahan selama iterasi hasil query
	if err = selectDB.Err(); err != nil {
		panic(err.Error()) // Menggunakan panic untuk error handling
	}

	return bioArray
}

// Fungsi untuk memasukkan data ke tabel biodata
func insertData(db *sql.DB) {
	// Menjalankan query untuk memasukkan data
	_, err := db.Exec("INSERT INTO biodata(nama, alamat) VALUES ('Reno', 'Malang')")
	if err != nil {
		panic(err.Error()) // Menggunakan panic untuk error handling
	} else {
		fmt.Println("Berhasil tambah data")
	}
}

func main() {
	// Membuka koneksi ke database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/latihan_go")
	if err != nil {
		panic(err.Error()) // Menggunakan panic untuk error handling
	}
	defer db.Close() // Pastikan koneksi database ditutup setelah selesai

	fmt.Println("berhasil konek")

	// Untuk memasukkan data baru
	// insertData(db)

	// Memanggil fungsi selectData dan menampilkan hasilnya
	data := selectData(db)
	fmt.Print(data[0], data[1])
	// for _, bio := range data {
	// 	fmt.Printf("Nim: %d, Nama: %s, Alamat: %s\n", bio.Nim, bio.Nama, bio.Alamat)
	// }
}
