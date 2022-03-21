package radiant

import (
	"html/template"
	"io"

	"github.com/foolin/goview"
)

const templateEngineKey = "foolin-goview-echoview"

// ViewEngine view engine for echo
type ViewEngine struct {
	*goview.ViewEngine
}

// New a new view engine
func ViewNew(config goview.Config) *ViewEngine {
	return Wrap(goview.New(config))
}

// Wrap wrap view engine for goview.ViewEngine
func Wrap(engine *goview.ViewEngine) *ViewEngine {
	return &ViewEngine{
		ViewEngine: engine,
	}
}

// Default new default config view engine
func Default() *ViewEngine {
	return ViewNew(goview.DefaultConfig)
}

// Render render template for echo interface
func (e *ViewEngine) Render(w io.Writer, name string, data interface{}, c Context) error {
	return e.RenderWriter(w, name, data)
}

// Render html render for template
// You should use helper func `Middleware()` to set the supplied
// TemplateEngine and make `Render()` work validly.
func Render(ctx Context, code int, name string, data interface{}) error {
	if val := ctx.Get(templateEngineKey); val != nil {
		if e, ok := val.(*ViewEngine); ok {
			return e.Render(ctx.Response().Writer, name, data, ctx)
		}
	}
	return ctx.Render(code, name, data)
}

// NewMiddleware echo middleware for func `echoview.Render()`
func NewMiddleware(config goview.Config) MiddlewareFunc {
	return Middleware(ViewNew(config))
}

// Middleware echo middleware wrapper
func Middleware(e *ViewEngine) MiddlewareFunc {
	return func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			c.Set(templateEngineKey, e)
			return next(c)
		}
	}
}
func setMapFunctions() template.FuncMap {
	return template.FuncMap{}
}

func (e *Radiant)RenderTemplate() {
	e.Renderer = ViewNew(goview.Config{
		Root:         "views",
		Extension:    ".gohtml",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        setMapFunctions(),
		DisableCache: false,
		Delims:       goview.Delims{Left: "{{", Right: "}}"},
	})
}
