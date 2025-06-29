package ws

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	logger "github.com/taigrr/log-socket/log"
)

var upgrader = websocket.Upgrader{} // use default options

func SetUpgrader(u websocket.Upgrader) {
	upgrader = u
}

func LogSocketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("upgrade:", err)
		return
	}
	defer c.Close()
	lc := logger.CreateClient()
	defer lc.Destroy()
	lc.SetLogLevel(logger.LTrace)
	logger.Info("Websocket client attached.")
	for {
		logEvent := lc.Get()
		logJSON, _ := json.Marshal(logEvent)
		err = c.WriteMessage(websocket.TextMessage, logJSON)
		if err != nil {
			logger.Warn("write:", err)
			break
		}
	}
}
