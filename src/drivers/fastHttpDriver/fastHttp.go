package fastHttpDriver

import (
	"time"

	"github.com/valyala/fasthttp"
)

//go:generate mockery -name=FastHttpClient
type FastHttpClient interface {
	DoTimeout(req *fasthttp.Request, resp *fasthttp.Response, timeout time.Duration) error
}
