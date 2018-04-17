/*
 * main.go
 *
 * This file contains the source code of the main function of the forensic's
 * tool InspectorUnix
 *
 * Version: Apr-17-2018
*/


package main


import (
	"os"
	"os/exec"
	"strings"
	"log"
	"fmt"
)


/**
 * Function collect system information, and outputs it to a file
 */
func makeoutfile(){

	Outstring := "UnixDetective tool output: \n\n"
	fmt.Println("Gathering mounted filesystem info... \n")
	Outstring = joinstrings(Outstring, runcmd("fdisk","-l"), "Mounted Filesystem info:" )
	fmt.Println("Gathering Time and Date info... \n")
	Outstring = joinstrings(Outstring, runcmd("timedatectl", " "),"Time and date at time of analysis:")
	fmt.Println("Gathering User info... \n")
	Outstring = joinstrings(Outstring, runcmd("who", " "),"Logged in users at time of analysis:")
	fmt.Println("Gathering Connection Info... \n")
	Outstring = joinstrings(Outstring, runcmd("netstat","--tcp"), "Connections at time of analsyis:")
	fmt.Println("Gathering Listening Ports Info... \n")
	Outstring = joinstrings(Outstring, runcmd("netstat","-lntu"), "Listening ports at time of analsyis:")
	fmt.Println("Finding most recently modified files... \n")
	Outstring = joinstrings(Outstring, runcmdother("find /etc -type f -mtime 30"), "Modified files in /etc in past month")
	fmt.Println("Finding most recently modified files... \n")
	Outstring = joinstrings(Outstring, runcmdother("find ~/.ssh -type f -mtime 30"), "Modified files in ~/.ssh in past month")
	writetofile(Outstring, "out.txt")
	fmt.Println("Done! \n")
	fmt.Println(Outstring)
	fmt.Println("Output also written to file out.txt")
}


/**
 * Runs the command specified in the cmdout string
 *
 * Param:   mycmd   Command for the function to execute 
 */
func runcmdother(mycmd string) (cmdout string){
	cmdoutbyte, err := exec.Command("sh","-c",mycmd).Output()
	if err!=nil {
		fmt.Printf(mycmd," command failed")
	}
	cmdout = string(cmdoutbyte)
	return
}


/**
 * Runs a command with a supplied arguement string.
 *
 * Param:   cmd     Command for the function to execute
 * Param:   flag    Command arguement to run
 */
func runcmd(cmd string,flag string ) (cmdout string){
	if flag == " " {
		cmdoutbyte, err := exec.Command(cmd).Output()
		if err!=nil {
			fmt.Printf(cmd," command failed")
		}
		cmdout = string(cmdoutbyte)
		return
	}
	cmdoutbyte, err := exec.Command(cmd,flag).Output()
	if err!=nil {
		fmt.Printf(cmd," command failed")
	}
	cmdout = string(cmdoutbyte)
	return
}


/**
 * Function which will join the arguements into a single string
 */
func joinstrings(string1, string2 string, message string) (mashstring string){
	var strs []string
	strs = append(strs, string1)
	strs = append(strs, "\n")
	strs = append(strs, message)
	strs = append(strs, generateline(message))
	strs = append(strs, string2)
	mashstring = strings.Join(strs, "")
	return
}


/**
 * Prepares countstr for printing information
 */
func generateline(counstr string)(line string){
	var strs []string
	strs = append(strs,"\n")
	for i := len(counstr); i > 0 ; i-- {
		strs = append(strs,"-")
	}
	strs = append(strs,"\n")
	line = strings.Join(strs,"")
	return
}


/**
 * Opens the file, and prints Outstring to that file
 */
func writetofile(Outstring string, filename string){
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, Outstring)
}
