package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

var dbName string = "test"

func main() {
	socketUrl := "ws://localhost:1111/ws"

	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		fmt.Println("Error! :", err)
		return
	}
	defer conn.Close()

	// handle std input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(dbName + "> ")
		inputQuery, _ := reader.ReadString('\n')
		if inputQuery == "" {
			println("zeroval")
			continue
		}

		// Send an echo packet every second
		err = conn.WriteMessage(websocket.TextMessage, []byte(inputQuery))
		if err != nil {
			fmt.Println("Error! :", err)
			return
		}

		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error in receive:", err)
			return
		}
		fmt.Println(string(msg))

	}
}
