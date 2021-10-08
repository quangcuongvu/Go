package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {

	//check file (file<<2MB)
	r.Body = http.MaxBytesReader(w, r.Body, 2 * 1024 * 1024) // 2 Mb
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()


	// Check info file
	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)

	
	// Upload file
	resFile, err := os.Create("./Upload/" + handler.Filename)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	defer resFile.Close()
	if err == nil {
		io.Copy(resFile, file)
		defer resFile.Close()
		fmt.Fprintf(w, "Successfully Uploaded Original File\n")
	}
}

//start server which will listen and server on pot 8083
func startServer() {

	http.HandleFunc("/upload", uploadFile)

	//server is listening on port 8080
	http.ListenAndServe(":8080", nil)
}

func main() {

	fmt.Println("Start run localhost port 8080. With API: localhost:8080/upload")
	
	// Create folder Upload not exist
	_, err := os.Stat("Upload")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("Upload", 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}
	startServer()
}
