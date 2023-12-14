package notifier

import (
	"github.com/gorilla/websocket"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/service"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSNotifier struct {
	driverService *service.DriverService
}

func NewWSNotifier(driverService *service.DriverService) *WSNotifier {
	return &WSNotifier{driverService: driverService}
}

func (w *WSNotifier) TripsHandler(rw http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		loggy.Warnln("can't upgrade connection to ws:", err)
		return
	}
	defer connection.Close()

	for {
		mt, msg, err := connection.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			break
		}

		loggy.Infoln(string(msg))

	}
}
