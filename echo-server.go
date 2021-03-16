package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>")
}

func cookieHandler(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))
	http.SetCookie(w, &http.Cookie{
		Name:  "cookie_field",
		Value: "cookie_value",
		Path:  "/",
	})
	w.WriteHeader(http.StatusOK)
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	http.HandleFunc("/cookie", cookieHandler)
	log.Println("start http listning :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
