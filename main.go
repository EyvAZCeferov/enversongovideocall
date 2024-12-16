package main

import (
	"log"
	"net/http"
	"signalling_app/server"
)

func main() {
	server.AllRooms.Init()

	http.HandleFunc("/create", server.CreateRoomRequestHandler)
	http.HandleFunc("/join", server.JoinRoomRequestHandler)

	log.Println("Starting Server on Port 3290")
	err := http.ListenAndServe(":3290", nil)
	if err != nil {
		log.Fatal(err)
	}
}
