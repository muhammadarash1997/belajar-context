package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Printf("handler started")
	defer log.Printf("handler ended")

	select {
	case <-time.After(5 * time.Second):	// <-- Ini ceritanya penanganan request membutuhkan 5 detik
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():	// <-- Ketika client membatalkan requestnya sintaks ini akan dijalankan dan proses penanganan request dihentikan
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
