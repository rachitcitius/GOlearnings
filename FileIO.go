// FileIO
package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	//"bufio"
	"os"
)

const filePath string="File.txt"

func main() {
	// ReadFile()
	WriteFile()
}

func chkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile() {
	//Simply reading a file & printing the content in it
	filedata, err := ioutil.ReadFile(filePath)
	chkErr(err)
	fmt.Print(string(filedata))
	
	//Opening a file to do certain operations
	f, err := os.Open(filePath)
	chkErr(err)
	
	//Reading 5 bytes from beginning
	b1 := make([]byte, 5)
	bf1, err := f.Read(b1)
	chkErr(err)
	fmt.Printf("\n %d bytes: %s \n", bf1, string(b1))
	
	//Seek to a position & read from there
	s1, err := f.Seek(8,0)
	chkErr(err)
	b2 := make([]byte, 4)
	bf2, err := f.Read(b2)
	chkErr(err)
	fmt.Printf("\n %d bytes starting from %d: %s \n", bf2, s1, string(b2))
	
	
	defer f.Close()
}

func WriteFile() {
	//Basic way to write to a file
	d1 := []byte("Hello from Go.\n It is really easy to write to a file!\n")
	err := ioutil.WriteFile("CreatedFromGo.txt", d1, 0644)
	chkErr(err)
	
	//Create file & write with more control
	fl, err := os.Create("CreatedFromGoNew.txt")
	chkErr(err)
	
	defer fl.Close()
	
	data := []byte{65, 112, 13, 97, 13}
	n2, err := fl.Write(data)
	chkErr(err)
	
	fmt.Printf("Wrote %d bytes to file", n2)
	
	fl.WriteString("\nAdded another line!")
	fl.Sync()
}