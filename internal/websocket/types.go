package websocket

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Client struct {
	ID     uint
	Conn   *websocket.Conn
	Send   chan []byte
	UserID uint
}

type Manager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mutex      sync.Mutex
}
