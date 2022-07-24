package neisgo

import (
	"net/url"
	"strings"
)

func (p Parameters) format() string {
	var params []string

	for name, value := range p {
		params = append(params, url.QueryEscape(name)+"="+url.QueryEscape(value))
	}

	return "?" + strings.Join(params, "&")
}
