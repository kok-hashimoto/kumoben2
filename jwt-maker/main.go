package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jws"
)

func main() {
	cmd := &cobra.Command{
		Use:  "jwt-token mail-address service-account-key.json",
		Args: cobra.ExactArgs(2),
		Run:  run,
	}
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// https://cloud.google.com/api-gateway/docs/authenticate-service-account?hl=ja#making_an_authenticated_request

func run(cmc *cobra.Command, args []string) {
	serviceAccountEmail := args[0]
	serviceAccountKeyFile := args[1]
	now := time.Now().Unix()
	var expiryLength int64 = 3600
	jwt := &jws.ClaimSet{
		Iat:           now,
		Exp:           now + expiryLength,
		Iss:           serviceAccountEmail,
		Aud:           "audience",
		Sub:           serviceAccountEmail,
		PrivateClaims: map[string]interface{}{"email": serviceAccountEmail},
	}
	jwsHeader := &jws.Header{
		Algorithm: "RS256",
		Typ:       "JWT",
	}

	// Extract the RSA private key from the service account keyfile.
	sa, err := ioutil.ReadFile(serviceAccountKeyFile)
	if err != nil {
		log.Fatalf("failed to read service account key file: %s: %s", serviceAccountKeyFile, err.Error())
	}
	conf, err := google.JWTConfigFromJSON(sa)
	if err != nil {
		log.Fatalf("Could not parse service account JSON: %s", err.Error())
	}
	block, _ := pem.Decode(conf.PrivateKey)
	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("private key parse error: %s", err.Error())
	}
	rsaKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		log.Fatal("private key failed rsa.PrivateKey type assertion")
	}
	encoded, err := jws.Encode(jwsHeader, jwt, rsaKey)
	if err != nil {
		log.Fatalf("failed to encode JWT: %s", err.Error())
	}
	fmt.Println(encoded)
}
