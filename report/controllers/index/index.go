package index

import(
	"net/http"
	"log"
)

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Printf("report receive. 404 request, data:%v, method:%v, uri:%v, host:%v, ip:%v", r.Form,  r.Method, r.RequestURI, r.Host, r.RemoteAddr)

	http.Redirect(w, r, "", 404)
}
