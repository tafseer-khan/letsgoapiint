package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	// Getting the string from the given link
	res, err := http.Get("https://cloudflare-ipfs.com/ipfs/QmP8jTG1m9GSDJLCbeWhVSVgEzCPPwXRdCRuJtQ5Tz9Kc9")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// Reads the response's body
	resdata, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// uses response as string and prints it
	var content = string(resdata)
	fmt.Println("File content received from link: \n\n" + content)

	// Using the ipfs-go api establishes connection to current daemon, then adds the file
	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(strings.NewReader(content))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s", cid+"\n\n")

	// Using CID from previous function, retreives file from ipfs
	// var id = string(cid)
	out, err := sh.Cat(cid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	// output from previous function is in io.ReadCloser,so is then converted as a string.
	buf := new(bytes.Buffer)
	buf.ReadFrom(out)
	newout := buf.String()
	println("Content from CID " + cid + "\n\n" + newout)
}
