package main

import (
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/jonasprogrammer/docker-plugin-auth-debug/plugin"
	"os/user"
	"strconv"
	"flag"
)

func main() {
	p := plugin.NewLogAuthorizationPlugin()

	flag.BoolVar(&p.Headers, "headers", false, "show headers")
	flag.BoolVar(&p.Body, "body", false, "show request/response body")

	flag.Parse()

	h := authorization.NewHandler(p)
	u, _ := user.Lookup("root")
	gid, _ := strconv.Atoi(u.Gid)
	h.ServeUnix("auth_debug_plugin", gid)
}
