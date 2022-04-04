package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()

	// Kita ingin dapat response dari server tidak lebih dari 2 detik, maka kita set WithTimeout 2 second,
	// ketika lebih dari 2 detik maka client membatalkan request dan juga penanganan request di server dihentikan,
	// jika tanpa context maka proses penanganan request di server masih tetap dijalankan walaupun client membatalkan
	// requestnya.
	ctx, cancel := context.WithTimeout(ctx, 2 * time.Second)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	io.Copy(os.Stdout, res.Body)
}
