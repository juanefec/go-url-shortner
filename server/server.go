package server

import (
	"log"
	"net/http"
	"time"

	"github.com/juanefec/go-url-shortner/shortner"
)

var handler http.HandlerFunc

func init() {
	handler = URLShortnerHandler
}

type Server struct {
	*http.Server
}

func NewServer() *Server {
	s := &http.Server{
		Addr:           ":4444",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{s}
}

func URLShortnerHandler(res http.ResponseWriter, req *http.Request) {
	url := req.URL
	switch url.Path {
	case "/":
		if id := url.Query().Get("i"); id != "" {
			og, e := shortner.GetOriginal(id)
			log.Printf("\nGOTO URL: %v\nOG URL: %v", req.URL.String(), og)
			if e != nil {
				respond(res, req, http.StatusBadRequest, e.Error())
				return
			}
			if og != "" {
				http.Redirect(res, req, og, 303)
				return
			}
		}
		respond(res, req, http.StatusBadRequest, "wtf: no se obtubo la query i")
		return
	case "/set":
		s, e := shortner.Shorten(url.Query().Get("url"))
		if e != nil {
			respond(res, req, http.StatusBadRequest, e.Error())
			return
		}
		respond(res, req, http.StatusOK, s)
		return
	}
	res.WriteHeader(http.StatusNotFound)
}

func respond(res http.ResponseWriter, req *http.Request, statusCode int, body string) {
	log.Printf("\nURL stored: %v \nNew url: %v", req.URL.String(), body)
	res.WriteHeader(statusCode)
	if body != "" {
		res.Write([]byte(body))
	}
}
