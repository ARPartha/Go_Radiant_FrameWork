package radiant

import (
	"fmt"
	"net/http"
	"radiant/radiant/core"

	"github.com/getsentry/sentry-go"
)

func SentryInit() bool {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: core.Configure.Error["SentryDSN"],
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
		return false
	}
	return true
}

func CustomHTTPErrorHandler(err error, c Context) {

	var renderErr error
	if he, ok := err.(*HTTPError); ok {
		vars := map[string]interface{}{
			"code":    he.Code,
			"message": he.Message,
		}

		renderErr = c.Render(he.Code, "error.gohtml", vars)
	} else {
		vars := map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		}
		renderErr = c.Render(http.StatusInternalServerError, "error.gohtml", vars)
	}

	if renderErr != nil {
		c.Logger().Fatal(renderErr)
	}
}
