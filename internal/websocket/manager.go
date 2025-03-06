package websocket

func NewManager() *Manager {
	return &Manager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (m *Manager) Run() {
	for {
		select {
		case client := <-m.register:
			m.mutex.Lock()
			m.clients[client] = true
			m.mutex.Unlock()

		case client := <-m.unregister:
			if _, ok := m.clients[client]; ok {
				m.mutex.Lock()
				delete(m.clients, client)
				close(client.Send)
				m.mutex.Unlock()
			}

		case message := <-m.broadcast:
			m.mutex.Lock()
			for client := range m.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(m.clients, client)
				}
			}
			m.mutex.Unlock()
		}
	}
}

func (m *Manager) RegisterClient(client *Client) {
	m.register <- client
}

func (m *Manager) Broadcast(message []byte) {
	m.broadcast <- message
}
