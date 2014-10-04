package whois

import (
	"net/url"
	"strings"
)

type nrAdapter struct {
	DefaultAdapter
}

func (a *nrAdapter) Resolve(req *Request) error {
	labels := strings.SplitN(req.Query, ".", 2)
	values := url.Values{}
	values.Set("subdomain", labels[0])
	values.Set("tld", labels[1])
	req.URL = "http://cenpac.net.nr/dns/whois.html?" + values.Encode()
	req.Body = ""
	return nil
}

func init() {
	BindAdapter(
		&nrAdapter{},
		"cenpac.net.nr",
	)
}
