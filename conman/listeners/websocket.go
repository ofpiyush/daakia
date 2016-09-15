package listeners

import (
	"golang.org/x/net/websocket"
	"log"
	"net"
	"net/http"
)

type WebSocketListener struct {
	Addr   string
	mx     *http.ServeMux
	Logger *log.Logger
}

func NewWebSocket() *WebSocketListener {
	return &WebSocketListener{
		mx: http.NewServeMux(),
	}
}
func (w *WebSocketListener) Handle(pattern string, h func(net.Conn) error) {
	w.mx.Handle(pattern, websocket.Handler(func(ws *websocket.Conn) {
		h(ws)
	}))
}

func (w *WebSocketListener) Listen(addr string) error {
	return http.ListenAndServe(addr, w.mx)
}
