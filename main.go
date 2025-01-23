package main

import (
	"flag"
	"fmt"
	v1 "github.com/kyzrfranz/buntesdach/api/v1"
	"github.com/kyzrfranz/interventure-cli/pkg/client"
	"github.com/kyzrfranz/interventure-cli/pkg/xls"
	"log"
	"os"
)

var (
	apiUrl string
)

func main() {

	flag.StringVar(&apiUrl, "api-url", envOrString("API_URL", "http://localhost:8080"), "buntesdach API URL")

	cli, _ := client.NewBuntesdachClient(apiUrl)

	list := cli.Politicians().List()

	bios := make([]v1.Politician, 0)
	counter := 0
	for _, politicianListEntry := range list {
		if politicianListEntry.Id.Status == "Aktiv" {
			bio := cli.Politicians().Bio(politicianListEntry.Id.Value)
			fmt.Printf("#%d %s, %s, %v\n", counter, politicianListEntry.Name, politicianListEntry.Id.Value)
			bios = append(bios, *bio)
			counter++
			//if counter > 500 {
			//	break
			//}
		}
	}

	xlsGenerator := xls.NewGenerator()
	err := xlsGenerator.Generate("./adds.xlsx", bios)
	noErrorOrExit(err)

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
