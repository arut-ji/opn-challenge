package main

import (
	"fmt"
	"io"
	"opn-challenge/cipher"
	"os"
)

func main() {
	f, err := os.Open("./data/fng.1000.csv.rot128")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	rot128Reader, err := cipher.NewRot128Reader(f)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)

	decryptedResult := ""

	for {
		n, err := rot128Reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			decryptedResult += string(buf[:n])
		} else {
			break
		}
	}

	err = os.WriteFile("./data/fng.1000.csv", []byte(decryptedResult), 0644)
	if err != nil {
		panic(err)
	}
}
