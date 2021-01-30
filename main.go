package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type UserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email_address"`
}

func hello() {
	fmt.Println("hello!")
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(handle))
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)

	if r.URL.Path == "/" || r.URL.Path == "/index.html" {
		indexBytes, _ := ioutil.ReadFile("./index.html")
		w.Write(indexBytes)
	} else if r.URL.Path == "/connect" {
		upgrader := websocket.Upgrader{}
		ws, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			fmt.Println(err)
		}

		jsonUserInfo, _ := json.MarshalIndent(UserInfo{
			ID:    "ergerkgberg",
			Email: "bla@email.com",
		}, "", "    ")

		ws.WriteMessage(websocket.TextMessage, jsonUserInfo)

		for {
			messageType, data, err := ws.ReadMessage()
			if err != nil {
				log.Printf("error reading message from websocket: %v\n", err)
				break
			}

			fmt.Printf("message received: %d %s\n", messageType, string(data))
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
