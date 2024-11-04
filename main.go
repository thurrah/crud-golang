package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var books []string
	
func main(){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pilih Opsi: ")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Lihat Semua Buku")
		fmt.Println("3. Hapus buku")
		fmt.Println("4. Keluar")
		fmt.Print("Masukkan pilihan anda: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Masukkan Judul buku baru: ")
			scanner.Scan()
			book := scanner.Text()
			books = append(books, book)
			fmt.Println("Tugas Berhasil Ditambahkan")
		case "2":
			fmt.Println("Daftar buku :")
			for i, book := range books {
				fmt.Printf("%d. Judul : '%s'\n", i+1, book)
			}
		case "3":
			fmt.Println("Masukkan nomor buku yang ingin dihapus: ")
			scanner.Scan()
			numStr := scanner.Text()
			num, err := strconv.Atoi(numStr)
			if err != nil || num < 1 || num < len(books){
				fmt.Println("Nomor buku tidak valid.")
			} else {
				books = append(books[:num-1], books[num:]...)
				fmt.Println("Buku berhasil dihapus!")
			}
		case "4":
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}

	}
	
}