package main

import (
	"fmt"
	"golangbeginer/structinterface/rumus"
)

func main() {
	var nilaiPanjang float64
	fmt.Printf("Masukkan Panjang: ")
	fmt.Scan(&nilaiPanjang)

	var nilaiLebar float64
	fmt.Printf("Masukkan Lebar: ")
	fmt.Scan(&nilaiLebar)

	var nilaiTinggi float64
	fmt.Printf("Masukkan Tinggi: ")
	fmt.Scan(&nilaiTinggi)

	if nilaiPanjang <= 0 && nilaiLebar <= 0 && nilaiTinggi <= 0 {
		fmt.Println("Angka tidak boleh kosong/ tidak boleh string")
	} else {

		bangun := rumus.Bangun{Panjang: nilaiPanjang, Lebar: nilaiLebar, Tinggi: nilaiTinggi}
		fmt.Println("===================")
		fmt.Println("Hasilnya Adalah: ")
		hitung(bangun)

	}

}
func hitung(nilai rumus.Hitung) {
	fmt.Println("Luas : ", nilai.Luas())
	fmt.Println("Keliling : ", nilai.Keliling())
	fmt.Println("Volume : ", nilai.Volume())
}
