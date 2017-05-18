package context // import "gopkg.in/webnice/web.v1/context"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import "gopkg.in/webnice/web.v1/context/route"
import "gopkg.in/webnice/web.v1/context/errors"
import "gopkg.in/webnice/web.v1/context/handlers"
import (
	stdContext "context"
	"net/http"
)

// New returns a new routing context object
// You can pass the following types of objects as arguments:
// from "net/http" type *http.Request;
// fom "context" interface context.Context.
// If an invalid argument type is passed, the function will return nil
func New(obj ...interface{}) Interface {
	var ctx *impl
	var i int

	for i = range obj {
		switch val := obj[i].(type) {
		case *http.Request:
			ctx = request(val)
		case stdContext.Context:
			ctx = context(val)
		default:
			// invalid argument type is passed
			return nil
		}
		if ctx != nil {
			return ctx
		}
	}

	ctx = new(impl)
	ctx.route = route.New()
	ctx.errors = errors.New()
	ctx.handlers = handlers.New()

	return ctx
}

// Get the routing Context object from a http context
func context(cx stdContext.Context) (ret *impl) {
	var ok bool
	if ret, ok = cx.Value(_ContextKey).(*impl); ok {
		return
	}
	return
}

// Get the routing context object from a http context
func request(rq *http.Request) *impl { return context(rq.Context()) }

// IsContext Check if a context not empty in net/http context
func IsContext(rq *http.Request) bool { return request(rq) != nil }

// Route interface
func (ctx *impl) Route() route.Interface { return ctx.route }

// Error interface
func (ctx *impl) Errors() errors.Interface { return ctx.errors }

// Handlers interface
func (ctx *impl) Handlers() handlers.Interface { return ctx.handlers }

// NewRequest Creates new http request and copy context from parent request to new request
func (ctx *impl) NewRequest(rq *http.Request) *http.Request {
	return rq.WithContext(stdContext.WithValue(rq.Context(), _ContextKey, ctx))
}
