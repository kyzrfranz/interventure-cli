package cmd

import (
	"fmt"
	v1 "github.com/kyzrfranz/buntesdach/api/v1"
	"github.com/kyzrfranz/interventure-cli/pkg/client"
	"log"
	"os"
	"strconv"
)

func EnvOrString(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func EnvOrInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		intVal, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return intVal
	}
	return fallback
}

func NoErrorOrExit(err error) {
	if err == nil {
		return
	}
	log.Fatalf("well shit!: %v", err)
}

func FetchPoliticians(url string, max int) []v1.Politician {
	cli, _ := client.NewBuntesdachClient(url)

	list := cli.Politicians().List()

	bios := make([]v1.Politician, 0)
	counter := 0
	for _, politicianListEntry := range list {
		if politicianListEntry.Id.Status == "Aktiv" {
			fmt.Printf("processing %s %s \n", politicianListEntry.Name, politicianListEntry.Id)
			bio := cli.Politicians().Bio(politicianListEntry.Id.Value)
			bios = append(bios, *bio)
			counter++

			if max > 0 && counter > max {
				break
			}
		}
	}

	return bios
}
