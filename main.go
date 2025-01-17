package main

import (
	"flag"
	"fmt"
	"github.com/kyzrfranz/interventure-cli/pkg/client"
	"log"
	"os"
)

var (
	apiUrl string
)

func main() {

	flag.StringVar(&apiUrl, "api-url", envOrString("API_URL", "http://localhost:8080"), "buntesdach API URL")

	cli, _ := client.NewBuntesdachClient(apiUrl)

	//list := cli.Committees().List()
	detail := cli.Committees().Detail("a11")
	//fmt.Printf("list: %v\n", list)
	fmt.Printf("detail: %v\n", detail)

}

func envOrString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func noErrorOrExit(err error) {
	if err == nil {
		return
	}
	log.Fatalf("error: %v", err)
}
