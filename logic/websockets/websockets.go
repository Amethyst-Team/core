package websockets

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[string]map[*websocket.Conn]bool)
	broadcast = make(chan Message)
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	mu sync.Mutex
)

type Message struct {
	Route string
	Text  string
}

var WS *mux.Router

func init() {
	WS = mux.NewRouter()

	srv := &http.Server{
		Handler: WS,
		Addr:    "localhost:9000",
	}

	go func() {
		log.Printf("Listening on %s\n", "localhost:9000")

		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start WebSocket server: %v", err)
		}
	}()

	go handleMessages()
}

func AddWSRoute(path string, f func(http.ResponseWriter, *http.Request)) {
	WS.HandleFunc(path, f)
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v", err)
		return
	}
	defer ws.Close()

	route := r.URL.Path
	mu.Lock()
	if clients[route] == nil {
		clients[route] = make(map[*websocket.Conn]bool)
	}
	clients[route][ws] = true
	mu.Unlock()

	for {
		// _, message
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			mu.Lock()
			delete(clients[route], ws)
			mu.Unlock()
			break
		}

		//log.Printf("Received message on %s: %s", route, message)
	}
}

func BroadcastMessage(route, msg string) {
	broadcast <- Message{Route: route, Text: msg}
}

func handleMessages() {
	for {
		msg := <-broadcast
		mu.Lock()
		for client := range clients[msg.Route] {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg.Text))
			if err != nil {
				client.Close()
				delete(clients[msg.Route], client)
			}
		}
		mu.Unlock()
	}
}
