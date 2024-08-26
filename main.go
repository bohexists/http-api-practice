package main

import (
	"encoding/json"
	"io/ioutil"
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

	http.HandleFunc("/users", handleUsers)
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
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		addUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func addUsers(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var user User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	users = append(users, user)
}
