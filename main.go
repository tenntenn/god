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

func conf() *oauth.Config {
	var configJSON []byte
	if configJSON, err := ioutil.ReadFile(apiConfig); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
	conf := NewConfigFromJSON(configJSON)
	oauthConf := conf.OAuthConfig()

	return oauthConf
}

func service(conf *oauth.Conf) *drive.Service {
	t := &oauth.Transport{Config: conf}
	client := t.Client()

	var driveService *drive.Service
	if driveService, err := drive.New(oauthHttpClient); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}

	return driveService
}

func main() {
	flag.Parse()

	conf := conf()
	driveService := service(conf)
}
