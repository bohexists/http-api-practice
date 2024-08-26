package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

var (
	users = []User{{1, "Vasya"}, {2, "Petya"}}
)

func main() {
	//coincapClient, err := coincap.NewClient(time.Second * 10)
	//if err != nil {
	//	log.Fatal(err)
	//}

	http.HandleFunc("/user", handleUsers)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

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
	//
	//fmt.Println(bitcoin.Info())

}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		w.Write(resp)
	}
}
