package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	// 	"net/http/httputil"
)

// less /usr/lib/golang/src/net/http/cookiejar/jar.go

func main() {
	jar, _ := cookiejar.New(nil)
	client := http.Client{
		Jar: jar,
	}

	for i := 0; i < 2; i++ {
		_, err := client.Get("http://localhost:18888/cookie")
		set_cookie_url, _ := url.Parse("http://localhost:18888/cookie")
		log.Println("---")
		log.Println(client.Jar.Cookies(set_cookie_url))
		log.Println("---")

		if err != nil {
			panic(err)
		}
		// 		dump, _ := httputil.DumpResponse(resp, true)
		// 		log.Println(string(dump))
		//  2021/03/16 11:00:24 HTTP/1.1 200 OK
		//  Content-Length: 31
		//  Content-Type: text/html; charset=utf-8
		//  Date: Tue, 16 Mar 2021 11:00:24 GMT

		//  <html><body>hello</body></html>
	}
}
