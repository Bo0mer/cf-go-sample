package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type appConfig struct {
	AppName          string       `json:"application_name"`
	AppUris          []string     `json:"application_uris"`
	Limits           limitsConfig `json:"limits"`
	AppSpaceName     string       `json:"space_name"`
	AppInstanceIndex string
}

type limitsConfig struct {
	Disk int `json:"disk"`
	FDs  int `json:"fds"`
	Mem  int `json:"mem"`
}

func main() {
	t, err := template.New("index.html").Parse(indexHTMLTemplate)
	if err != nil {
		log.Fatalf("error parsing index.html template: %v", err)
	}

	srv := &server{t: t}
	http.ListenAndServe(":7777", srv)
}

type server struct {
	t *template.Template
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		s.serveTemplate(w, req)
		return
	}
	http.FileServer(http.Dir("./public")).ServeHTTP(w, req)
}

func (s *server) serveTemplate(w http.ResponseWriter, req *http.Request) {
	cfg, err := appConfigFromEnv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.t.ExecuteTemplate(w, "index.html", cfg)
}

func appConfigFromEnv() (*appConfig, error) {
	vcapApp := os.Getenv("VCAP_APPLICATION")
	var cfg appConfig
	if err := json.Unmarshal([]byte(vcapApp), &cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON from VCAP_APPLICATION env: %v", err)
	}

	cfg.AppInstanceIndex = os.Getenv("INSTANCE_INDEX")
	return &cfg, nil
}
