package main

import (
	"os"
	. "github.com/Coosis/cos-goflv"
)

func main() {
	// Reading an flv
	// Place the sample.flv file in root dir
	f, err := os.Open("sample.flv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	flv := Flv{}
	if err := flv.Read(f); err != nil {
		panic(err)
	}

	// Printing info of the flv
	// Go to print_data.go to see the printflv function
	printflv(&flv)

	// Writing an flv
	// Will create a copy of the sample.flv file in root dir
	f2, err := os.Create("sample_cp.flv")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	if err := flv.Write(f2); err != nil {
		panic(err)
	}
}
