package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	IMG_SIZE = 10 << 20 // 10 MB
	IMG_NAME = "file"
)

func main() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			fmt.Println("POST Only")
			return
		}
		fmt.Println("upload start")

		r.ParseMultipartForm(IMG_SIZE)
		file, handler, err := r.FormFile(IMG_NAME)
		defer file.Close()
		if err != nil {
			fmt.Print("Read file failed")
			fmt.Println(err)
			return
		}
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		dst, err := os.Create(handler.Filename)
		defer dst.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Successfully Uploaded File\n")
	})
	http.ListenAndServe(":8888", nil)
}
func init() {
	fmt.Println("server started")
	fmt.Println("using curl -X POST -F 'file=@/your file's dir/xxx.png' http://localhost:8888/upload")
}
