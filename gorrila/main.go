package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	fmt.Println("Start ")

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		fmt.Println("Connected: ")

		fmt.Println("Enter text: ")

		for {
			var input string
			fmt.Scanln(&input)
			msg := []byte("Server " + input)
			// if input == "" {
			// 	_, msg, err := conn.ReadMessage()
			// 	if err != nil {
			// 		return
			// 	}
			// 	msg = []byte("Client " + string(msg))

			// }
			fmt.Printf("%s\n", string(msg))
			// Write message back to browser
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}
		}

		// for {
		// 	// Read message from browser
		// 	msgType, msg, err := conn.ReadMessage()
		// 	if err != nil {
		// 		return
		// 	}

		// 	// Print the message to the console
		// 	fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// }
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}
