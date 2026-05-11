package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var UserRegistrations = promauto.NewCounter(prometheus.CounterOpts{
	Name: "user_registrations_total",
	Help: "Total number of successful user registrations.",
})

var UserLogins = promauto.NewCounter(prometheus.CounterOpts{
	Name: "user_logins_total",
	Help: "Total number of successful user logins.",
})

var CartCheckouts = promauto.NewCounter(prometheus.CounterOpts{
	Name: "cart_checkouts_total",
	Help: "Total number of cart checkout events published.",
})

var OrdersCreated = promauto.NewCounter(prometheus.CounterOpts{
	Name: "orders_created_total",
	Help: "Total number of orders successfully created from checkout events.",
})

var ProductsCreated = promauto.NewCounter(prometheus.CounterOpts{
	Name: "products_created_total",
	Help: "Total number of products created.",
})

var StoresCreated = promauto.NewCounter(prometheus.CounterOpts{
	Name: "stores_created_total",
	Help: "Total number of stores created.",
})

// ── Metrics HTTP server ───────────────────────────────────────────────────────

// StartServer starts a lightweight HTTP server exposing /metrics and /healthz.
// Call it in a goroutine from each service's main().
//
// Example:
//
//	go metrics.StartServer(9100)
func StartServer(port int) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		// Non-fatal: log and continue. The gRPC server is the main process.
		_ = err
	}
}
