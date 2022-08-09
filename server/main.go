package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func imgHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	name := header.Filename

	fmt.Println(name)

	// Make images directory if not already there
	err = os.MkdirAll("./images", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./images/%s.jpg", name))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World API!")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", imgHandler).Methods("POST")
	r.HandleFunc("/", getHandler).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":3004", nil))
}
