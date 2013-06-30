package main

import (
	"encoding/json"

	"code.google.com/p/goauth2/oauth"
	"code.google.com/p/google-api-go-client/drive/v2"
)

type Config struct {
	AuthURI                 string   `json:"auth_uri"`
	ClientSecret            string   `json:"client_secret"`
	TokenURI                string   `json:"token_uri"`
	ClientEmail             string   `json:"client_email"`
	RedirectURIs            []string `json:"redirect_uris"`
	ClientX509CertURL       string   `json:"client_x509_cert_url"`
	ClientId                string   `json:"client_id"`
	AuthProviderX509CertURL string   `json:"auth_provider_x509_cert_url"`
}

func NewConfigFromJSON(jsonStr []byte) *Config {

	// including "installed"
	var wrapedConf map[string]interface{}
	json.Unmarshal(jsonStr, &wrapedConf)

	// extract inner conf in "installed"
	jsonStr2, _ := json.Marshal(result["installed"])
	var conf Config
	json.Unmarshal(jsonStr2, &conf)

	return &conf
}

func (c *config) OAuthConfig() *oauth.Config {
	var config = &oauth.Config{
		ClientId:     c.ClientId,
		ClientSecret: c.ClientSecret,
		Scope:        drive.DriveScope,
		AuthURL:      c.AuthURI,
		TokenURL:     c.TokenURI,
	}
}
