package main

import (
	"bytes"
	"image/png"
	"log"
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func generateBarcode(url string) ([]byte, error) {
	// QR kod oluştur
	qrCode, err := qr.Encode(url, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	// QR kod resmini oluştur
	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		return nil, err
	}

	// Resmi PNG formatında encode et
	var buf bytes.Buffer
	if err := png.Encode(&buf, qrCode); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func barcodeHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL parametresi gereklidir", http.StatusBadRequest)
		return
	}

	barcodeImg, err := generateBarcode(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Resmi yanıt olarak gönder
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", string(len(barcodeImg)))
	if _, err := w.Write(barcodeImg); err != nil {
		log.Println("Resim yazma hatası:", err)
	}
}

func main() {
	http.HandleFunc("/barcode", barcodeHandler)

	log.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
