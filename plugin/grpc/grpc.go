package grpc

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/debug"
	"github.com/coredns/coredns/plugin/pkg"
	"github.com/coredns/coredns/request"

	"github.com/miekg/dns"
	ot "github.com/opentracing/opentracing-go"
)

// GRPC represents a plugin instance that can proxy requests to another (DNS) server via gRPC protocol.
// It has a list of proxies each representing one upstream proxy.
type GRPC struct {
	proxies []*Proxy
	p       pkg.Policy

	from    string
	ignored []string

	tlsConfig     *tls.Config
	tlsServerName string

	Next plugin.Handler
}

// ServeDNS implements the plugin.Handler interface.
func (g *GRPC) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}
	if !pkg.Match(g, state, g.Name(), g.ignored) {
		return plugin.NextOrFailure(g.Name(), g.Next, ctx, w, r)
	}

	var (
		span, child ot.Span
		ret         *dns.Msg
		err         error
		i           int
	)
	span = ot.SpanFromContext(ctx)
	list := g.list()
	deadline := time.Now().Add(defaultTimeout)

	for time.Now().Before(deadline) {
		if i >= len(list) {
			// reached the end of list without any answer
			if ret != nil {
				// write empty response and finish
				w.WriteMsg(ret)
			}
			break
		}

		proxy := list[i]
		i++

		if span != nil {
			child = span.Tracer().StartSpan("query", ot.ChildOf(span.Context()))
			ctx = ot.ContextWithSpan(ctx, child)
		}

		ret, err = proxy.query(ctx, r)
		if err != nil {
			// Continue with the next proxy
			continue
		}

		if child != nil {
			child.Finish()
		}

		// Check if the reply is correct; if not return FormErr.
		if !state.Match(ret) {
			debug.Hexdumpf(ret, "Wrong reply for id: %d, %s %d", ret.Id, state.QName(), state.QType())

			formerr := new(dns.Msg)
			formerr.SetRcode(state.Req, dns.RcodeFormatError)
			w.WriteMsg(formerr)
			return 0, nil
		}

		w.WriteMsg(ret)
		return 0, nil
	}

	return 0, nil
}

// NewGRPC returns a new GRPC.
func newGRPC() *GRPC {
	g := &GRPC{
		p: new(pkg.Random),
	}
	return g
}

// Name implements the Handler interface.
func (g *GRPC) Name() string { return "grpc" }

// Len returns the number of configured proxies.
func (g *GRPC) len() int { return len(g.proxies) }

// List returns a set of proxies to be used for this client depending on the policy in p.
//func (g *GRPC) list() []*Proxy { return g.p.List(g.proxies).([]*Proxy) }
func (g *GRPC) list() []*Proxy {
	slice := g.p.List(structToInterface(g.proxies))
	proxySlice := make([]*Proxy, len(slice))
	for _, s := range slice {
		proxySlice = append(proxySlice, s.(*Proxy))
	}
	return proxySlice
}

func (g *GRPC) From() string { return g.from }

const defaultTimeout = 5 * time.Second

func structToInterface(proxies []*Proxy) []interface{} {
	res := make([]interface{}, len(proxies))

	for _, p := range proxies {
		res = append(res, p)
	}
	return res
}
