/*
main.go

This file contains the source code of the main function of the forensic's
tool InspectorUnix

Version: Mar--2018
*/


package main

import (
	"fmt"
	"sync"
)

// Main routine of the program
func main() {
	var wg sync.WaitGroup
	var pathtodrive string
	fmt.Println("Input path to drive to image")
	fmt.Scan(&pathtodrive)
	var wheretoput string
	fmt.Println("Input where image should go")
	fmt.Scan(&wheretoput)
	var name string
	fmt.Println("Input desired name of image")
	fmt.Scan(&name)
	wg.Add()
    go ImageDrive(pathtodrive,wheretoput,name, &wg )
	var directory string
	fmt.Println("Input directory to put on flash drive")
	fmt.Scan(&directory)
	var flashpath string
	fmt.Println("Input flash drive mounted path")
	fmt.Scan(&flashpath)
	wg.Add()
    go ImageMemory( directory, flashpath, &wg)
    makeoutfile()
    fmt.Println("Summary of computer information created, but drive and memory imaging still running")

}
