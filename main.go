package main

import (
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/jonasprogrammer/docker-plugin-auth-debug/plugin"
	"os/user"
	"strconv"
)

func main() {
	p := plugin.NewLogAuthorizationPlugin()
	h := authorization.NewHandler(p)
	u, _ := user.Lookup("root")
	gid, _ := strconv.Atoi(u.Gid)
	h.ServeUnix("auth_debug_plugin", gid)
}
