package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	"code.google.com/p/goauth2/oauth"
	"code.google.com/p/google-api-go-client/drive/v2"
)

var (
	apiConfig = flag.string("c", "client_secret.json", "client secret JSON file")
)

func main() {
	flag.Parse()

	var configJSON []byte
	if configJSON, err := ioutil.ReadFile(apiConfig); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
	}
	conf := NewConfigFromJSON(configJSON)
}
