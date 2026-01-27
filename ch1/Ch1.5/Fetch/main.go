package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// method 1
	/*
		for _, url := range os.Args[1:] {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
			b, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
				os.Exit(1)
			}
			fmt.Printf("%s", b)
		}
	*/
	// Exercise 1.7: use of io.Copy(dst, src) inplace of ioutils.ReadAll
	// for _, url := range os.Args[1:] {
	// 	resp, err := http.Get(url)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	dst := os.Stdout
	// 	b, err := io.Copy(dst, resp.Body)
	// 	defer resp.Body.Close()
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Printf("%v", b)
	// }

	// Exercise 1.8: using strings.Hasprefix to check for http:// prefixes
	/*
		for _, url := range os.Args[1:] {
			if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
				url = "http://" + url
			}
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
			dst := os.Stdout
			b, err := io.Copy(dst, resp.Body)
			defer resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetchL reading %s: %v\n", url, err)
				os.Exit(1)
			}
			fmt.Printf("%v", b)
		}
	*/

	// Exercise 1.9: modify to print http status code
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Status code: ", resp.Status)
		dst := os.Stdout
		b, err := io.Copy(dst, resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%v", b)
	}
}
