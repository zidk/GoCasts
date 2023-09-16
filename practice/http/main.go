package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://www.google.es")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	/*
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}(resp.Body)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(body))
	*/

	/*
		bs := make([]byte, 99999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/
	// io.Copy(os.Stdout, resp.Body)

	lw := logWritter{}
	io.Copy(lw, resp.Body)

}

type logWritter struct {
}

func (writter logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
