package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func imgHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved request!")

	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("file")

	if err != nil {
		//fmt.Printf("error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	fn := header.Filename

	fmt.Println(fn)

	tempFile, err := ioutil.TempFile("images", fn+"-*.jpg")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tempFile.Write(fileBytes)

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
