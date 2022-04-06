package main

import (
	"flag"
	"belajar-context/log"
	"fmt"
	"net/http"
	"time"
)

// if you're logging all the things that your program does you might want to have
// request ID to identify in a single server all the things that are related to
// one single request then we could use context value, and we are going to create
// wrapper named Decorate, this is for update context sent by client in order to be
// having ID each request sent by client

func main() {
	flag.Parse()	// <-- untuk mengambil parameter yang dilempar dari command line ketika memanggil aplikasi
	http.HandleFunc("/", log.Decorate(handler))	// <-- this Decorate is just for updating context
	panic(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler ended")

	select {
	case <-time.After(5 * time.Second): // <-- Ini ceritanya penanganan request membutuhkan 5 detik
		fmt.Fprintln(w, "hello")
	case <-ctx.Done(): // <-- Ketika client membatalkan requestnya sintaks ini akan dijalankan dan proses penanganan request dihentikan
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}