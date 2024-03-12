package restapi

import (
	"fmt"
	"net/http"
)

func ImageHook(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	etag := r.URL.Query().Get("etag")
	listname := r.URL.Query().Get("listname")
	sitename := r.URL.Query().Get("sitename")
	tenantname := r.URL.Query().Get("tenant")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(200)
	fmt.Fprint(w, " id:", id)
	fmt.Fprint(w, " etag:", etag)
	fmt.Fprint(w, " listname", listname)
	fmt.Fprint(w, " sitename", sitename)
	fmt.Fprint(w, " tenantname", tenantname)

}
