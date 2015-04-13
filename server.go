package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/go-zoo/claw"
	"github.com/go-zoo/claw/mw"
)

type Frame struct {
	Head   template.HTML
	Header template.HTML
	Footer template.HTML
	Foot   template.HTML
}

type Service struct {
	Frame    Frame
	Services template.HTML
}

const (
	INDEX    = "index.html"
	CONTACT  = "contact.html"
	OUTILS   = "outils.html"
	ABOUT    = "about.html"
	SERVICE  = "service.html"
	NOTFOUND = "404.html"

	HEAD   = "template/head.html"
	HEADER = "template/header.html"
	FOOTER = "template/footer.html"
	FOOT   = "template/foot.html"
)

var (
	muxx = bone.New()

	WebFrame    = Frame{}
	ServiceList = Service{}
)

func init() {
	head, _ := ioutil.ReadFile(HEAD)
	header, _ := ioutil.ReadFile(HEADER)
	footer, _ := ioutil.ReadFile(FOOTER)
	foot, _ := ioutil.ReadFile(FOOT)

	WebFrame.Head = template.HTML(head)
	WebFrame.Header = template.HTML(header)
	WebFrame.Footer = template.HTML(footer)
	WebFrame.Foot = template.HTML(foot)

	ServiceList.Frame = WebFrame
}

func main() {
	muxx.NotFound(NotFound)
	middle := claw.New(mw.Logger)

	muxx.Get("/", http.HandlerFunc(IndexHandler))
	muxx.Get("/contact", http.HandlerFunc(ContactHandler))
	muxx.Get("/outils", http.HandlerFunc(OutilsHandler))
	muxx.Get("/about", http.HandlerFunc(AboutHandler))
	muxx.Get("/services/:name", http.HandlerFunc(ServiceHandler))

	muxx.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	muxx.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	muxx.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))
	muxx.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	muxx.Handle("/skins/", http.StripPrefix("/skins/", http.FileServer(http.Dir("skins"))))

	http.ListenAndServe(":8080", middle.Merge(muxx))
}

func IndexHandler(rw http.ResponseWriter, req *http.Request) {
	if req.RequestURI != "/" {
		muxx.HandleNotFound(rw, req)
		return
	}
	tmpl, err := template.ParseFiles(INDEX)
	if err != nil {
		muxx.HandleNotFound(rw, req)
		return
	}
	tmpl.Execute(rw, WebFrame)
}

func ContactHandler(rw http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles(CONTACT)
	if err != nil {
		muxx.HandleNotFound(rw, req)
		return
	}
	tmpl.Execute(rw, WebFrame)
}

func OutilsHandler(rw http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles(OUTILS)
	if err != nil {
		muxx.HandleNotFound(rw, req)
		return
	}
	tmpl.Execute(rw, WebFrame)
}

func AboutHandler(rw http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles(ABOUT)
	if err != nil {
		muxx.HandleNotFound(rw, req)
		return
	}
	tmpl.Execute(rw, WebFrame)
}

func ServiceHandler(rw http.ResponseWriter, req *http.Request) {
	value := bone.GetValue(req, "name")
	prodData, err := ioutil.ReadFile(fmt.Sprintf("template/services/%s.html", value))
	if err != nil {
		muxx.HandleNotFound(rw, req)
		return
	}
	ServiceList.Services = template.HTML(prodData)
	tmpl, err := template.ParseFiles(SERVICE)
	if err != nil {
		muxx.HandleNotFound(rw, req)
		return
	}
	tmpl.Execute(rw, ServiceList)
}

func NotFound(rw http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles(NOTFOUND)
	if err != nil {
		muxx.HandleNotFound(rw, req)
		return
	}
	tmpl.Execute(rw, WebFrame)
}
