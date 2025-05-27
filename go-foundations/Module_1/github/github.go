package main

// Given a github user login, return name and number of public repos

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main () {
	resp, err := http.Get("http://api.github.com/users/ardanlabs");
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: bad status - %s\n", resp.Status)
		return
	}

	ctype := resp.Header.Get("Content-Type")
	fmt.Println("content-type: ", ctype)

	io.Copy(os.Stdout, resp.Body)
}
