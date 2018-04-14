/*
imageSystem.go

This file contains the source code which implements the drive image
collection, and the generation of a image file hash.

Version: Apr-02-2018
*/


package main

import (
    "os/exec"
    "os"
    "fmt"
    "log"
    "sync"
)


// Function ImageDrive will follow the path to a device file, and image the
// drive located there.
//
// Param:   drivePath       the path to a device file of a drive
// Param:   imageDirectory  the path to the directory to store the file
// Param:   imageName       the string to name the image file with
func ImageDrive( drivePath string, imageDirectory string, imageName string, wg *sync.WaitGroup ){

    // Create the image file
    var imagePath string = fmt.Sprintf( "%s/%s", imageDirectory,imageName )
    var arg1 string = fmt.Sprintf( "if=%s", drivePath )
    var arg2 string = fmt.Sprintf( "of=%s", imagePath )

    out, err := exec.Command("dd", arg1, arg2 ).Output()

    // Handle any errors that may arise
    if err != nil {
    	wg.Done()
        fmt.Printf("Drive imaging has failed\n")
        log.Fatal(err)
    }

    // Print that imaging was successfull
    fmt.Printf("Drive imaging was sucessfull %s\n", out)
    // No reason to hash memory, changes too often
    if imageName != "memoryDump.dd" {
        // Calculate the hashes
        hashImage( drivePath, imageDirectory )
    }
    defer wg.Done()
}


// Function ImageMem will compile fmem, run the accompaning script and collect a
// memory dump of the system.
//
// Param:   imageDirectory  the path to the directory to store the file
// Param:   driveMount      the path to where the drive has been mounted
func ImageMemory( imageDirectory string, driveMount string, memSize string ){
    // Compile fmem
    cmd := exec.Command( "make" )
    var fmemPath string = fmt.Sprintf( "%s/bin/fmem", driveMount )
    cmd.Dir = fmemPath
    err := cmd.Run()

    // Handle any errors that may arise
    if err != nil {
        fmt.Printf("Building the memory module has failed\n")
        log.Fatal(err)
    }

    // Run the fmem script
    cmd = exec.Command( "./run.sh" )
    cmd.Dir = fmemPath
    err = cmd.Run()
    // Handle any errors that may arise
    if err != nil {
        fmt.Printf("Memory imaging has failed\n")
        log.Fatal(err)
    }

    // Create the image file
    var arg2 string = fmt.Sprintf( "of=%s/memory.dd", imageDirectory )
    var arg3 string = fmt.Sprintf( "count=%s", memSize )

    out, err := exec.Command("dd", "if=/dev/fmem", arg2, "bs=1MB", arg3 ).Output()

    // Handle any errors that may arise
    if err != nil {
        fmt.Printf("Drive imaging has failed\n")
        log.Fatal(err)
    }

    // Print that imaging was successfull
    fmt.Printf("Drive imaging was sucessfull %s\n", out)
}


// Function hashImage will hash the image located at imagePath, and print the
// MD5 and SHA256 hashes to a file.
//
// Param:   imageDirectory  the path to the directory to store the file
// Param:   imagePath       the path to where the image to be hashed is
func hashImage( targetPath string, imageDirectory string ){

    // Create the hashfile
    var hashPath string = fmt.Sprintf( "%s/%s", imageDirectory, "imageHashes.txt" )
    file, err := os.Create( hashPath )
    // Handle any errors that may arise
    if err != nil {
        fmt.Printf("Hashfile creatation has failed.\n")
        log.Fatal(err)
    }

    // Calculate MD5 hash
    file.WriteString( "MD5 Hash\n" )
    file.WriteString( "--------\n" )
    out, err := exec.Command(
        "md5sum", targetPath ).Output()
    // Handle any errors that may arise
    if err != nil {
        fmt.Printf("Calculating the MD5 hash has failed.\n")
        log.Fatal(err)
    }
    // Print the MD5 hash to file
    file.WriteString( string(out) )

    // Calculate SHA256 hash
    file.WriteString( "\nSHA256\n" )
    file.WriteString( "------\n" )
    out, err = exec.Command(
        "shasum", "-a", "256", targetPath ).Output()
    // Handle any errors that may arise
    if err != nil {
        fmt.Printf("Calculating the SHA256 hash has failed.\n")
        log.Fatal(err)
    }
    // Print the MD5 hash to file
    file.WriteString( string(out) )
}
