package main

import (
	"fmt"
	"github.com/bohexists/http-api-practice/client/ipstack"
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

	//Coincap
	//coincapClient, err := coincap.NewClient(time.Second * 10)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//IPstack
	ipstackClient, err := ipstack.NewClient("aeccdd0513bae8c553250e5761328461", time.Second*10)
	if err != nil {
		log.Fatal(err)
	}

	//Coincap
	//assets, err := coincapClient.GetAssets()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, asset := range assets {
	//	fmt.Println(asset.Info())
	//}

	//bitcoin, err := coincapClient.GetAsset("bitcoin")
	//if err != nil {
	//	log.Fatal(err)
	//}

	// IPstack
	ipInfo, err := ipstackClient.GetIPInfo("134.201.250.155")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IP: %s\nCountry: %s\nCity: %s\n", ipInfo.IP, ipInfo.CountryName, ipInfo.City)

	//fmt.Println(bitcoin.Info())
}
