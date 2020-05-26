package exporter

import (
	"context"
	"fmt"

	"github.com/falcosecurity/client-go/pkg/api/output"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

// A RecvFunc waits for subscribed events and forwards to metric counters.
type RecvFunc func() error

var (
	eventsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "falco_events",
		},
		[]string{
			"rule",
			"priority",
			"hostname",
		},
	)
)

func init() {
	prometheus.MustRegister(eventsCounter)
}

// Subscribe to a ServiceClient to receive a stream of Falco output events.
// If success, it returns a RecvFunc that can be used.
// To stop the subscription, use Cancel() on the context provided.
func Subscribe(ctx context.Context, outputClient output.ServiceClient, opts ...grpc.CallOption) (RecvFunc, error) {

	// Keepalive true means that the client will wait indefinitely for new events to come
	fcs, err := outputClient.Subscribe(ctx, &output.Request{Keepalive: true}, opts...)
	if err != nil {
		return nil, err
	}

	return func() error {
		for {
			res, err := fcs.Recv()
			if err != nil {
				return err
			}

			eventsCounter.With(prometheus.Labels{
				"rule":     res.Rule,
				"priority": fmt.Sprintf("%d", res.Priority),
				"hostname": res.Hostname,
			}).Inc()

			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
		}
	}, nil
}
