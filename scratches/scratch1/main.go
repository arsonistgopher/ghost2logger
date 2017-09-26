// This code is an attempt to build a simple demo of a foreground, Docker and journald app with basic logging.
//
// Our contrived example launches a HTTP server and offers a route for '/'.
// Logs are generated at various stages of service/app start/stop/run and servicing of calls.
//
// Copyright David Gee 2017, arsonistgopher.com
// This is released under the CRAPL License: http://matt.might.net/articles/crapl/

package main

import (
	"fmt"
	"net/http"

	// This package gives us good logging capabilities inline with our requirements.
	// It claims to handle stdout/stderr/journal!
	"github.com/coreos/go-log/log"

	// This is our foundation service.
	// This allows us to run as a foreground application as well as systemd service.
	"github.com/kardianos/service"
)

// Required in order to instantiate the methods required for the service.
type program struct{}

// This is a HTTP handler and not part of the service package.
func handler(w http.ResponseWriter, r *http.Request) {
	log.Info("[Scratch1] Request Received for: ", r.URL)
	fmt.Fprintf(w, "Hi Dave, I love you man!")
}

// This is part of the service package and represents our Start() entry point.
func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	log.Info("[Scratch1] Entered p.Start(s)")

	// Call/launch run as a Go Routine
	go p.run()
	return nil
}

func (p *program) run() {
	// Our basic run() method
	log.Info("[Scratch1] Entered p.run(s)")

	// Register the HTTP handler
	http.HandleFunc("/", handler)

	log.Info("[Scratch1] Registered handler() for HTTP / call")
	log.Info("[Scratch1] Serving http://localhost:8181")

	// ListenAndServe does not return, allowing our GR
	http.ListenAndServe(":8181", nil)
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	log.Info("[Scratch1] Entered p.Stop(s)")
	return nil
}

func main() {

	// Pretty comments
	svcConfig := &service.Config{
		Name:        "Scratch1 Service",
		DisplayName: "Go Scratch1 Service",
		Description: "This is the Scratch1 multi-purpose award winning service and foreground app.",
	}

	// Create a program struct
	prg := &program{}

	// Create a new service
	s, err := service.New(prg, svcConfig)

	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatalln(err)
	}

	// This runs the exported Run() method which calls Start() in turn.
	err = s.Run()

	if err != nil {
		log.Fatalln(err)
	}
}
