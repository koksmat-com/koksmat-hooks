package restapi

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func ImageHook(w http.ResponseWriter, r *http.Request) {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, _ := nats.Connect(url)
	defer nc.Drain()

	id := r.URL.Query().Get("id")
	app := r.URL.Query().Get("app")
	version := r.URL.Query().Get("version")
	listname := r.URL.Query().Get("listname")
	sitename := r.URL.Query().Get("sitename")
	tenantname := r.URL.Query().Get("tenantname")
	subject := app + "." + "hook"
	rep, err := nc.Request(subject, nil, time.Second*10)
	color := "green"
	if err != nil {
		log.Println(subject, "error calling nats", err)
		color = "red"
	} else {
		fmt.Println(string(rep.Data))
	}
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(200)
	svg := `<?xml version="1.0" encoding="UTF-8"?>
<svg width="10" height="10" xmlns="http://www.w3.org/2000/svg">
  <circle cx="5" cy="5" r="4" stroke="black" stroke-width="0" fill="` + color + `" />
</svg>`
	fmt.Fprint(w, svg)

	// fmt.Fprint(w, " id:", id)
	// fmt.Fprint(w, " version: ", version)
	// fmt.Fprint(w, " listname: ", listname)
	// fmt.Fprint(w, " sitename: ", sitename)
	// fmt.Fprint(w, " tenantname: ", tenantname)

	log.Println("ImageHook called", app, id, version, listname, sitename, tenantname)

}
