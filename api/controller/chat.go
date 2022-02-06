package controller

import (
	"fmt"
	"learning-go/api/helpers"
	"net/http"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,

// We'll need to check the origin of our connection
// this will allow us to make requests from our React
// development server to here.
// For now, we'll do no checking and just allow any connection
// CheckOrigin: func(ctx *http.Request) bool { return true },
// }

// func reader(conn *websocket.Conn) {
// 	for {
// 		// read in a message
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		// Print out the message for clarity
// 		fmt.Println(string(p))
// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		// print out the message for clarity
// 		fmt.Println(string(p))

// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

func ServeWs(pool *helpers.Pool, w http.ResponseWriter, r *http.Request) {
	// ws, err := helpers.Upgrade(c.Writer, c.Request)
	conn, err := helpers.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &helpers.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
	// reader(ws)
	// go helpers.Writer(ws)
	// helpers.Reader(ws)
}
