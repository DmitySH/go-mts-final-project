package notifier

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/entity"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
	"net/http"
	"sync"
)

const (
	driverIDHeader = "user_id"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSNotifier struct {
	clients   map[string]*websocket.Conn
	clientsMu sync.Mutex
}

func NewWSNotifier() *WSNotifier {
	return &WSNotifier{
		clients:   make(map[string]*websocket.Conn),
		clientsMu: sync.Mutex{},
	}
}

func (w *WSNotifier) TripsHandler(rw http.ResponseWriter, r *http.Request) {
	_, span := tracy.Start(r.Context())
	defer span.End()

	driverID := r.Header.Get(driverIDHeader)
	if driverID == "" {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "user_id header is not set")
		return
	}

	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		loggy.Errorln("can't upgrade connection to ws:", err)
		return
	}
	defer conn.Close()

	w.clientsMu.Lock()
	w.clients[driverID] = conn
	w.clientsMu.Unlock()

	defer func() {
		w.clientsMu.Lock()
		delete(w.clients, driverID)
		w.clientsMu.Unlock()
	}()

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			break
		}
		loggy.Infoln(msg)
	}
}

func (w *WSNotifier) NotifyDriver(ctx context.Context, driverID string, offer entity.Offer) error {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	w.clientsMu.Lock()
	driverClient, ok := w.clients[driverID]
	w.clientsMu.Unlock()

	if !ok {
		return nil
	}

	err := driverClient.WriteJSON(offer)
	if err != nil {
		return fmt.Errorf("can't write to client: %w", err)
	}

	return nil
}
