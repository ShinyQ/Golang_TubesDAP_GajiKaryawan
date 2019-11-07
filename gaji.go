package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

/**
	+-------------------------------------+
	|    DECLARE VARIABLE AND RECORD	  |
    +-------------------------------------+
**/

var (
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
	+-------------------------------------+
	|    	MENU AND MAIN FUNCTION    	  |
    +-------------------------------------+
**/

func menu() {
	var Menu int

	fmt.Println("+----------------------------------+--------------------------------------+")
	fmt.Println("| Menu 1 : Input Data Karyawan	   |	 Menu 4 : Input Gaji Karyawan	  |")
	fmt.Println("| Menu 2 : Lihat Data Karyawan	   |	 Menu 5 : Histori Data Gaji	  |")
	fmt.Println("| Menu 3 : Cari Data Karyawan 	   |	 Menu 6 : Cari Histori Data Gaji  |")
	fmt.Println("+----------------------------------+--------------------------------------+\n")

	fmt.Print("Silahakan Pilih Menu : ")
	fmt.Scanln(&Menu)
	fmt.Println("")

}

func main() {
	menu()
}
