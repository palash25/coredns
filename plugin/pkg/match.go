package pkg

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"

	"github.com/miekg/dns"
)

type forwardingPlugin interface {
	ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error)
	Name() string
	From() string
}

func Match(p forwardingPlugin, state request.Request, name string, ignored []string) bool {
	if !plugin.Name(name).Matches(state.Name()) || !isAllowedDomain(p, state.Name(), ignored) {
		return false
	}

	return true
}

func isAllowedDomain(p forwardingPlugin, name string, ignored []string) bool {
	if dns.Name(name) == dns.Name(p.From()) {
		return true
	}

	for _, ignore := range ignored {
		if plugin.Name(ignore).Matches(name) {
			return false
		}
	}
	return true
}
