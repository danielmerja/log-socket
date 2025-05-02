package browser

import (
	_ "embed"
	"html/template"
	"net/http"
	"strings"
)

//go:embed viewer.html
var webpage string

func LogSocketViewHandler(w http.ResponseWriter, r *http.Request) {
	wsResource := r.Host + r.URL.Path
	if r.TLS != nil {
		wsResource = "wss://" + wsResource
	} else {
		wsResource = "ws://" + wsResource
	}
	wsResource = strings.TrimSuffix(wsResource, "/") + "/ws"
	homeTemplate.Execute(w, wsResource)
}

var homeTemplate = template.Must(template.New("").Parse(webpage))
