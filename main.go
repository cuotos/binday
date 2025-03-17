package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/cuotos/binday/bin"
)

const (
	urlFormat = "https://online.bcpcouncil.gov.uk/bcp-apis/?api=BinDayLookup&uprn=%s"
)

var uprn string

func main() {

	flag.StringVar(&uprn, "uprn", "", "UPRN of the property to check. use -uprn or UPRN envvar")
	flag.Parse()

	if uprn == "" && os.Getenv("UPRN") != "" {
		uprn = os.Getenv("UPRN")
	}

	if uprn == "" {
		flag.Usage()
		os.Exit(1)
	}

	url := fmt.Sprintf(urlFormat, uprn)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)

	bins, err := run(respBody)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bins.Next())
}

func run(input []byte) (bin.Bins, error) {

	bins := bin.Bins{}

	err := json.Unmarshal(input, &bins)
	if err != nil {
		return nil, err
	}

	return bins, nil
}
