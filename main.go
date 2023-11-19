package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Input nama pengguna
	fmt.Print("Masukkan nama Anda: ")
	reader := bufio.NewReader(os.Stdin)
	nama, _ := reader.ReadString('\n')

	// Kumpulan soal, pilihan jawaban, dan jawaban yang benar
	soal := []string{
		"Apa ibu kota Indonesia?",
		"Siapakah presiden Indonesia saat ini?",
		"Berapa jumlah provinsi di Indonesia?",
	}
	pilihan := [][]string{
		{"1. Jakarta", "2. Bandung", "3. Surabaya"},
		{"1. Joko Widodo", "2. Megawati Soekarnoputri", "3. Susilo Bambang Yudhoyono"},
		{"1. 28", "2. 34", "3. 42"},
	}
	jawabanBenar := []string{"1", "1", "2"}

	// Menghitung jumlah soal
	jumlahSoal := len(soal)

	// Inisialisasi hasil jawaban
	jawabanBenarCount := 0
	jawabanSalahCount := 0

	// Mengolah setiap soal
	for i := 0; i < jumlahSoal; i++ {
		fmt.Printf("\nSoal %d: %s\n", i+1, soal[i])
		for _, pilihanJawaban := range pilihan[i] {
			fmt.Println(pilihanJawaban)
		}
		fmt.Print("Jawaban Anda: ")
		jawabanPengguna, _ := reader.ReadString('\n')

		// Membandingkan jawaban pengguna dengan jawaban yang benar
		jawabanPengguna = strings.TrimSpace(jawabanPengguna)
		if strings.EqualFold(jawabanPengguna, jawabanBenar[i]) {
			jawabanBenarCount++
		} else {
			jawabanSalahCount++
		}
	}
	// Menampilkan keterangan hasil
	fmt.Printf("\n%s, berikut adalah hasilnya:\n", strings.TrimSpace(nama))
	fmt.Println("Jumlah Soal: ", jumlahSoal)
	fmt.Println("Jawaban Benar: ", jawabanBenarCount)
	fmt.Println("Jawaban Salah: ", jawabanSalahCount)
}
