/*
main.go

This file contains the source code of the main function of the forensic's
tool InspectorUnix

Version: Mar--2018
*/


package main

import (
    "InspectorUnix/imageDrive"
)

// Main routine of the program
func main() {
    imageDrive.ImageDrive( "/dev/sdb", "/home/michaelboc/", "test.dd" )
}
