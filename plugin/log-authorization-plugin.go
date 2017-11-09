package plugin

import (
	"github.com/docker/go-plugins-helpers/authorization"
	"fmt"
	"encoding/json"
	"os"
)

type LogAuthorizationPlugin struct {
	Headers bool
	Body    bool
}

func NewLogAuthorizationPlugin() *LogAuthorizationPlugin {
	return &LogAuthorizationPlugin{
		Headers: false,
		Body:    false,
	}
}

func (p *LogAuthorizationPlugin) AuthZReq(req authorization.Request) authorization.Response {
	if !p.Headers {
		req.RequestHeaders = nil
	}
	if !p.Body {
		req.RequestBody = nil
	}

	output, err := json.Marshal(req)
	if err == nil {
		fmt.Printf("request: %s\n", string(output))
	} else {
		fmt.Fprintf(os.Stderr, "request: marshalling error %v\n", err)
	}

	return authorization.Response{Allow: true}
}

func (p *LogAuthorizationPlugin) AuthZRes(res authorization.Request) authorization.Response {
	if !p.Headers {
		res.RequestHeaders = nil
		res.ResponseHeaders = nil
	}
	if !p.Body {
		res.RequestBody = nil
		res.ResponseBody = nil
	}

	output, err := json.Marshal(res)
	if err == nil {
		fmt.Printf("response: %s\n", string(output))
	} else {
		fmt.Fprintf(os.Stderr, "response: marshalling error %v\n", err)
	}

	return authorization.Response{Allow: true}
}
