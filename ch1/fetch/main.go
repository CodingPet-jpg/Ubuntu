package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		/*
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
				os.Exit(1)
			}
		*/
		//fmt.Printf("%s\n", b)
		fmt.Printf("The status of current resp:%s\n", resp.Status)
		time.Sleep(1000)
		_, error := io.Copy(os.Stdout, resp.Body)
		if error != nil {
			fmt.Fprintf(os.Stderr, "copy:%s\t%v\n", url, error)
		}
		resp.Body.Close()
		fmt.Println()
	}
}
