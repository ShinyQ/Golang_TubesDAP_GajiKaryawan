package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

/**
	+-------------------------------------+
	|    DECLARE VARIABLE AND RECORD	  |
    +-------------------------------------+
**/

var (
	Menu         int
	itemKaryawan []Karyawan
	itemGaji     []Gaji

	dataKaryawan = DataKaryawan{itemKaryawan}

	ErrorPrint   = color.New(color.FgRed).Add(color.BgWhite)
	SuccessPrint = color.New(color.FgBlue).Add(color.BgWhite)
	scanner      = bufio.NewScanner(os.Stdin)
)

type Karyawan struct {
	Golongan, Umur, JumlahAnak int
	Nama, Alamat, KodePegawai  string
}

type DataKaryawan struct {
	ItemsKaryawan []Karyawan
}

type Gaji struct {
	KodePegawai, Bulan  string
	JamKerja, TotalGaji int
}

type DataGaji struct {
	ItemsGaji []Gaji
}

/**
	+----------------------------------------+
	| END OF VARIABLE AND RECORD DECLARATION |
    +----------------------------------------+
	|      START OF KARYAWAN FUNCTION		 |
    +----------------------------------------+
**/

func (dataKaryawan *DataKaryawan) tambahKaryawan(itemKaryawan Karyawan) []Karyawan {
	dataKaryawan.ItemsKaryawan = append(dataKaryawan.ItemsKaryawan, itemKaryawan)
	return dataKaryawan.ItemsKaryawan
}

func inputKaryawan() {
	var (
		i, Golongan, Umur, JumlahAnak int
		Nama, Alamat, KodePegawai     string
		inputLagi                     string
		validGolongan, Selesai        bool
	)

	for i = len(dataKaryawan.ItemsKaryawan); Selesai != true; i++ {
		fmt.Println("")
		fmt.Println("Masukkan Data Pegawai")

		fmt.Print("Kode Pegawai : \t")
		fmt.Scanln(&KodePegawai)

		fmt.Print("Golongan : \t")
		fmt.Scanln(&Golongan)
		for validGolongan != true {
			if Golongan > 3 && Golongan > 0 {
				ErrorPrint.Println(" Golongan Tersebut Tidak Ada ! ")
				fmt.Print("Golongan : \t")
				fmt.Scanln(&Golongan)
			} else {
				validGolongan = true
			}
		}
		fmt.Print("Umur Pegawai: \t")
		fmt.Scanln(&Umur)

		fmt.Print("Nama Pegawai: \t")
		scanner.Scan()
		Nama = scanner.Text()

		fmt.Print("Jumlah Anak : \t")
		fmt.Scanln(&JumlahAnak)

		fmt.Print("Alamat : \t")
		scanner.Scan()
		Alamat = scanner.Text()

		karyawan := Karyawan{
			Nama:        Nama,
			KodePegawai: KodePegawai,
			Golongan:    Golongan,
			Umur:        Umur,
			JumlahAnak:  JumlahAnak,
			Alamat:      Alamat,
		}

		dataKaryawan.tambahKaryawan(karyawan)
		itemKaryawan = dataKaryawan.ItemsKaryawan

		fmt.Println("\nData Berhasil Diinputkan :")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Kode Pegawai", "Nama", "Golongan", "Umur", "Jumlah Anak", "Alamat"})
		table.Append(
			[]string{
				dataKaryawan.ItemsKaryawan[i].KodePegawai,
				dataKaryawan.ItemsKaryawan[i].Nama,
				strconv.Itoa(dataKaryawan.ItemsKaryawan[i].Golongan),
				strconv.Itoa(dataKaryawan.ItemsKaryawan[i].Umur),
				strconv.Itoa(dataKaryawan.ItemsKaryawan[i].JumlahAnak),
				dataKaryawan.ItemsKaryawan[i].Alamat,
			},
		)
		table.Render()

		fmt.Print("\nInput Lagi ? (Ya / Tidak) : ")
		fmt.Scanln(&inputLagi)

		if inputLagi == "Tidak" || inputLagi == "tidak" {
			Selesai = true
		}
	}
	menu()
}

func tampilKaryawan() {
	if len(itemKaryawan) != 0 {

		SuccessPrint.Println(" Terdapat", len(itemKaryawan), "Data Karyawan \n")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Kode Pegawai", "Nama", "Golongan", "Umur", "Jumlah Anak", "Alamat"})

		for i := 0; i < len(itemKaryawan); i++ {
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
		}
		table.Render()
	} else {
		ErrorPrint.Println(" Belum Ada Data Karyawan ")
	}

	menu()
}

/**
	+-------------------------------------+
	|    	MENU AND MAIN FUNCTION    	  |
    +-------------------------------------+
**/

func menu() {
	fmt.Println("+----------------------------------+--------------------------------------+")
	fmt.Println("| Menu 1 : Input Data Karyawan	   |	 Menu 4 : Input Gaji Karyawan	  |")
	fmt.Println("| Menu 2 : Lihat Data Karyawan	   |	 Menu 5 : Histori Data Gaji	  |")
	fmt.Println("| Menu 3 : Cari Data Karyawan 	   |	 Menu 6 : Cari Histori Data Gaji  |")
	fmt.Println("+----------------------------------+--------------------------------------+\n")

	fmt.Print("Silahakan Pilih Menu : ")
	fmt.Scanln(&Menu)
	fmt.Println("")

	if Menu == 1 {
		inputKaryawan()
	} else if Menu == 2 {
		tampilKaryawan()
	}
}

func main() {
	menu()
}
