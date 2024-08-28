package main

import (
	"fmt"
	"github.com/bohexists/http-api-practice/coincap"
	"github.com/bohexists/http-api-practice/server"
	"log"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

var (
	users = []User{{1, "Vasya"}, {2, "Petya"}}
)

func main() {

	go server.StartServer()

	coincapClient, err := coincap.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	assets, err := coincapClient.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	for _, asset := range assets {
		fmt.Println(asset.Info())
	}

	bitcoin, err := coincapClient.GetAsset("bitcoin")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bitcoin.Info())

}
