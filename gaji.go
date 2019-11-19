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
	Menu int

	itemKaryawan []Karyawan
	itemGaji     []Gaji

	ErrorPrint   = color.New(color.FgRed).Add(color.BgWhite)
	SuccessPrint = color.New(color.FgBlue).Add(color.BgWhite)
	scanner      = bufio.NewScanner(os.Stdin)
)

type Karyawan struct {
	Golongan, Umur, JumlahAnak int
	Nama, Alamat, KodePegawai  string
}

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

func inputKaryawan() {
	var (
		i, Golongan, Umur, JumlahAnak                int
		Nama, Alamat, KodePegawai                    string
		inputLagi                                    string
		validGolongan, validUmur, validAnak, Selesai bool
	)

	for i = len(itemKaryawan); !Selesai; i++ {
		fmt.Println("Masukkan Data Pegawai")

		fmt.Print("Kode Pegawai : \t")
		fmt.Scanln(&KodePegawai)

		fmt.Print("Golongan : \t")
		fmt.Scanln(&Golongan)

		for !validGolongan {
			if Golongan > 3 || Golongan < 0 {
				ErrorPrint.Println("\n Golongan Tersebut Tidak Ada ! \n")
				fmt.Print("Golongan : \t")
				fmt.Scanln(&Golongan)
			} else {
				validGolongan = true
			}
		}

		fmt.Print("Umur Pegawai: \t")
		fmt.Scanln(&Umur)
		for !validUmur {
			if Umur < 1 {
				ErrorPrint.Println("\n Umur Tidak Valid ! \n")
				fmt.Print("Umur Pegawai: \t")
				fmt.Scanln(&Umur)
			} else {
				validUmur = true
			}
		}

		fmt.Print("Nama Pegawai: \t")
		scanner.Scan()
		Nama = scanner.Text()

		fmt.Print("Jumlah Anak : \t")
		fmt.Scanln(&JumlahAnak)
		for !validAnak {
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

		karyawan := Karyawan{Golongan, Umur, JumlahAnak, Nama, Alamat, KodePegawai}
		itemKaryawan = append(itemKaryawan, karyawan)

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

		fmt.Print("\nInput Lagi ? (Y / T) : ")
		fmt.Scanln(&inputLagi)
		fmt.Println("")
		if inputLagi == "T" || inputLagi == "t" {
			Selesai = true
		}
	}
	menu()
}

func prosesCariKaryawan(kode string) int {
	var Selesai bool
	var index int

	for i := 0; i < len(itemKaryawan) && !Selesai; i++ {
		if itemKaryawan[i].KodePegawai == kode {
			index = i
			Selesai = true
		} else {
			index = -1
		}
	}
	return index
}

func cariKaryawan() {
	var KodePegawai string
	var index int

	fmt.Print("Masukkan Kode Pegawai : ")
	fmt.Scanln(&KodePegawai)

	if len(itemKaryawan) == 0 {
		index = -1
	} else {
		index = prosesCariKaryawan(KodePegawai)
	}

	if index != -1 {
		SuccessPrint.Println("\n Data Ditemukan : \n")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Kode Pegawai", "Nama", "Golongan", "Umur", "Jumlah Anak", "Alamat"})
		table.Append(
			[]string{
				itemKaryawan[index].KodePegawai,
				itemKaryawan[index].Nama,
				strconv.Itoa(itemKaryawan[index].Golongan),
				strconv.Itoa(itemKaryawan[index].Umur),
				strconv.Itoa(itemKaryawan[index].JumlahAnak),
				itemKaryawan[index].Alamat,
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

func sortKaryawanGolongan() {
	var (
		sortKaryawan                 []Karyawan
		sorted                       bool
		Golongan, jumlahDataKaryawan int
	)
	fmt.Print("Masukkan Golongan : ")
	fmt.Scanln(&Golongan)

	for i := 0; i < len(itemKaryawan); i++ {
		if itemKaryawan[i].Golongan == Golongan {

			karyawan := Karyawan{
				Nama:        itemKaryawan[i].Nama,
				KodePegawai: itemKaryawan[i].KodePegawai,
				Golongan:    itemKaryawan[i].Golongan,
				Umur:        itemKaryawan[i].Umur,
				JumlahAnak:  itemKaryawan[i].JumlahAnak,
				Alamat:      itemKaryawan[i].Alamat,
			}
			sortKaryawan = append(sortKaryawan, karyawan)
		}
	}

	jumlahDataKaryawan = len(sortKaryawan)
	if jumlahDataKaryawan != 0 {
		for !sorted {
			swapped := false
			for i := 0; i < jumlahDataKaryawan-1; i++ {
				if sortKaryawan[i].Nama > sortKaryawan[i+1].Nama {
					sortKaryawan[i+1], sortKaryawan[i] = sortKaryawan[i], sortKaryawan[i+1]
					swapped = true
				}
			}
			if !swapped {
				sorted = true
			}
			jumlahDataKaryawan--
		}

		SuccessPrint.Println("\n Terdapat", jumlahDataKaryawan+1, "Karyawan Golongan", Golongan, ":\n")

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
func validasiBulan(Bulan string) bool {
	var validBulan bool
	arrBulan := [12]string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli",
		"Agustus", "September", "Oktober", "November", "Desember",
	}

	for j := 0; j < len(arrBulan); j++ {
		if Bulan == arrBulan[j] {
			validBulan = true
		}
	}
	return validBulan
}

func inputGaji() {
	var (
		KodePegawai, Bulan                         string
		JamKerja, TotalGaji, GajiTetap             int
		Selesai, validKode, validBulan, validWaktu bool
		i, Golongan, JumlahAnak                    int
		inputLagi                                  string
	)

	for i = len(itemGaji); !Selesai; i++ {
		fmt.Println("Masukkan Gaji Pegawai")

		fmt.Print("Kode Pegawai \t : ")
		fmt.Scanln(&KodePegawai)
		for !validKode {
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
		for !validBulan {
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

		for !validWaktu {
			if JamKerja >= 0 {
				validWaktu = true
			} else {
				ErrorPrint.Println("\n Jumlah Waktu Kerja Tidak Valid ! \n")
				fmt.Print("Jumlah Jam Kerja : ")
				fmt.Scanln(&JamKerja)
			}
		}

		index := prosesCariKaryawan(KodePegawai)
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

		gaji := Gaji{KodePegawai, Bulan, JamKerja, TotalGaji}
		itemGaji = append(itemGaji, gaji)

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

		fmt.Print("\nInput Lagi ? (Y / T) : ")
		fmt.Scanln(&inputLagi)
		fmt.Println("")

		if inputLagi == "T" || inputLagi == "t" {
			Selesai = true
		}
	}
	menu()
}

func tampilGaji() {
	if len(itemGaji) != 0 {

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

func prosesCariGaji(kode string) int {
	var Selesai bool
	var index int

	for i := 0; i < len(itemGaji) && !Selesai; i++ {
		if itemGaji[i].KodePegawai == kode {
			index = i
			Selesai = true
		} else {
			index = -1
		}
	}
	return index
}

func cariGaji() {
	var KodePegawai string
	var index int

	fmt.Print("Masukkan Kode Pegawai : ")
	fmt.Scanln(&KodePegawai)

	if len(itemGaji) == 0 {
		index = -1
	} else {
		index = prosesCariGaji(KodePegawai)
	}

	if index != -1 {
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

func sortGajiBulan() {
	var (
		sortGaji       []Gaji
		sorted         bool
		Bulan          string
		jumlahDataGaji int
	)
	fmt.Print("Masukkan Bulan Penggajian : ")
	fmt.Scanln(&Bulan)

	for i := 0; i < len(itemGaji); i++ {
		fmt.Println(itemGaji[i].Bulan, " , ", Bulan)
		if itemGaji[i].Bulan == Bulan {

			gaji := Gaji{
				KodePegawai: itemGaji[i].KodePegawai,
				Bulan:       itemGaji[i].Bulan,
				JamKerja:    itemGaji[i].JamKerja,
				TotalGaji:   itemGaji[i].TotalGaji,
			}

			sortGaji = append(sortGaji, gaji)
		}
	}

	jumlahDataGaji = len(sortGaji)
	if jumlahDataGaji != 0 {
		for !sorted {
			swapped := false
			for i := 0; i < jumlahDataGaji-1; i++ {
				if sortGaji[i].KodePegawai < sortGaji[i+1].KodePegawai {
					sortGaji[i+1], sortGaji[i] = sortGaji[i], sortGaji[i+1]
					swapped = true
				}
			}
			if !swapped {
				sorted = true
			}
			jumlahDataGaji--
		}

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
