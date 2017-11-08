package plugin

import (
	"github.com/docker/go-plugins-helpers/authorization"
	"fmt"
)

type LogAuthorizationPlugin struct {
}

func NewLogAuthorizationPlugin() *LogAuthorizationPlugin {
	return &LogAuthorizationPlugin{}
}

func (*LogAuthorizationPlugin) AuthZReq(req authorization.Request) authorization.Response {
	fmt.Println(req)

	return authorization.Response{Allow: true}
}

func (*LogAuthorizationPlugin) AuthZRes(res authorization.Request) authorization.Response {
	fmt.Println(res)

	return authorization.Response{Allow: true}
}
