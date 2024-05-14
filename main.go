package main

import (
	"fmt"
	"sync"

	"github.com/skip2/go-qrcode"
)

func main() {

	url := []string{
		"https://www.github.com/AIDK",
	}

	var wg sync.WaitGroup
	wg.Add(len(url))

	for i, url := range url {
		go func(i int, url string) {
			defer wg.Done()
			generateQR(i, url)
		}(i, url)
	}

	wg.Wait()
}

func generateQR(i int, url string) {

	err := qrcode.WriteFile(url, qrcode.High, 256, fmt.Sprintf("%v.png", i))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("QR code generated and saved as %v.png", i)
}
