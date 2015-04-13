package main

import (
	"io/ioutil"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/go-zoo/claw"
	"github.com/go-zoo/claw/mw"
)

const (
	INDEX   = "index.html"
	CONTACT = "contact.html"
	OUTILS  = "outils.html"
)

var (
	IndexCont   []byte
	ContactCont []byte
	OutilsCont  []byte
)

func init() {
	IndexCont, _ = ioutil.ReadFile(INDEX)
	ContactCont, _ = ioutil.ReadFile(CONTACT)
	OutilsCont, _ = ioutil.ReadFile(OUTILS)
}

func main() {
	muxx := bone.New()
	middle := claw.New(mw.Logger)

	muxx.Get("/", http.HandlerFunc(IndexHandler))
	muxx.Get("/contact", http.HandlerFunc(ContactHandler))
	muxx.Get("/outils", http.HandlerFunc(OutilsHandler))

	muxx.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	muxx.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	muxx.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))
	muxx.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	muxx.Handle("/skins/", http.StripPrefix("/skins/", http.FileServer(http.Dir("skins"))))

	http.ListenAndServe(":8080", middle.Merge(muxx))
}

func IndexHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write(IndexCont)
}

func ContactHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write(ContactCont)
}

func OutilsHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write(OutilsCont)
}
