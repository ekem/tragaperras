package trasto

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Make_a_file(filename string) *os.File {
	// Create a file on the host file system.
	outfile, err := os.Create(filename)

	Check(err)

	return outfile
}

func Get_a_file(filename string, url string) {
	log.Printf("Downloading %s.", url)
	outfile, err := os.Create(filename)
	Check(err)
	defer outfile.Close()

	response := Make_a_request(url)

	defer response.Body.Close()
	n, err := io.Copy(outfile, response.Body)
	Check(err)

	log.Printf("%d bytes written to disk at %s.", n, filename)
}

func Make_a_request(uri string) (r *http.Response) {
	r, err := http.Get(uri)
	Check(err)
	return
}
