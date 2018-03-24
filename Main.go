package main

import (
	"os"
	"os/exec"
	"strings"
	"log"
	"fmt"
)

func main(){
	makeoutfile()



}

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
	Outstring = joinstrings(Outstring, runcmdother("find /etc -type f -mtime 1"), "Modified files in /etc in past day")
	fmt.Println("Finding most recently modified files... \n")
	Outstring = joinstrings(Outstring, runcmdother("find ~/.ssh -type f -mtime 1"), "Modified files in ~/.ssh in past day")
	writetofile(Outstring, "out.txt")
	fmt.Println("Done! \n")
	fmt.Println(Outstring)
	fmt.Println("Output also written to file out.txt")
}

func runcmdother(mycmd string) (cmdout string){
	cmdoutbyte, err := exec.Command("sh","-c",mycmd).Output()
	if err!=nil {
		fmt.Printf(mycmd," command failed")
	}
	cmdout = string(cmdoutbyte)
	return
}


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

func writetofile(Outstring string, filename string){
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, Outstring)
}
