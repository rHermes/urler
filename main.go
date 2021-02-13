package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/url"
	"os"
)

// MUrl is our url type, that mimics the url.URL struct
type MUrl struct {
	*url.URL
	Values url.Values
}

func newMUrl(line string) (*MUrl, error) {
	u, err := url.Parse(line)
	if err != nil {
		return nil, err
	}
	k := &MUrl{URL: u, Values: u.Query()}
	return k, nil
}

func main() {
	scn := bufio.NewScanner(os.Stdin)
	wew := json.NewEncoder(os.Stdout)
	for scn.Scan() {
		m, err := newMUrl(scn.Text())
		if err != nil {
			log.Fatalf("couldn't parse url: %s\n", err.Error())
		}

		if err := wew.Encode(m); err != nil {
			log.Fatalf("couldn't encode murl: %s\n", err.Error())
		}

	}
	if err := scn.Err(); err != nil {
		log.Fatalf("error scanning: %s\n", err.Error())
	}
}
