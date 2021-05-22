package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("http://www.gooogle.com/")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Try to investigate how Read function in Reader interface get value
	// bs := make([]byte, 99999)

	// res.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(Dst, Src)
	// Dst: will call Write in Writer interfaces
	// Src: will call Read in Reader interfaces
	// io.Copy(os.Stdout, res.Body)

	log := logWriter{}
	io.Copy(log, res.Body)
}

func (logWriter) Write (bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
