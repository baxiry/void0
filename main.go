package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
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
	kinput()

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

func kinput() {

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		if key == keyboard.KeyEsc {
			break
		}
	}
}
