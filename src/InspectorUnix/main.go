/*
main.go

This file contains the source code of the main function of the forensic's
tool InspectorUnix

Version: Mar--2018
*/


package main

import (
    "InspectorUnix/imageSystem"
    "os"
)

// Main routine of the program
func main() {
    imageSystem.ImageMemory( os.Args[1], os.Args[2] )
}
