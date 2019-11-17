package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

/**
	+-------------------------------------+
	|    DECLARE VARIABLE AND RECORD	  |
    +-------------------------------------+
**/

var (
	Menu int // Variable Menu Untuk Berpindah Function Pada Main Sesuai Dengan Angka Yang Dimasukkan

	itemKaryawan []Karyawan // Array itemKaryawan berupa record karyawan yang akan menampung data item Karyawan
	itemGaji     []Gaji     // Array itemGaji berupa record gaji karyawan yang akan menampung data Gaji Karyawan

	ErrorPrint   = color.New(color.FgRed).Add(color.BgWhite)  // Implementasi Penggunaan Library Color Untuk Error
	SuccessPrint = color.New(color.FgBlue).Add(color.BgWhite) // Implementasi Penggunaan Library Color Untuk Sukses Menampilkan Data
	scanner      = bufio.NewScanner(os.Stdin)                 // Deklarasi Scanner untuk data yang memerlukan spasi dalam inputannya
)

// Deklarasi record Karyawan
type Karyawan struct {
	Golongan, Umur, JumlahAnak int
	Nama, Alamat, KodePegawai  string
}

// Deklarasi record Gaji
type Gaji struct {
	KodePegawai, Bulan  string
	JamKerja, TotalGaji int
}

/**
	+----------------------------------------+
	| END OF VARIABLE AND RECORD DECLARATION |
    +----------------------------------------+
	|      START OF KARYAWAN FUNCTION		 |
    +----------------------------------------+
**/

// Function Untuk Menginputkan Data Karyawan
func inputKaryawan() {
	var (
		i, Golongan, Umur, JumlahAnak                int
		Nama, Alamat, KodePegawai                    string
		inputLagi                                    string
		validGolongan, validUmur, validAnak, Selesai bool
	)

	// Melakukan Perulangan Inputan karyawan sampai user selesai menginputkan data karyawan (Parameter Selesai)
	for i = len(itemKaryawan); Selesai != true; i++ {
		fmt.Println("Masukkan Data Pegawai")

		fmt.Print("Kode Pegawai : \t")
		fmt.Scanln(&KodePegawai)

		fmt.Print("Golongan : \t")
		fmt.Scanln(&Golongan)

		// Perulangan Validasi Golongan
		for validGolongan != true {
			if Golongan > 3 || Golongan <= 0 {
				ErrorPrint.Println("\n Golongan Tersebut Tidak Ada ! \n")
				fmt.Print("Golongan : \t")
				fmt.Scanln(&Golongan)
			} else {
				validGolongan = true
			}
		}

		fmt.Print("Umur Pegawai: \t")
		fmt.Scanln(&Umur)

		// Perulangan Validasi Umur
		for validUmur != true {
			if Umur < 1 {
				fmt.Print("Umur Pegawai: \t")
				fmt.Scan(&Umur)
			} else {
				validUmur = true
			}
		}

		fmt.Print("Nama Pegawai: \t")
		scanner.Scan()
		Nama = scanner.Text()

		fmt.Print("Jumlah Anak : \t")
		fmt.Scanln(&JumlahAnak)

		// Perulangan Validasi Jumlah Anak
		for validAnak != true {
			if JumlahAnak < 0 {
				ErrorPrint.Println("\n Jumlah Anak Tidak Valid ! \n")
				fmt.Print("Jumlah Anak : \t")
				fmt.Scanln(&JumlahAnak)
			} else {
				validAnak = true
			}
		}

		fmt.Print("Alamat : \t")
		scanner.Scan()
		Alamat = scanner.Text()

		// Proses membuat record karyawan baru sesuai inputan user
		karyawan := Karyawan{
			Nama:        Nama,
			KodePegawai: KodePegawai,
			Golongan:    Golongan,
			Umur:        Umur,
			JumlahAnak:  JumlahAnak,
			Alamat:      Alamat,
		}

		// Proses menambahkan karyawan kedalam array ItemKaryawan dengan append (Slice)
		itemKaryawan = append(itemKaryawan, karyawan)

		// Proses menampilkan data karyawan yang telah di inputkan sebelumnya
		fmt.Println("\nData Berhasil Diinputkan :")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Kode Pegawai", "Nama", "Golongan", "Umur", "Jumlah Anak", "Alamat"})
		table.Append(
			[]string{
				itemKaryawan[i].KodePegawai,
				itemKaryawan[i].Nama,
				strconv.Itoa(itemKaryawan[i].Golongan),
				strconv.Itoa(itemKaryawan[i].Umur),
				strconv.Itoa(itemKaryawan[i].JumlahAnak),
				itemKaryawan[i].Alamat,
			},
		)
		table.Render()

		// Konfirmasi Apakah User Ingin Menginput Kembali Karyawan Baru
		fmt.Print("\nInput Lagi ? (Y / T) : ")
		fmt.Scanln(&inputLagi)
		fmt.Println("")
		if inputLagi == "T" || inputLagi == "t" {
			Selesai = true
		}
	}
	menu()
}

/*
   Function Proses mencari karyawan sesuai dengan kode Karyawan
   dimana akan melakukan return pada array indeks ke berapa karyawan tersebut
*/
func prosesCariKaryawan(kode string) int {
	var Selesai bool
	var data int

	for i := 0; i < len(itemKaryawan) && Selesai != true; i++ {
		if itemKaryawan[i].KodePegawai == kode {
			data = i
			Selesai = true
		} else {
			data = -1
		}
	}
	return data
}

func cariKaryawan() {
	// Function cari karyawan dimana user menginputkan kode karyawan yang akan dicari
	var KodePegawai string
	var data int

	fmt.Print("Masukkan Kode Pegawai : ")
	fmt.Scanln(&KodePegawai)

	if len(itemKaryawan) == 0 {
		data = -1
	} else {
		data = prosesCariKaryawan(KodePegawai)
	}

	// Proses Percabangan apabila indeks array karyawan ditemukan atau tidak
	if data != -1 {

		// Proses menampilkan data karyawan yang telah ditemukan sesuai dengan indeks array yang di return
		SuccessPrint.Println("\n Data Ditemukan : \n")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Kode Pegawai", "Nama", "Golongan", "Umur", "Jumlah Anak", "Alamat"})
		table.Append(
			[]string{
				itemKaryawan[data].KodePegawai,
				itemKaryawan[data].Nama,
				strconv.Itoa(itemKaryawan[data].Golongan),
				strconv.Itoa(itemKaryawan[data].Umur),
				strconv.Itoa(itemKaryawan[data].JumlahAnak),
				itemKaryawan[data].Alamat,
			},
		)
		table.Render()
	} else {
		ErrorPrint.Println("\n Maaf, Data Karyawan Tidak Ditemukan ")
	}
	menu()
}

/*
    Function Tampil Karyawan dimana function ini menampilkan seluruh data karyawan yang
	terdapat dalam array itemKaryawan dengan melakukan perulangan
*/
func tampilKaryawan() {
	if len(itemKaryawan) != 0 {
		SuccessPrint.Println(" Terdapat", len(itemKaryawan), "Data Karyawan \n")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"No", "Kode Pegawai", "Nama", "Golongan", "Umur", "Jumlah Anak", "Alamat"})

		for i := 0; i < len(itemKaryawan); i++ {
			table.Append(
				[]string{
					strconv.Itoa(i+1) + ".",
					itemKaryawan[i].KodePegawai,
					itemKaryawan[i].Nama,
					strconv.Itoa(itemKaryawan[i].Golongan),
					strconv.Itoa(itemKaryawan[i].Umur),
					strconv.Itoa(itemKaryawan[i].JumlahAnak),
					itemKaryawan[i].Alamat,
				},
			)
		}
		table.Render()
	} else {
		ErrorPrint.Println(" Belum Ada Data Karyawan ")
	}

	menu()
}

/*
   Function Sort Golongan Karyawan Untuk Melakukan sorting data karyawan
   sesuai dengan pencarian golongan secara ascending ( kecil ke terbesar )
*/
func sortKaryawanGolongan() {
	var (
		sortKaryawan []Karyawan
		sorted       bool
		Golongan, n  int
	)
	fmt.Print("Masukkan Golongan : ")
	fmt.Scanln(&Golongan)

	// Melakukan pencarian data karyawan berdasarkan golongan yang diinputkan user
	for i := 0; i < len(itemKaryawan); i++ {
		if itemKaryawan[i].Golongan == Golongan {

			// Data karyawan yang sesuai dengan golongan yang dicari dibuat kemudian dimasukkan datanya kedalam array sementara sortKaryawan
			karyawan := Karyawan{
				Nama:        itemKaryawan[i].Nama,
				KodePegawai: itemKaryawan[i].KodePegawai,
				Golongan:    itemKaryawan[i].Golongan,
				Umur:        itemKaryawan[i].Umur,
				JumlahAnak:  itemKaryawan[i].JumlahAnak,
				Alamat:      itemKaryawan[i].Alamat,
			}

			// Proses memasukkan data karyawan yang sesuai kedalam array sortKaryawan
			sortKaryawan = append(sortKaryawan, karyawan)
		}
	}

	n = len(sortKaryawan)
	// Seleksi apakah data pada sortKaryawan ada atau tidak
	if n != 0 {

		// Proses looping untuk sorting array sortKaryawan dengan menggunakan metode selection sort
		for !sorted {
			swapped := false
			for i := 0; i < n-1; i++ {
				if sortKaryawan[i].Nama > sortKaryawan[i+1].Nama {
					sortKaryawan[i+1], sortKaryawan[i] = sortKaryawan[i], sortKaryawan[i+1]
					swapped = true
				}
			}
			if !swapped {
				sorted = true
			}
			n--
		}

		// Proses melakukan looping data array sortKaryawan yang telah di sorting
		SuccessPrint.Println("\n Terdapat", n+1, "Karyawan Golongan", Golongan, ":\n")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"No", "Kode Pegawai", "Nama", "Golongan", "Umur", "Jumlah Anak", "Alamat"})
		for i := 0; i < len(sortKaryawan); i++ {
			table.Append(
				[]string{
					strconv.Itoa(i+1) + ".",
					sortKaryawan[i].KodePegawai,
					sortKaryawan[i].Nama,
					strconv.Itoa(sortKaryawan[i].Golongan),
					strconv.Itoa(sortKaryawan[i].Umur),
					strconv.Itoa(sortKaryawan[i].JumlahAnak),
					sortKaryawan[i].Alamat,
				},
			)
		}
		table.Render()
	} else {
		ErrorPrint.Println("\n Data Dengan Golongan Tersebut Tidak Ditemukan ")
	}

	menu()
}

/**
	+-------------------------------------+
	|		END OF KARYAWAN FUNCTION	  |
    +-------------------------------------+
	|		START OF GAJI FUNCTION		  |
	+-------------------------------------+
**/

// Function untuk melakukan validasi inputan bulan penggajian
func validasiBulan(Bulan string) bool {
	var validBulan bool
	arrBulan := [12]string{"Janurai", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	for j := 0; j < len(arrBulan); j++ {
		if Bulan == arrBulan[j] {
			validBulan = true
		}
	}
	return validBulan
}

// Function untuk melakukan input gaji karyawan sesuai kode pegawai yang dimasukkan pengguna
func inputGaji() {
	var (
		KodePegawai, Bulan                         string
		JamKerja, TotalGaji, GajiTetap             int
		Selesai, validKode, validBulan, validWaktu bool
		i, Golongan, JumlahAnak, index             int
		inputLagi                                  string
	)

	// Melakukan Perulangan inputan gaji sampai user mengkonfirmasi  selesai melakukan input
	for i = len(itemGaji); Selesai != true; i++ {
		fmt.Println("Masukkan Gaji Pegawai")

		fmt.Print("Kode Pegawai \t : ")
		fmt.Scanln(&KodePegawai)

		// Validasi apakah kode pegawai untuk input gaji ada atau tidak
		for validKode != true {
			if prosesCariKaryawan(KodePegawai) == -1 {
				ErrorPrint.Println("\n Kode Pegawai Tidak Valid ! \n")
				fmt.Print("Kode Pegawai \t : ")
				fmt.Scanln(&KodePegawai)
			} else {
				validKode = true
			}
		}

		fmt.Print("Bulan \t\t : ")
		fmt.Scanln(&Bulan)

		// Validasi bulan inputan user
		for validBulan != true {
			if validasiBulan(Bulan) {
				validBulan = true
			} else {
				ErrorPrint.Println("\n Bulan Tersebut Tidak Ada ! \n")
				fmt.Print("Bulan \t\t : ")
				fmt.Scanln(&Bulan)
			}
		}

		fmt.Print("Jumlah Jam Kerja : ")
		fmt.Scanln(&JamKerja)

		// Validasi jumlah waktu kerja
		for validWaktu != true {
			if JamKerja >= 0 {
				validWaktu = true
			} else {
				ErrorPrint.Println("\n Jumlah Waktu Kerja Tidak Valid ! \n")
				fmt.Print("Jumlah Jam Kerja : ")
				fmt.Scanln(&JamKerja)
			}
		}

		// Melakukan pencarian data karyawan untuk mendapatkan Golongan dan Jumlah Anak untuk proses hitung gaji
		index = prosesCariKaryawan(KodePegawai)
		Golongan = itemKaryawan[index].Golongan
		JumlahAnak = itemKaryawan[index].JumlahAnak

		if Golongan == 1 {
			TotalGaji = JamKerja * 5000
			TotalGaji = TotalGaji + (500000 * JumlahAnak)
			GajiTetap = 500000
		} else if Golongan == 2 {
			TotalGaji = JamKerja * 3000
			TotalGaji = TotalGaji + (400000 * JumlahAnak)
			GajiTetap = 300000
		} else {
			TotalGaji = JamKerja * 20000
			TotalGaji = TotalGaji + (300000 * JumlahAnak)
			GajiTetap = 250000
		}

		if Bulan == "Januari" || Bulan == "Agustus" || Bulan == "Oktober" {
			TotalGaji = TotalGaji + (GajiTetap * 10 / 100)
		}

		// Proses memasukkan data gaji karyawan sesuai record dengan inputan pengguna
		gaji := Gaji{
			KodePegawai: KodePegawai,
			Bulan:       Bulan,
			JamKerja:    JamKerja,
			TotalGaji:   TotalGaji,
		}

		// Proses memasukkan data gaji pengguna yang telah di inputkan kedalam array itemGaji
		itemGaji = append(itemGaji, gaji)

		// Proses Menampilkan data gaji yang telah di inputkan pengguna sebelumnya
		fmt.Println("\nData Berhasil Diinputkan :")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Kode Pegawai", "Bulan", "Jam Kerja", "Total Gaji"})

		table.Append(
			[]string{
				itemGaji[i].KodePegawai,
				itemGaji[i].Bulan,
				strconv.Itoa(itemGaji[i].JamKerja),
				"Rp" + humanize.Comma(int64(itemGaji[i].TotalGaji)) + ",00",
			},
		)
		table.Render()

		// Proses validasi apakah pengguna akan melakukan input gaji pegawai lagi
		fmt.Print("\nInput Lagi ? (Y / T) : ")
		fmt.Scanln(&inputLagi)
		fmt.Println("")

		if inputLagi == "T" || inputLagi == "t" {
			Selesai = true
		}
	}
	menu()
}

// Function menampilkan seluruh data gaji karyawan pada array item gaji
func tampilGaji() {

	// Proses percabangan apakah array itemGaji berisi data gaji atau tidak
	if len(itemGaji) != 0 {

		// Proses looping data gaji karyawan
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"No", "Kode Pegawai", "Bulan", "Jam Kerja", "Total Gaji"})

		for i := 0; i < len(itemGaji); i++ {
			table.Append(
				[]string{
					strconv.Itoa(i+1) + ".",
					itemGaji[i].KodePegawai,
					itemGaji[i].Bulan,
					strconv.Itoa(itemGaji[i].JamKerja),
					"Rp" + humanize.Comma(int64(itemGaji[i].TotalGaji)) + ",00",
				},
			)
		}
		table.Render()
	} else {
		ErrorPrint.Println(" Belum Ada Data Gaji ")
	}

	menu()
}

// Function proses melakukan pencarian gaji sesuai kode pegawai dan akan mengembalikan nilai indeks array dari data pegawai
func prosesCariGaji(kode string) int {
	var Selesai bool
	var data int

	// Perulangan untuk melakukan pencarian gaji sesuai dengan kode pegawai
	for i := 0; i < len(itemGaji) && Selesai != true; i++ {
		if itemGaji[i].KodePegawai == kode {
			data = i
			Selesai = true
		} else {
			data = -1
		}
	}
	return data
}

// Function Cari Gaji untuk melakukan inputan kode pegawai serta menampilkan seluruh data gaji sesuai kode pegawai
func cariGaji() {
	var KodePegawai string
	var data int

	fmt.Print("Masukkan Kode Pegawai : ")
	fmt.Scanln(&KodePegawai)

	// Seleksi dan Proses Pencarian Data Gaji Sesuai Kode Pegawai
	if len(itemGaji) == 0 {
		data = -1
	} else {
		data = prosesCariGaji(KodePegawai)
	}

	// Proses percabangan apakah ditemukan kode pegawai atau tidak
	if data != -1 {

		// Proses menampilkan data gaji pegawai sesuai dengan indeks array yang ditemukan pada pencarian kode pegawai
		SuccessPrint.Println("\n Data Ditemukan : \n")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Kode Pegawai", "Bulan", "Jam Kerja", "Total Gaji"})

		for i := 0; i < len(itemGaji); i++ {
			if itemGaji[i].KodePegawai == KodePegawai {
				table.Append(
					[]string{
						itemGaji[i].KodePegawai,
						itemGaji[i].Bulan,
						strconv.Itoa(itemGaji[i].JamKerja),
						"Rp" + humanize.Comma(int64(itemGaji[i].TotalGaji)) + ",00",
					},
				)
			}
		}
		table.Render()
	} else {
		ErrorPrint.Println("\n Maaf, Data Gaji Tidak Ditemukan ")
	}

	menu()
}

// Proses melakukan sorting berdasarkan pencarian bulan berdasarkan kode pegawai secara descendend (Tinggi Ke Rendah)
func sortGajiBulan() {
	var (
		sortGaji []Gaji
		sorted   = false
		Bulan    string
		n        int
	)
	fmt.Print("Masukkan Bulan Penggajian : ")
	fmt.Scanln(&Bulan)

	// Proses melakukan pencarian dan insert data array sementara yang merupakan data gaji pada bulan yang dicari pengguna
	for i := 0; i < len(itemGaji); i++ {
		if itemGaji[i].Bulan == Bulan {

			// Proses membuat data array baru sesuai dengan indeks data array sesuai bulan yang dicari
			gaji := Gaji{
				KodePegawai: itemGaji[i].KodePegawai,
				Bulan:       itemGaji[i].Bulan,
				JamKerja:    itemGaji[i].JamKerja,
				TotalGaji:   itemGaji[i].TotalGaji,
			}

			// Proses memasukkan data array baru kedalam array data sort gaji sementara
			sortGaji = append(sortGaji, gaji)
		}
	}

	n = len(sortGaji)
	// Proses seleksi apakah pencarian array sort gaji menghasilkan data array atau tidak
	if n != 0 {

		// Proses melakukan sorting data pada array sort gaji berdasarkan kode pegawai secara descendend
		for !sorted {
			swapped := false
			for i := 0; i < n-1; i++ {
				if sortGaji[i].KodePegawai < sortGaji[i+1].KodePegawai {
					sortGaji[i+1], sortGaji[i] = sortGaji[i], sortGaji[i+1]
					swapped = true
				}

			}
			if !swapped {
				sorted = true
			}
			n--
		}

		// Proses menampilkan data array sort gaji yang telah di sorting
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"No", "Kode Pegawai", "Bulan", "Jam Kerja", "Total Gaji"})
		for i := 0; i < len(sortGaji); i++ {
			table.Append(
				[]string{
					strconv.Itoa(i+1) + ".",
					sortGaji[i].KodePegawai,
					sortGaji[i].Bulan,
					strconv.Itoa(sortGaji[i].JamKerja),
					"Rp" + humanize.Comma(int64(sortGaji[i].TotalGaji)) + ",00",
				},
			)
		}
		table.Render()
	} else {
		ErrorPrint.Println("\n Data Dengan Bulan Tersebut Tidak Ditemukan ")
	}

	menu()
}

/**
	+-------------------------------------+
	|    	MENU AND MAIN FUNCTION    	  |
    +-------------------------------------+
**/

// Function menu awal yang bertujuan untuk user memilih menu yang di inginkan
func menu() {

	fmt.Println("\n+----------------------------------+--------------------------------------+-------------------------------------------------------+")
	fmt.Println("| Menu 1 : Input Data Karyawan	   |	 Menu 4 : Input Gaji Karyawan	  |	 Menu 7 : Cari Karyawan Berdasarkan Golongan	  |")
	fmt.Println("| Menu 2 : Lihat Data Karyawan	   |	 Menu 5 : Histori Data Gaji	  |	 Menu 8 : Cari Gaji Berdasarkan Bulan	      	  |")
	fmt.Println("| Menu 3 : Cari Data Karyawan 	   |	 Menu 6 : Cari Histori Data Gaji  |  	 Menu 9 : Keluar Program	  		  |")
	fmt.Println("+----------------------------------+--------------------------------------+-------------------------------------------------------+\n")

	fmt.Print("Silahakan Pilih Menu : ")
	fmt.Scanln(&Menu)
	fmt.Println("")

	if Menu == 1 {
		inputKaryawan()
	} else if Menu == 2 {
		tampilKaryawan()
	} else if Menu == 3 {
		cariKaryawan()
	} else if Menu == 4 {
		inputGaji()
	} else if Menu == 5 {
		tampilGaji()
	} else if Menu == 6 {
		cariGaji()
	} else if Menu == 7 {
		sortKaryawanGolongan()
	} else if Menu == 8 {
		sortGajiBulan()
	} else if Menu == 9 {
		defer SuccessPrint.Println(" Sukses Keluar Program ")
	} else {
		ErrorPrint.Println(" Menu Tersebut Tidak Ada ! ")
		menu()
	}

}

func main() {
	menu()
}
