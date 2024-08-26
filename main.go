package main

import (
	"context"
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

	http.HandleFunc("/users", authMiddleware(loggingMiddleware(handleUsers)))
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

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idFromCtx := r.Context().Value("id")
		userID, ok := idFromCtx.(string)
		if !ok {
			log.Printf("Failed to get user ID from context")
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Printf("Request: %s %s by user %s", r.Method, r.URL.Path, userID)
		next(w, r)
	}

}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Здесь можно добавить логику логирования, например:
		userID := r.Header.Get("x-id")
		if userID == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", userID)
		r = r.WithContext(ctx)
		next(w, r)
	}

}

func loggerMiddleware(next interface{}) {

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
