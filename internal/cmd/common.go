package cmd

import (
	v1 "github.com/kyzrfranz/buntesdach/api/v1"
	"github.com/kyzrfranz/interventure-cli/pkg/client"
	"log"
	"os"
)

func EnvOrString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func NoErrorOrExit(err error) {
	if err == nil {
		return
	}
	log.Fatalf("well shit!: %v", err)
}

func FetchPoliticians(url string) []v1.Politician {
	cli, _ := client.NewBuntesdachClient(url)

	list := cli.Politicians().List()

	bios := make([]v1.Politician, 0)
	counter := 0
	for _, politicianListEntry := range list {
		if politicianListEntry.Id.Status == "Aktiv" {
			bio := cli.Politicians().Bio(politicianListEntry.Id.Value)
			bios = append(bios, *bio)
			counter++

			if counter > 3 {
				break
			}
		}
	}

	return bios
}
