package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/allantabilog/visualiser/internal/visualizer"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Server represents the web server for the visualizer
type Server struct {
	router     *mux.Router
	visualizer *visualizer.Visualizer
	upgrader   websocket.Upgrader
	clients    map[*websocket.Conn]bool
	broadcast  chan []byte
}

// NewServer creates a new server instance
func NewServer(viz *visualizer.Visualizer) *Server {
	server := &Server{
		router:     mux.NewRouter(),
		visualizer: viz,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Allow all origins for development
				return true
			},
		},
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
	}
	
	server.setupRoutes()
	server.startBroadcaster()
	
	return server
}

// setupRoutes configures the HTTP routes
func (s *Server) setupRoutes() {
	// Serve static files
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", 
		http.FileServer(http.Dir("./web/static/"))))
	
	// API routes
	api := s.router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/snapshots", s.handleGetSnapshots).Methods("GET")
	api.HandleFunc("/snapshots/{id}", s.handleGetSnapshot).Methods("GET")
	api.HandleFunc("/operations", s.handleGetOperations).Methods("GET")
	api.HandleFunc("/operations/{id}", s.handleGetDataStructureOperations).Methods("GET")
	
	// WebSocket endpoint
	s.router.HandleFunc("/ws", s.handleWebSocket)
	
	// Main page
	s.router.HandleFunc("/", s.handleIndex).Methods("GET")
}

// startBroadcaster starts the message broadcaster for WebSocket clients
func (s *Server) startBroadcaster() {
	go func() {
		for {
			message := <-s.broadcast
			for client := range s.clients {
				err := client.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Printf("WebSocket error: %v", err)
					client.Close()
					delete(s.clients, client)
				}
			}
		}
	}()
	
	// Subscribe to visualizer operations
	go func() {
		operationChan := s.visualizer.Subscribe()
		for operation := range operationChan {
			// Wrap in a message envelope
			envelope := map[string]interface{}{
				"type": "operation",
				"data": operation,
			}
			
			envelopeJSON, err := json.Marshal(envelope)
			if err != nil {
				log.Printf("Error creating message envelope: %v", err)
				continue
			}
			
			s.broadcast <- envelopeJSON
		}
	}()
}

// handleIndex serves the main HTML page
func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/templates/index.html")
}

// handleWebSocket handles WebSocket connections
func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()
	
	// Register client
	s.clients[conn] = true
	log.Printf("New WebSocket client connected. Total clients: %d", len(s.clients))
	
	// Send current state to new client
	snapshots, err := s.visualizer.GetAllSnapshots()
	if err == nil {
		envelope := map[string]interface{}{
			"type": "initial_state",
			"data": snapshots,
		}
		
		message, err := json.Marshal(envelope)
		if err == nil {
			conn.WriteMessage(websocket.TextMessage, message)
		}
	}
	
	// Listen for messages from client
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			delete(s.clients, conn)
			break
		}
		
		// Handle client messages (e.g., requests for specific snapshots)
		s.handleWebSocketMessage(conn, message)
	}
}

// handleWebSocketMessage processes messages from WebSocket clients
func (s *Server) handleWebSocketMessage(conn *websocket.Conn, message []byte) {
	var request map[string]interface{}
	if err := json.Unmarshal(message, &request); err != nil {
		log.Printf("Error parsing WebSocket message: %v", err)
		return
	}
	
	messageType, ok := request["type"].(string)
	if !ok {
		log.Printf("Invalid message type in WebSocket message")
		return
	}
	
	switch messageType {
	case "get_snapshot":
		if id, ok := request["id"].(string); ok {
			snapshot, err := s.visualizer.GetSnapshot(id)
			if err != nil {
				log.Printf("Error getting snapshot: %v", err)
				return
			}
			
			response := map[string]interface{}{
				"type": "snapshot",
				"data": snapshot,
			}
			
			responseJSON, err := json.Marshal(response)
			if err != nil {
				log.Printf("Error serializing snapshot response: %v", err)
				return
			}
			
			conn.WriteMessage(websocket.TextMessage, responseJSON)
		}
	}
}

// handleGetSnapshots returns all current snapshots
func (s *Server) handleGetSnapshots(w http.ResponseWriter, r *http.Request) {
	snapshots, err := s.visualizer.GetAllSnapshots()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snapshots)
}

// handleGetSnapshot returns a specific snapshot by ID
func (s *Server) handleGetSnapshot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	snapshot, err := s.visualizer.GetSnapshot(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snapshot)
}

// handleGetOperations returns all operations
func (s *Server) handleGetOperations(w http.ResponseWriter, r *http.Request) {
	operations := s.visualizer.GetOperations()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(operations)
}

// handleGetDataStructureOperations returns operations for a specific data structure
func (s *Server) handleGetDataStructureOperations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	operations := s.visualizer.GetOperationsForDataStructure(id)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(operations)
}

// Start starts the HTTP server
func (s *Server) Start(port int) error {
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server on %s", addr)
	return http.ListenAndServe(addr, s.router)
}