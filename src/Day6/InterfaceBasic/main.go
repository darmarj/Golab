package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Speaker interface {
	Say(string)
}

func SpeakAlphabet(speaker Speaker) {
	speaker.Say("abcdefghijk...xyz")
}

type Person struct {
	name string
}

func (p *Person) Say(message string) {
	fmt.Println(p.name, ":", message)
}

type SpeakWriter struct {
	w io.Writer
}

func (sw *SpeakWriter) Say(message string) {
	io.WriteString(sw.w, message)
}

type FileWriter struct {
	filename string
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	err = ioutil.WriteFile(fw.filename, p, 0644)
	n = 0
	return
}

func main() {
	people01 := new(Person)
	people01.name = "James"
	SpeakAlphabet(people01)

	fmt.Println("===========")
	console := new(SpeakWriter)
	console.w = os.Stdout
	SpeakAlphabet(console)

	fmt.Println("===========")
	fileWriter := new(FileWriter)
	fileWriter.filename = "GoBuffet/src/Day6/Interface/abc.txt"

	fileSpeakWriter := new(SpeakWriter)
	fileSpeakWriter.w = fileWriter
	SpeakAlphabet(fileSpeakWriter)
}
