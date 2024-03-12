package restapi

import (
	"fmt"
	"log"
	"net/http"
)

func ImageHook(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	version := r.URL.Query().Get("version")
	listname := r.URL.Query().Get("listname")
	sitename := r.URL.Query().Get("sitename")
	tenantname := r.URL.Query().Get("tenantname")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(200)
	fmt.Fprint(w, " id:", id)
	fmt.Fprint(w, " version: ", version)
	fmt.Fprint(w, " listname: ", listname)
	fmt.Fprint(w, " sitename: ", sitename)
	fmt.Fprint(w, " tenantname: ", tenantname)

	log.Println("ImageHook called", id, version, listname, sitename, tenantname)

}
