package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func add(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 100 MB files
	r.ParseMultipartForm(100 << 20)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("File")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	dst, err := os.Create(fmt.Sprintf("%s/%s", "images", handler.Filename))
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func remove(w http.ResponseWriter, r *http.Request) {
	var filelist []string
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(body, &filelist)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(filelist)
	for _, file := range filelist {
		err := os.Remove(file)
		if err != nil {
			fmt.Println(err, file)
		}
	}

}

func run_feh() {
	for {
		cmd := exec.Command("feh", "-Z", "-x", "-F", "-Y", "-q", "-D", "10", "-R", "60", "images/")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Print(err)
			return
		}
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	var files = []string{}

	root := "images/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path != root {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	bytes, _ := json.Marshal(files)
	fmt.Println(string(bytes))
	w.Write(bytes)
}

func main() {
	fs := http.FileServer(http.Dir("./pages/build"))
	http.Handle("/", fs)
	http.HandleFunc("/add", add)
	http.HandleFunc("/remove", remove)
	http.HandleFunc("/list", list)

	go run_feh()
	http.ListenAndServe(":8080", nil)
}
