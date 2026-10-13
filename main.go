package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"schedsim/internal/scheduler"
)

//go:embed web/index.html
var indexHTML []byte

type simReq struct {
	Algo      string              `json:"algo"`
	Quantum   int                 `json:"quantum"`
	Processes []scheduler.Process `json:"processes"`
}

var simCount atomic.Int64

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(indexHTML)
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "ok") })
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "# TYPE schedsim_simulations_total counter\nschedsim_simulations_total %d\n", simCount.Load())
	})
	http.HandleFunc("/simulate", func(w http.ResponseWriter, r *http.Request) {
		var req simReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res, err := scheduler.Simulate(req.Algo, req.Processes, req.Quantum)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		simCount.Add(1)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"result": res, "metrics": res.Metrics()})
	})
	log.Println("SchedSim GUI: open http://localhost" + *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
