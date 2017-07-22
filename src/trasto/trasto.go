package trasto

import (
	"io"
	"log"
	"net/http"
	"os"
)

var (
	Debug bool
)

func Check(err error) error {
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func MakeFile(filename string) *os.File {
	// Create a file on the host file system.
	outfile, err := os.Create(filename)

	Check(err)

	return outfile
}

func DownloadFile(filename string, url string) {
	outfile, err := os.Create(filename)
	Check(err)
	defer outfile.Close()

	if Debug {
		log.Printf("Downloading %s.", url)
	}

	response := MakeRequest(url)

	defer response.Body.Close()

	n, err := io.Copy(outfile, response.Body)
	Check(err)

	if Debug {
		log.Printf("%d bytes written to disk at %s.", n, filename)
	}
}

func MakeRequest(uri string) (r *http.Response) {
	r, err := http.Get(uri)
	Check(err)
	return
}
