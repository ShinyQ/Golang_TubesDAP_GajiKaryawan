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
	dataGaji     = DataGaji{itemGaji}

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

func prosesCariKaryawan(kode string) int {
	var Selesai bool
	var data int

	for i := 0; i < len(dataKaryawan.ItemsKaryawan) && Selesai != true; i++ {
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
	var KodePegawai string
	var i int
	fmt.Print("Masukkan Kode Pegawai : ")
	fmt.Scanln(&KodePegawai)

	if prosesCariKaryawan(KodePegawai) != -1 {
		i = prosesCariKaryawan(KodePegawai)
		SuccessPrint.Println("\n Data Ditemukan : \n")
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
	} else {
		ErrorPrint.Println("\n Maaf, Data Karyawan Tidak Ditemukan ")
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
	|		END OF KARYAWAN FUNCTION	  |
    +-------------------------------------+
	|		START OF GAJI FUNCTION		  |
	+-------------------------------------+
**/

func (dataGaji *DataGaji) tambahGaji(itemGaji Gaji) []Gaji {
	dataGaji.ItemsGaji = append(dataGaji.ItemsGaji, itemGaji)
	return dataGaji.ItemsGaji
}

func inputGaji() {
	var (
		KodePegawai, Bulan             string
		JamKerja, TotalGaji, GajiTetap int
		Selesai                        bool
		i, Golongan, JumlahAnak, index int
		inputLagi                      string
	)

	for i = len(dataGaji.ItemsGaji); Selesai != true; i++ {
		fmt.Println("")
		fmt.Println("Masukkan Gaji Pegawai")

		fmt.Print("Kode Pegawai \t : ")
		fmt.Scanln(&KodePegawai)
		for Selesai != true {

			if prosesCariKaryawan(KodePegawai) == -1 {
				fmt.Print("Kode Pegawai \t : ")
				fmt.Scanln(&KodePegawai)
			} else {
				Selesai = true
			}
		}

		fmt.Print("Bulan \t\t : ")
		fmt.Scanln(&Bulan)

		fmt.Print("Jumlah Jam Kerja : ")
		fmt.Scanln(&JamKerja)

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

		gaji := Gaji{
			KodePegawai: KodePegawai,
			Bulan:       Bulan,
			JamKerja:    JamKerja,
			TotalGaji:   TotalGaji,
		}

		dataGaji.tambahGaji(gaji)
		itemGaji = dataGaji.ItemsGaji

		fmt.Println("\nData Berhasil Diinputkan :")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Kode Pegawai", "Bulan", "Jam Kerja", "Total Gaji"})

		table.Append(
			[]string{
				dataGaji.ItemsGaji[i].KodePegawai,
				dataGaji.ItemsGaji[i].Bulan,
				strconv.Itoa(dataGaji.ItemsGaji[i].JamKerja),
				strconv.Itoa(dataGaji.ItemsGaji[i].TotalGaji),
			},
		)
		table.Render()

		fmt.Print("\nInput Lagi ? (Y / T) : ")
		fmt.Scanln(&inputLagi)

		if inputLagi == "T" || inputLagi == "t" {
			Selesai = true
		}
	}
	menu()
}

func tampilGaji() {
	if len(itemGaji) != 0 {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Kode Pegawai", "Bulan", "Jam Kerja", "Total Gaji"})

		for i := 0; i < len(itemGaji); i++ {
			table.Append(
				[]string{
					itemGaji[i].KodePegawai,
					itemGaji[i].Bulan,
					strconv.Itoa(itemGaji[i].JamKerja),
					strconv.Itoa(itemGaji[i].TotalGaji),
				},
			)
		}
		table.Render()
	} else {
		ErrorPrint.Println(" Belum Ada Data Gaji ")
	}

	menu()
}

func prosesCariGaji(kode string) int {
	var Selesai bool
	var data int

	for i := 0; i < len(dataGaji.ItemsGaji) && Selesai != true; i++ {
		if itemGaji[i].KodePegawai == kode {
			data = i
			Selesai = true
		} else {
			data = -1
		}
	}
	return data
}

func cariGaji() {
	var KodePegawai string
	fmt.Print("Masukkan Kode Pegawai : \t")
	fmt.Scanln(&KodePegawai)

	if prosesCariGaji(KodePegawai) != -1 {
		SuccessPrint.Println("\n Data Ditemukan : \n")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Kode Pegawai", "Bulan", "Jam Kerja", "Total Gaji"})

		for i := 0; i < len(dataGaji.ItemsGaji); i++ {
			if itemGaji[i].KodePegawai == KodePegawai {
				table.Append(
					[]string{
						itemGaji[i].KodePegawai,
						itemGaji[i].Bulan,
						strconv.Itoa(itemGaji[i].JamKerja),
						strconv.Itoa(itemGaji[i].TotalGaji),
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

/**
	+-------------------------------------+
	|    	MENU AND MAIN FUNCTION    	  |
    +-------------------------------------+
**/

func menu() {

	fmt.Println("\n+----------------------------------+--------------------------------------+")
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
	} else if Menu == 3 {
		cariKaryawan()
	} else if Menu == 4 {
		inputGaji()
	} else if Menu == 5 {
		tampilGaji()
	} else if Menu == 6 {
		cariGaji()
	} else {
		ErrorPrint.Println(" Menu Tersebut Tidak Ada ! ")
		menu()
	}
}

func main() {
	menu()
}
