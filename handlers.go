package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	tmpl, _ := template.ParseFiles("templates/index.html")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


type ConnectUser struct {
	Websocket *websocket.Conn
	ClientIP string
}

func newConnectUser(ws *websocket.Conn, clientIP string) *ConnectUser  {
	return &ConnectUser{
		Websocket: ws,
		ClientIP: clientIP,
	}
}

var users = make(map[ConnectUser]int)

func WebsocketHandler(w http.ResponseWriter, r *http.Request)  {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	
	ws, _ := upgrader.Upgrade(w, r, nil)

	defer func() {
		if err := ws.Close(); err != nil {
			log.Println("Websocket could not be closed", err.Error())
		}
	}()

	log.Println("Client connected:", ws.RemoteAddr().String())
	var socketClient *ConnectUser = newConnectUser(ws, ws.RemoteAddr().String())
	users[*socketClient] = 0
	log.Println("Number client connected ...", len(users))

	for {
		messageType, message, err := ws.ReadMessage()
		if  err != nil {
			log.Println("Ws disconnect waiting", err.Error())
			delete(users, *socketClient)
			log.Println("Number of client still connected ...", len(users))
			return
		}

		user := Messages{Message: string(message)}
		DB.Create(&user)


		for client := range users {
			if err = client.Websocket.WriteMessage(messageType, message); err != nil {
				log.Println("Cloud not send Message to ", client.ClientIP, err.Error())
			}
		}

	}
}