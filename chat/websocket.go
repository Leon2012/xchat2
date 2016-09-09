package main

import (
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func ws_init(serverName, addr string) error {
	var (
		bind         string = addr
		listener     *net.TCPListener
		httpServeMux = http.NewServeMux()
		server       *http.Server
		tcpaddr      *net.TCPAddr
		err          error
	)

	httpServeMux.HandleFunc("/chat", ws_handleChat)
	if tcpaddr, err = net.ResolveTCPAddr("tcp4", bind); err != nil {
		return err
	}

	if listener, err = net.ListenTCP("tcp4", tcpaddr); err != nil {
		return err
	}

	server = &http.Server{Handler: httpServeMux}
	go func(host string) {
		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}(bind)

	return nil
}

func ws_handleChat(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		appLogger.Error("Websocket Upgrade error(%v), userAgent(%s)", err, req.UserAgent())
		return
	}
	defer ws.Close()
	sess := globals.sessionStore.Create(ws, "")
	appLogger.Info("%s session connected at %s", sess.sid, sess.remoteAddr)

	go ws_writePump(sess)
	ws_readPump(sess)
}

func ws_writePump(sess *Session) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()

		ws_exit(sess)
	}()
	for {
		select {
		case packet, ok := <-sess.stop:
			if !ok {
				ws_write(sess, websocket.CloseMessage, []byte{})
				return
			}
			if packet != nil {
				ws_write(sess, websocket.TextMessage, packet)
			}
			return
			break
		case packet, ok := <-sess.send:
			if !ok {
				// The hub closed the channel.
				ws_write(sess, websocket.CloseMessage, []byte{})
				return
			}
			sess.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if packet != nil {
				ws_write(sess, websocket.TextMessage, packet)
			}
			break
		case <-ticker.C:
			if err := ws_write(sess, websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func ws_readPump(sess *Session) {
	defer func() {
		sess.ws.Close()
		ws_exit(sess)
	}()
	sess.ws.SetReadLimit(maxMessageSize)
	sess.ws.SetReadDeadline(time.Now().Add(pongWait))
	sess.ws.SetPongHandler(func(string) error { sess.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	sess.remoteAddr = sess.ws.RemoteAddr().String()
	for {
		_, raw, err := sess.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				appLogger.Error("error: %s", err.Error())
			}
			break
		} else {
			sess.dispathRaw(raw)
		}
	}
}

func ws_write(sess *Session, mt int, payload []byte) error {
	sess.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return sess.ws.WriteMessage(mt, payload)
}

func ws_exit(sess *Session) {

}
