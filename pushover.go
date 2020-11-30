package pushover

import (
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

var (
	appTokens []string
	tokenIdx  int
	userKey   string
	enabled   bool
)

func SetupPushover(appTokens_ []string, userKey_ string) {
	appTokens = appTokens_
	userKey = userKey_
	enabled = true
	log.Info("Pushover alerts have been enabled")
}

func SendAlert(priority string, heading string, message string) {
	if enabled {
		data := url.Values{
			"priority": {priority},
			"retry":    {"60"},
			"expire":   {"10"},
			"html":     {"1"},
			"token":    {appTokens[tokenIdx]},
			"user":     {userKey},
			"message":  {"<b>" + heading + "</b>: " + message},
		}

		tokenIdx = (tokenIdx + 1) % len(appTokens)
		_, err := http.PostForm("https://api.pushover.net/1/messages.json", data)

		if err != nil {
			log.Error("Error sending message on Pushover: ", heading, " ", message, " ", err)
		}
	}
}

var (
	NormalPriority    = "0"
	HighPriority      = "1"
	EmergencyPriority = "2"
)
