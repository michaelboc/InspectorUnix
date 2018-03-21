/*
imageDrive.go

This file contains the source code which implements the drive image
collection, and the generation of a image file hash.

Version: Mar-20-2018
*/


package imageDrive

import (
    "os/exec"
    "os"
    "fmt"
    "log"
)


// Function ImageDrive will follow the path to a device file, and image the
// drive located there.
//
// Param:   drivePath       the path to a device file of a drive
func ImageDrive( drivePath string, imageDirectory string, imageName string ){

    // Create the image file
    var imagePath string = fmt.Sprintf( "%s%s", imageDirectory, imageName )
    var arg1 string = fmt.Sprintf( "if=%s", drivePath )
    var arg2 string = fmt.Sprintf( "of=%s", imagePath )
    out, err := exec.Command("dd", arg1, arg2 ).Output()

    // Handle any errors that may arise
    if err != nil {
        fmt.Printf("Drive imaging has failed\n")
        log.Fatal(err)
    }

    // Print that imaging was successfull
    fmt.Printf("Drive imaging was sucessfull %s\n", out)
    // Calculate the hashes
    hashImage( imageDirectory, imagePath )
}


// Function hashImage will hash the image located at imagePath, and print the
// MD5 and SHA256 hashes to a file.
func hashImage( imageDirectory string, imagePath string ){
    
    // Create the hashfile
    var hashPath string = fmt.Sprintf( "%s%s", imageDirectory, "imageHashes.txt" ) 
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
        "md5sum", imagePath ).Output()
    // Handle any errors that may arise
    if err != nil {
        fmt.Printf("Calculating the MD5 hash has failed.\n")
        log.Fatal(err)
    }
    // Print the MD5 hash to file    
    file.WriteString( string(out) ) 

    // Calculate SHA256 hash 
    file.WriteString( "\nSHA1\n" ) 
    file.WriteString( "----\n" ) 
    out, err = exec.Command(
        "sha1sum", imagePath ).Output()
    // Handle any errors that may arise
    if err != nil {
        fmt.Printf("Calculating the SHA256 hash has failed.\n")
        log.Fatal(err)
    }
    // Print the MD5 hash to file    
    file.WriteString( string(out) ) 
}
