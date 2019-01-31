package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Name   string
	Awake  bool
	Hungry bool
}

func main() {
	// writeEmptyFile()
	// writeInFile("Hello World!", "example3.txt")
	// readDir(".")
	// copyFile("./example3.txt")
	c := Config{}
	_, err := toml.DecodeFile("example.toml", &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", c)
}

func copyFile(filetoCopy string) {
	// Read a File from the Disk
	from, err := os.Open(filetoCopy)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()
	// open a file, created if doesn't exist according to flag ( second arg)
	to, err := os.OpenFile("./this_is_a_copy.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}

func readDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Mode(), file.Name(), file.IsDir())
	}
}
func writeInFile(content string, fileName string) {
	fmt.Println([]byte(content)) //[72 101 108 108 111 32 87 111 114 108 100 33]
	err := ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func writeEmptyFile() {
	b := make([]byte, 0)
	//Hello World = [72 101 108 108 111 32 87 111 114 108 100 33]
	err := ioutil.WriteFile("example2.txt", b, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

// ReadFile read a file passed as an argument
func ReadFile(filename string) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileBytes)
	fileString := string(fileBytes)
	fmt.Println(fileString)
}
