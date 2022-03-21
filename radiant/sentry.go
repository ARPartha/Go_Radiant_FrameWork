package radiant

import (
	c "context"
	"net/http"

	"time"

	"github.com/getsentry/sentry-go"
)

const valuesKey = "sentry"

type Sentryhandler struct {
	repanic         bool
	waitForDelivery bool
	timeout         time.Duration
}

type SentryOptions struct {
	// Repanic configures whether Sentry should repanic after recovery, in most cases it should be set to true,
	// as radiant includes it's own Recover middleware what handles http responses.
	Repanic bool
	// WaitForDelivery configures whether you want to block the request before moving forward with the response.
	// Because radiant's Recover handler doesn't restart the application,
	// it's safe to either skip this option or set it to false.
	WaitForDelivery bool
	// Timeout for the event delivery requests.
	Timeout time.Duration
}

// New returns a function that satisfies HandlerFunc interface
// It can be used with Use() methods.
func SentryNew(options SentryOptions) MiddlewareFunc {
	timeout := options.Timeout
	if timeout == 0 {
		timeout = 2 * time.Second
	}
	return (&Sentryhandler{
		repanic:         options.Repanic,
		timeout:         timeout,
		waitForDelivery: options.WaitForDelivery,
	}).Sentryhandle
}

func (h *Sentryhandler) Sentryhandle(next HandlerFunc) HandlerFunc {
	return func(ctx Context) error {
		hub := sentry.GetHubFromContext(ctx.Request().Context())
		if hub == nil {
			hub = sentry.CurrentHub().Clone()
		}
		hub.Scope().SetRequest(ctx.Request())
		ctx.Set(valuesKey, hub)
		defer h.recoverWithSentry(hub, ctx.Request())
		return next(ctx)
	}
}

func (h *Sentryhandler) recoverWithSentry(hub *sentry.Hub, r *http.Request) {
	if err := recover(); err != nil {
		eventID := hub.RecoverWithContext(
			c.WithValue(r.Context(), sentry.RequestContextKey, r),
			err,
		)
		if eventID != nil && h.waitForDelivery {
			panic(err)
		}
	}
}

// GetHubFromContext retrieves attached *sentry.Hub instance from Context.
func GetHubFromContext(ctx Context) *sentry.Hub {
	if hub, ok := ctx.Get(valuesKey).(*sentry.Hub); ok {
		return hub
	}
	return nil
}
