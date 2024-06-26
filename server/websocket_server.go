package server

import (
	"net/http"
	"strconv"

	"github.com/craterone/go-protoo/logger"
	"github.com/craterone/go-protoo/transport"
	"github.com/gorilla/websocket"
)

type WebSocketServerConfig struct {
	Host          string
	Port          int
	CertFile      string
	KeyFile       string
	WebSocketPath string
}

func DefaultConfig() WebSocketServerConfig {
	return WebSocketServerConfig{
		Host:          "0.0.0.0",
		Port:          8443,
		WebSocketPath: "/ws",
	}
}

type WebSocketHandler interface {
	ServerWs(ws *transport.WebSocketTransport, request *http.Request)
}

type WebSocketServer struct {
	handleWebSocket WebSocketHandler
	// Websocket upgrader
	upgrader websocket.Upgrader
}

func NewWebSocketServer(handler WebSocketHandler) *WebSocketServer {
	var server = &WebSocketServer{
		handleWebSocket: handler,
	}
	server.upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return server
}

func (server *WebSocketServer) handleWebSocketRequest(writer http.ResponseWriter, request *http.Request) {
	responseHeader := http.Header{}
	responseHeader.Add("Sec-WebSocket-Protocol", "protoo")
	socket, err := server.upgrader.Upgrade(writer, request, responseHeader)
	if err != nil {
		panic(err)
	}
	wsTransport := transport.NewWebSocketTransport(socket)
	wsTransport.Start()

	server.handleWebSocket.ServerWs(wsTransport, request)
}

func (server *WebSocketServer) Bind(cfg WebSocketServerConfig) {
	// Websocket handle func
	http.HandleFunc(cfg.WebSocketPath, server.handleWebSocketRequest)

	if cfg.CertFile == "" || cfg.KeyFile == "" {
		logger.Infof("non-TLS WebSocketServer listening on: %s:%d", cfg.Host, cfg.Port)
		panic(http.ListenAndServe(cfg.Host+":"+strconv.Itoa(cfg.Port), nil))
	} else {
		logger.Infof("TLS WebSocketServer listening on: %s:%d", cfg.Host, cfg.Port)
		panic(http.ListenAndServeTLS(cfg.Host+":"+strconv.Itoa(cfg.Port), cfg.CertFile, cfg.KeyFile, nil))
	}
}
