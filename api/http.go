package api

import (
	"fmt"
	"github.com/sukhajata/devicetwin/pkg/loggerhelper"
	"log"
	"net/http"
)

type HealthCheckServer struct {
	Ready bool
	Live  bool
}

// readinessHandler reports ready status
func (s *HealthCheckServer) readinessHandler(w http.ResponseWriter, r *http.Request) {
	if !s.Ready {
		http.Error(w, "Not ready", http.StatusInternalServerError)
		return
	}
	_, err := fmt.Fprintf(w, "Ready")
	if err != nil {
		loggerhelper.WriteToLog(err.Error())
	}
}

// livenessHandler reports live status
func (s *HealthCheckServer) livenessHandler(w http.ResponseWriter, r *http.Request) {
	if !s.Live {
		http.Error(w, "Not live", http.StatusInternalServerError)
		return
	}
	_, err := fmt.Fprintf(w, "Live")
	if err != nil {
		loggerhelper.WriteToLog(err.Error())
	}
}

func NewHealthCheckServer() {
	s := &HealthCheckServer{
		Live:  true,
		Ready: true,
	}

	http.HandleFunc("/health/ready", s.readinessHandler)
	http.HandleFunc("/health/live", s.livenessHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
