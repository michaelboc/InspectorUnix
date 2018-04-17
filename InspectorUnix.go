/*
main.go

This file contains the source code of the main function of the forensic's
tool InspectorUnix

Version: Apr-10-2018
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
    // Memory Dump
    var directory string
    fmt.Println("Where to store memory dump?")
    fmt.Scan(&directory)
    var flashpath string
    fmt.Println( "Where the external drive is mounted?")
    fmt.Scan(&flashpath)
    var memSize string
    fmt.Println( "How many megabytes of memory should be imaged?")
    fmt.Scan(&memSize)
    // Drive imaging
    var pathtodrive string
    fmt.Println("Path to the suspect drive?")
    fmt.Scan(&pathtodrive)
    var wheretoput string
    fmt.Println("Where to store drive image?")
    fmt.Scan(&wheretoput)
    var name string
    fmt.Println("Desired name of image?")
    fmt.Scan(&name)
    ImageMemory( directory, flashpath, memSize )
    makeoutfile()
    fmt.Println("Summary of computer information created, but drive and memory imaging still running")
    go ImageDrive( pathtodrive,wheretoput,name, &wg )
    // Wait till everybody is done
    wg.Wait()
}
