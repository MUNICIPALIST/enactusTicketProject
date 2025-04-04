package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/skip2/go-qrcode"
)

func generateQRCode(w http.ResponseWriter, r *http.Request) {
	code := "https://github.com"

	qrCode, err := qrcode.Encode(code, qrcode.Medium, 256)
	if err != nil {
		log.Println("Error generating QR code:", err)
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")

	w.Write(qrCode)
}

func main() {
	http.HandleFunc("/generate-qrcode", generateQRCode)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
