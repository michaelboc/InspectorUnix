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
	wg.Add(1)
	var pathtodrive string
	fmt.Println("Path to the suspect drive:")
	fmt.Scan(&pathtodrive)
	var wheretoput string
	fmt.Println("Where to store drive image:")
	fmt.Scan(&wheretoput)
	var name string
	fmt.Println("Desired name of image:")
	fmt.Scan(&name)
        go ImageDrive(pathtodrive,wheretoput,name, &wg )
	var directory string
	fmt.Println("Where to store memory dump:")
	fmt.Scan(&directory)
	var flashpath string
	fmt.Println("Path to the external drive mount point:")
	fmt.Scan(&flashpath)
        makeoutfile()
        fmt.Println("Summary of computer information created, but drive and memory imaging still running")
        wg.Wait() 
        ImageMemory( directory, flashpath )
}
