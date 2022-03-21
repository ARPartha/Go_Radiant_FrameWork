package middleware

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"radiant/radiant"
	
)

type (
	// CORSConfig defines the config for CORS middleware.
	CORSConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper Skipper

		// AllowOrigin defines a list of origins that may access the resource.
		// Optional. Default value []string{"*"}.
		AllowOrigins []string `yaml:"allow_origins"`

		// AllowOriginFunc is a custom function to validate the origin. It takes the
		// origin as an argument and returns true if allowed or false otherwise. If
		// an error is returned, it is returned by the handler. If this option is
		// set, AllowOrigins is ignored.
		// Optional.
		AllowOriginFunc func(origin string) (bool, error) `yaml:"allow_origin_func"`

		// AllowMethods defines a list methods allowed when accessing the resource.
		// This is used in response to a preflight request.
		// Optional. Default value DefaultCORSConfig.AllowMethods.
		AllowMethods []string `yaml:"allow_methods"`

		// AllowHeaders defines a list of request headers that can be used when
		// making the actual request. This is in response to a preflight request.
		// Optional. Default value []string{}.
		AllowHeaders []string `yaml:"allow_headers"`

		// AllowCredentials indicates whether or not the response to the request
		// can be exposed when the credentials flag is true. When used as part of
		// a response to a preflight request, this indicates whether or not the
		// actual request can be made using credentials.
		// Optional. Default value false.
		AllowCredentials bool `yaml:"allow_credentials"`

		// ExposeHeaders defines a whitelist headers that clients are allowed to
		// access.
		// Optional. Default value []string{}.
		ExposeHeaders []string `yaml:"expose_headers"`

		// MaxAge indicates how long (in seconds) the results of a preflight request
		// can be cached.
		// Optional. Default value 0.
		MaxAge int `yaml:"max_age"`
	}
)

var (
	// DefaultCORSConfig is the default CORS middleware config.
	DefaultCORSConfig = CORSConfig{
		Skipper:      DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}
)

// CORS returns a Cross-Origin Resource Sharing (CORS) middleware.
// See: https://developer.mozilla.org/en/docs/Web/HTTP/Access_control_CORS
func CORS() radiant.MiddlewareFunc {
	return CORSWithConfig(DefaultCORSConfig)
}

// CORSWithConfig returns a CORS middleware with config.
// See: `CORS()`.
func CORSWithConfig(config CORSConfig) radiant.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultCORSConfig.Skipper
	}
	if len(config.AllowOrigins) == 0 {
		config.AllowOrigins = DefaultCORSConfig.AllowOrigins
	}
	if len(config.AllowMethods) == 0 {
		config.AllowMethods = DefaultCORSConfig.AllowMethods
	}

	allowOriginPatterns := []string{}
	for _, origin := range config.AllowOrigins {
		pattern := regexp.QuoteMeta(origin)
		pattern = strings.Replace(pattern, "\\*", ".*", -1)
		pattern = strings.Replace(pattern, "\\?", ".", -1)
		pattern = "^" + pattern + "$"
		allowOriginPatterns = append(allowOriginPatterns, pattern)
	}

	allowMethods := strings.Join(config.AllowMethods, ",")
	allowHeaders := strings.Join(config.AllowHeaders, ",")
	exposeHeaders := strings.Join(config.ExposeHeaders, ",")
	maxAge := strconv.Itoa(config.MaxAge)

	return func(next radiant.HandlerFunc) radiant.HandlerFunc {
		return func(c radiant.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			origin := req.Header.Get(radiant.HeaderOrigin)
			allowOrigin := ""

			preflight := req.Method == http.MethodOptions
			res.Header().Add(radiant.HeaderVary, radiant.HeaderOrigin)

			// No Origin provided
			if origin == "" {
				if !preflight {
					return next(c)
				}
				return c.NoContent(http.StatusNoContent)
			}

			if config.AllowOriginFunc != nil {
				allowed, err := config.AllowOriginFunc(origin)
				if err != nil {
					return err
				}
				if allowed {
					allowOrigin = origin
				}
			} else {
				// Check allowed origins
				for _, o := range config.AllowOrigins {
					if o == "*" && config.AllowCredentials {
						allowOrigin = origin
						break
					}
					if o == "*" || o == origin {
						allowOrigin = o
						break
					}
					if matchSubdomain(origin, o) {
						allowOrigin = origin
						break
					}
				}

				// Check allowed origin patterns
				for _, re := range allowOriginPatterns {
					if allowOrigin == "" {
						didx := strings.Index(origin, "://")
						if didx == -1 {
							continue
						}
						domAuth := origin[didx+3:]
						// to avoid regex cost by invalid long domain
						if len(domAuth) > 253 {
							break
						}

						if match, _ := regexp.MatchString(re, origin); match {
							allowOrigin = origin
							break
						}
					}
				}
			}

			// Origin not allowed
			if allowOrigin == "" {
				if !preflight {
					return next(c)
				}
				return c.NoContent(http.StatusNoContent)
			}

			// Simple request
			if !preflight {
				res.Header().Set(radiant.HeaderAccessControlAllowOrigin, allowOrigin)
				if config.AllowCredentials {
					res.Header().Set(radiant.HeaderAccessControlAllowCredentials, "true")
				}
				if exposeHeaders != "" {
					res.Header().Set(radiant.HeaderAccessControlExposeHeaders, exposeHeaders)
				}
				return next(c)
			}

			// Preflight request
			res.Header().Add(radiant.HeaderVary, radiant.HeaderAccessControlRequestMethod)
			res.Header().Add(radiant.HeaderVary, radiant.HeaderAccessControlRequestHeaders)
			res.Header().Set(radiant.HeaderAccessControlAllowOrigin, allowOrigin)
			res.Header().Set(radiant.HeaderAccessControlAllowMethods, allowMethods)
			if config.AllowCredentials {
				res.Header().Set(radiant.HeaderAccessControlAllowCredentials, "true")
			}
			if allowHeaders != "" {
				res.Header().Set(radiant.HeaderAccessControlAllowHeaders, allowHeaders)
			} else {
				h := req.Header.Get(radiant.HeaderAccessControlRequestHeaders)
				if h != "" {
					res.Header().Set(radiant.HeaderAccessControlAllowHeaders, h)
				}
			}
			if config.MaxAge > 0 {
				res.Header().Set(radiant.HeaderAccessControlMaxAge, maxAge)
			}
			return c.NoContent(http.StatusNoContent)
		}
	}
}
